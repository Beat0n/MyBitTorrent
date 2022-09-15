package tracker

import (
	"MyBitTorrent/bencode"
	"encoding/binary"
	"fmt"
	bc "github.com/jackpal/bencode-go"
	"net"
	"net/http"
	"net/url"
	"strconv"
)

const (
	PeerPort int = 6666
	IpLen    int = 4
	PortLen  int = 2
	PeerLen  int = IpLen + PortLen
)

const IDLEN int = 20

type PeerInfo struct {
	Ip   net.IP
	Port uint16
}

type TrackerResp struct {
	Interval int    `bencode:"interval"`
	Peers    string `bencode:"peers"`
}

func buildUrl(tf *bencode.TorrentFile, peerId [IDLEN]byte) (string, error) {
	Url, err := url.Parse(tf.Announce)
	if err != nil {

		fmt.Println("Announce Error: " + tf.Announce)
		return "", err
	}

	params := url.Values{
		"info_hash":  []string{string(tf.InfoHash[:])},
		"peer_id":    []string{string(peerId[:])},
		"port":       []string{strconv.Itoa(PeerPort)},
		"uploaded":   []string{"0"},
		"downloaded": []string{"0"},
		"compact":    []string{"1"},
		"left":       []string{strconv.Itoa(tf.Length)},
	}

	Url.RawQuery = params.Encode()
	return Url.String(), nil
}

func buildPeerInfo(peers []byte) []PeerInfo {
	num := len(peers) / PeerLen
	if len(peers)%PeerLen != 0 {
		fmt.Println("Received malformed peers")
		return nil
	}
	infos := make([]PeerInfo, num)
	for i := 0; i < num; i++ {
		offset := i * PeerLen
		infos[i].Ip = net.IP(peers[offset : offset+IpLen])
		infos[i].Port = binary.BigEndian.Uint16(peers[offset+IpLen : offset+PeerLen])
	}
	return infos
}

func FindPeers(tf *bencode.TorrentFile, peerId [IDLEN]byte) []PeerInfo {
	url, err := buildUrl(tf, peerId)
	if err != nil {
		fmt.Println("Build Tracker Url Error: " + err.Error())
		return nil
	}

	cli := &http.Client{Timeout: 15e9}
	resp, err := cli.Get(url)
	if err != nil {
		fmt.Println("Fail to Connect to Tracker: " + err.Error())
		return nil
	}
	defer resp.Body.Close()

	trackResp := new(TrackerResp)
	err = bc.Unmarshal(resp.Body, trackResp)
	if err != nil {
		fmt.Println("Tracker Response Error" + err.Error())
		return nil
	}

	return buildPeerInfo([]byte(trackResp.Peers))
}
