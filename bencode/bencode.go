package bencode

import (
	"bytes"
	"crypto/sha1"
	"fmt"
	"github.com/jackpal/bencode-go"
	"os"
)

const (
	HashLen = 20
)

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
	//文件块标志信息
	Pieces string `bencode:"pieces"`
}

type TorrentFile struct {
	Announce string
	//整个文件的标识信息
	InfoHash [20]byte
	//每个文件块的标识信息
	PieceHashes [][20]byte
	PieceLength int
	Length      int
	Name        string
}

func (info *BencodeInfo) hash() ([20]byte, error) {
	var buf bytes.Buffer
	err := bencode.Marshal(&buf, *info)
	if err != nil {
		return [20]byte{}, err
	}
	h := sha1.Sum(buf.Bytes())
	return h, nil
}

func (info *BencodeInfo) GetPieceHashes() ([][20]byte, error) {
	buf := []byte(info.Pieces)
	if len(buf)%HashLen != 0 && HashLen != 0 {
		return nil, fmt.Errorf("Wrong pieces of length: %d", len(buf))
	}
	numhashes := len(buf) / HashLen
	hashes := make([][20]byte, numhashes)

	for i := 0; i < numhashes; i++ {
		copy(hashes[i][:], buf[i*HashLen:(i+1)*HashLen])
	}
	return hashes, nil
}

func (bto *BencodeTorrent) toTorrenFile() (*TorrentFile, error) {
	h, err := bto.Info.hash()
	if err != nil {
		return nil, err
	}
	hashes, err := bto.Info.GetPieceHashes()
	if err != nil {
		return nil, err
	}
	t := &TorrentFile{
		bto.Announce,
		h,
		hashes,
		bto.Info.Piecelenth,
		bto.Info.Length,
		bto.Info.Name,
	}
	return t, nil
}

func Parse(torrentname string) (*TorrentFile, error) {
	file, err := os.OpenFile(torrentname, os.O_RDONLY, 0666)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	bto := BencodeTorrent{}
	err = bencode.Unmarshal(file, &bto)
	if err != nil {
		return nil, err
	}
	torrentfile, err := bto.toTorrenFile()
	if err != nil {
		return nil, err
	}
	return torrentfile, nil
}
