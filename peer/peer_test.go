package peer

import (
	"github.com/stretchr/testify/require"
	"net"
	"testing"
)

func TestPeer_GetAddress(t *testing.T) {
	peer := Peer{
		Ip:   net.IPv4(180, 150, 112, 23),
		Port: 50000,
	}
	expected := "180.150.112.23:50000"
	require.Equal(t, expected, peer.GetAddress())
}
