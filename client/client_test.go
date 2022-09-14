package client

import (
	"MyBitTorrent/bencode"
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGetPeersList(t *testing.T) {
	torrentfile, err := bencode.Parse("../testfile/debian-11.5.0-amd64-netinst.iso.torrent")
	require.Nil(t, err)

	Url, err := BuildURL(torrentfile)
	require.Nil(t, err)

	peers, err := GetPeersList(Url)
	require.Nil(t, err)
	for _, peer := range peers {
		fmt.Printf("%#v", peer)
	}
}
