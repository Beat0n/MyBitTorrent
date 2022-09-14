package client

import (
	"MyBitTorrent/bencode"
	"MyBitTorrent/peer"
	"MyBitTorrent/tracker"
	"crypto/rand"
	"encoding/binary"
	"fmt"
	bc "github.com/jackpal/bencode-go"
	"net"
	"net/http"
	"net/url"
	"strconv"
)

func BuildURL(tf *bencode.TorrentFile) (string, error) {
	Url, err := url.Parse(tf.Announce)
	if err != nil {
		aerr := fmt.Errorf("Announce error: %s", err)
		fmt.Println(aerr)
		return "", aerr
	}
	PeerId := [20]byte{}
	_, err = rand.Read(PeerId[:])
	if err != nil {
		return "", err
	}
	params := url.Values{
		"info_hash":  []string{string(tf.InfoHash[:])},
		"uploaded":   []string{"0"},
		"downloaded": []string{"0"},
		"peer_id":    []string{string(PeerId[:])},
		"port":       []string{strconv.Itoa(peer.PeerPort)},
	}
	Url.RawQuery = params.Encode()
	return Url.String(), nil
}

func GetPeersList(Url string) ([]peer.Peer, error) {
	cli := &http.Client{Timeout: 15e9}
	resp, err := cli.Get(Url)
	if err != nil {
		return nil, fmt.Errorf("request error: %s", err)
	}
	defer resp.Body.Close()
	trackResp := new(tracker.TrackerResp)
	err = bc.Unmarshal(resp.Body, trackResp)
	if err != nil {
		return nil, err
	}
	peers := []byte(trackResp.PeersList)

	if len(peers)%peer.PeerLen != 0 {
		return nil, fmt.Errorf("wrong peers bytes")
	}
	peernum := len(peers) / peer.PeerLen
	peerinfos := make([]peer.Peer, peernum)
	for i := 0; i < peernum; i++ {
		offset := i * peer.PeerLen
		peerinfos[i].Ip = net.IP(peers[offset : offset+peer.IpLen])
		peerinfos[i].Port = binary.BigEndian.Uint16(peers[offset+peer.IpLen : offset+peer.PeerLen])
	}

	return peerinfos, nil
}
