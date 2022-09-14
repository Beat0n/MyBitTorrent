package tracker

type TrackerResp struct {
	Interval  int    `bencode:"interval"`
	PeersList string `bencode:"peers"`
}
