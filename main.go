package main

import (
	"io"
	"log"
	"os"

	"github.com/official-taufiq/LeechFree/bencode"
)

type torrent struct {
	announce string
	info     struct {
		name     string
		infohash []byte
		length   string
		pieces   int
	}
}

func main() {
	file, err := os.Open("../Inglourious Basterds (2009) [2160p] [BluRay] [x265] [10bit] [5.1] [YTS.MX].torrent")
	if err != nil {
		log.Fatal(err)
	}
	b, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	_, _, err = bencode.Unmarshal(b)
	if err != nil {
		log.Fatal(err)
	}
}
