// Package torrent maps the unmarshalled bencode data into a struct.
package torrent

import (
	"crypto/sha1"
	"errors"
)

type File struct {
	Length int64
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
type MetaInfo struct {
	Info         Info
	Announce     string
	AnnounceList [][]string
	CreationDate int64
	Comment      string
	CreatedBy    string
	Encoding     string
	InfoHash     [20]byte
}

func getString(m map[string]interface{}, s string) string {
	v, _ := m[s].(string)
	return v
}

func getInt(m map[string]interface{}, s string) int64 {
	v, _ := m[s].(int64)
	return v
}

func Parse(v interface{}, infoHash []byte) (*MetaInfo, error) {
	root, ok := v.(map[string]interface{})
	if !ok {
		return nil, errors.New("root is not a dict")
	}

	info, ok := root["info"].(map[string]interface{})
	if !ok {
		return nil, errors.New("no info dict found")
	}
	t := &MetaInfo{
		Announce:     getString(root, "announce"),
		CreationDate: getInt(root, "creation date"),
		Comment:      getString(root, "comment"),
		CreatedBy:    getString(root, "created by"),
		Encoding:     getString(root, "encoding"),
		Info: Info{
			PieceLength: getInt(info, "piece length"),
			Pieces:      []byte(getString(info, "pieces")),
			Private:     getInt(info, "private"),
			Name:        getString(info, "name"),
			Length:      getInt(info, "length"),
		},
	}

	if tiers, ok := root["announce-list"].([]interface{}); ok {
		for _, tier := range tiers {
			var urls []string

			if links, ok := tier.([]interface{}); ok {
				for _, link := range links {
					if l, ok := link.(string); ok {
						urls = append(urls, l)
					}
				}
			}
			t.AnnounceList = append(t.AnnounceList, urls)
		}
	}

	if files, ok := info["files"].([]interface{}); ok {
		for _, file := range files {
			fm, ok := file.(map[string]interface{})
			if !ok {
				return nil, errors.New("bad file entry")
			}
			f := File{Length: getInt(fm, "length")}

			if p, ok := fm["path"].([]interface{}); ok {
				for _, seg := range p {
					if s, ok := seg.(string); ok {
						f.Path = append(f.Path, s)
					}
				}
			}
			t.Info.Files = append(t.Info.Files, f)
		}
	}
	t.InfoHash = sha1.Sum(infoHash)
	return t, nil
}
