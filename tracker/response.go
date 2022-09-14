package tracker

import (
	"time"
)

type TrackerResp struct {
	interval  time.Duration
	PeersList string
}
