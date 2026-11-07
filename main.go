package main

import (
	"crypto/rand"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"time"

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

	bc, infoHash, err := bencode.Unmarshal(b)
	if err != nil {
		log.Fatal(err)
	}

	metaInfo, err := torrent.Parse(bc, infoHash)
	if err != nil {
		log.Fatal(err)
	}

	peerID := make([]byte, 20)
	rand.Read(peerID)

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	u, err := url.Parse(metaInfo.Announce)
	if err != nil {
		log.Fatal(err)
	}

	v := url.Values{}
	v.Add("info_hash", string(metaInfo.InfoHash[:]))
	v.Add("peer_id", string(peerID))
	v.Add("port", "6887")
	v.Add("uploaded", "0")
	v.Add("downloaded", "0")
	v.Add("left", strconv.FormatInt(metaInfo.Info.Length, 10))
	u.RawQuery = v.Encode()
	fmt.Println(u.String())
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		log.Fatal(err)
	}
	_, err = client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
}
