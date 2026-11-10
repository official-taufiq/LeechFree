package main

import (
	"io"
	"log"
	"net/url"
	"os"

	"github.com/official-taufiq/LeechFree/bencode"
	"github.com/official-taufiq/LeechFree/httpconn"
	"github.com/official-taufiq/LeechFree/torrent"
	"github.com/official-taufiq/LeechFree/udpconn"
)

func main() {
	file, err := os.Open("../Inglourious Basterds (2009) [2160p] [BluRay] [x265] [10bit] [5.1] [YTS.MX].torrent")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	b, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	bc, infoHash, err := bencode.Unmarshal(b)
	if err != nil {
		log.Fatal(err)
	}

	metaInfo, err := torrent.Parse(bc, infoHash)
	if err != nil {
		log.Fatal(err)
	}

	u, err := url.Parse(metaInfo.Announce)
	if err != nil {
		log.Fatal(err)
	}

	switch u.Scheme {
	case "http":
		err := httpconn.Connect(u, metaInfo)
		if err != nil {
			log.Fatal(err)
		}
	case "https":
		err := httpconn.Connect(u, metaInfo)
		if err != nil {
			log.Fatal(err)
		}
	case "udp":
		err := udpconn.Connect(u, metaInfo)
		if err != nil {
			log.Fatal(err)
		}
	default:
		log.Fatalf("invalid Scheme:%s, only supports http & udp", u.Scheme)
	}
}
