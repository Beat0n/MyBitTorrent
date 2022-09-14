package bencode

import (
	"encoding/json"
	"flag"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"testing"
)

var update = flag.Bool("update", false, "update .json file")

func TestParse(t *testing.T) {
	torrent, err := Parse("testdata/archlinux-2019.12.01-x86_64.iso.torrent")
	require.Nil(t, err)

	jsonPath := "testdata/archlinux-2019.12.01-x86_64.iso.torrent.json"
	if *update {
		serialized, err := json.MarshalIndent(torrent, "", "  ")
		require.Nil(t, err)
		ioutil.WriteFile(jsonPath, serialized, 0644)
	}

	expected := TorrentFile{}
	jsonFile, err := ioutil.ReadFile(jsonPath)
	require.Nil(t, err)
	err = json.Unmarshal(jsonFile, &expected)
	require.Nil(t, err)

	assert.Equal(t, expected, *torrent)
}
