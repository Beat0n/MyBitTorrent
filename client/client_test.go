package client

import (
	"MyBitTorrent/bencode"
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGetPeersList(t *testing.T) {
	torrentfile, err := bencode.Parse("../testfile/debian-iso.torrent")
	require.Nil(t, err)

	Url, err := BuildURL(torrentfile)
	require.Nil(t, err)
	fmt.Println(Url)

	peers, err := GetPeersList(Url)
	require.Nil(t, err)
	for i, peer := range peers {
		fmt.Printf("Peer: %d, IP: %s, Port: %d", i, peer.Ip, peer.Port)
	}
}
