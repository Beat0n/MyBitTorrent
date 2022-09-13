package bencode

type bencodeTorrent struct {
	//Tracker URL
	announce string
	//file information
	info bencodeInfo
	//optional tracker url
	announcelist []string
	//comment information
	comment string
	//
	createdby string
}

type bencodeInfo struct {
	length     int
	name       string
	piecelenth int
	pieces     []piece
}

type piece struct {
}
