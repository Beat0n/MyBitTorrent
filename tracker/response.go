package tracker

import "MyBitTorrent/peer"

type Response struct {
	PeersList []peer.Peer
}
