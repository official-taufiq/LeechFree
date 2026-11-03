// Package torrent maps the unmarshalled bencode data into a struct.
package torrent

type File struct {
	Length string
	Path   []string
}
type Info struct {
	PieceLength int64
	Pieces      []byte
	Private     int64
	Name        string
	Length      int64  // single-file mode
	Files       []File // multi-file mode
}
type TorrentData struct {
	Info         Info
	Announce     string
	AnnounceList [][]string
	CreationDate int64
	Comment      string
	CreatedBy    string
	Encoding     string
}

func Parse(i interface{}) (*TorrentData, error) {
	t := &TorrentData{
		Announce: "www.kyahorhahai.com",
	}
	return t, nil
}
