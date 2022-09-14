package peer

import (
	"fmt"
	"net"
)

const (
	PeerPort int = 6666
	IpLen    int = 4
	PortLen  int = 2
	PeerLen  int = IpLen + PortLen
)

type Peer struct {
	Ip   net.IP
	Port uint16
	Id   [20]byte
}

func (peer Peer) GetAddress() string {
	return fmt.Sprintf("%s:%d", peer.Ip, peer.Port)
}

/*func (peer Peer) ConnectTracker(rawURL string) (*tracker.Response, error){
	Url, err := url.Parse(rawURL)
	if err != nil {
		return nil, err
	}
	params := url.Values{
		"infohash":
	}
}*/
