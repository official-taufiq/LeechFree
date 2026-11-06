package main

import (
	"io"
	"log"
	"os"

	"github.com/official-taufiq/LeechFree/bencode"
	"github.com/official-taufiq/LeechFree/torrent"
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

	res, err := bencode.Unmarshal(b)
	if err != nil {
		log.Fatal(err)
	}

	metaInfo, err := torrent.Parse(res)
	if err != nil {
		log.Fatal(err)
	}
}
