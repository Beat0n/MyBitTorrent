package peer

import (
	"fmt"
	"net"
)

type Peer struct {
	Ip   net.IP `bencode:"ip"`
	Port uint16 `bencode:"port"`
	//Id   [20]byte
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
