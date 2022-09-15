package tracker

import (
	"MyBitTorrent/bencode"
	"bytes"
	"crypto/rand"
	"fmt"
	bc "github.com/jackpal/bencode-go"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"net/http"
	"testing"
	"time"
)

func TestTracker(t *testing.T) {

	tf, _ := bencode.Parse("../testfile/debian-iso.torrent")

	var peerId [IDLEN]byte
	_, _ = rand.Read(peerId[:])
	url, _ := buildUrl(tf, peerId)
	fmt.Println(url)

	cli := &http.Client{Timeout: 15 * time.Second}
	resp, _ := cli.Get(url)

	defer resp.Body.Close()

	peers := FindPeers(tf, peerId)
	for i, p := range peers {
		fmt.Printf("Peer %d, Ip: %s, Port: %d\n", i, p.Ip, p.Port)
		fmt.Println(p.Ip)
	}
}

func TestFindPeers(t *testing.T) {
	tf, _ := bencode.Parse("../testfile/debian-iso.torrent")

	var peerId [IDLEN]byte
	_, _ = rand.Read(peerId[:])
	url, _ := buildUrl(tf, peerId)
	fmt.Println(url)

	peers := FindPeers(tf, peerId)
	for i, peer := range peers {
		fmt.Printf("Peer: %d, Ip: %s, Port: %d\n", i, peer.Ip, peer.Port)
	}

	cli := &http.Client{Timeout: 15e9}
	resp, err := cli.Get(url)
	require.Nil(t, err)
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	require.Nil(t, err)
	fmt.Println(string(data))

	reader := bytes.NewReader(data)

	trackResp := new(TrackerResp)
	bc.Unmarshal(reader, trackResp)
	fmt.Println(trackResp)
}
