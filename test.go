package main

import (
	"MyBitTorrent/bencode"
	"MyBitTorrent/client"
	"fmt"
)

func main() {
	torrentfile, _ := bencode.Parse("./testfile/debian-iso.torrent")

	Url, _ := client.BuildURL(torrentfile)

	peers, _ := client.GetPeersList(Url)

	fmt.Println(peers)
}
