package bencode

type BencodeTorrent struct {
	//Tracker URL
	Announce string `bencode:"announce"`
	//file information
	Info BencodeInfo `bencode:"info"`
}

type BencodeInfo struct {
	//总文件大小
	Length int    `bencode:"length"`
	Name   string `bencode:"name"`
	//文件块大小
	Piecelenth int `bencode:"piece length"`
	//文件标志信息
	Pieces string `bencode:"pieces"`
}

type TorrentFile struct {
	Announce    string
	Length      int
	PieceLength int
	Name        string
	InfoHash    [20]byte
	PieceHashes [][20]byte
}

func (bto *BencodeTorrent) toTorrenFile() (*TorrentFile, error) {
	t := &TorrentFile{
		bto.Announce,
		bto.Info.Length,
		bto.Info.Piecelenth,
		bto.Info.Name,
	}
	return t, nil
}
