// Package httpconn implements http connection for annouce with an http scheme.
package httpconn

import (
	"crypto/rand"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/official-taufiq/LeechFree/torrent"
)

func Connect(u *url.URL, t *torrent.MetaInfo) error {
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	peerID := make([]byte, 20)
	rand.Read(peerID)

	v := url.Values{}
	v.Add("info_hash", string(t.InfoHash[:]))
	v.Add("peer_id", string(peerID))
	v.Add("port", "6887")
	v.Add("uploaded", "0")
	v.Add("downloaded", "0")
	v.Add("left", strconv.FormatInt(t.Info.Length, 10))
	u.RawQuery = v.Encode()

	fmt.Println(u.String())

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	resp.Body.Close()

	fmt.Println(resp.Body)

	return nil
}
