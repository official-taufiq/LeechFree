package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("../Inglourious Basterds (2009) [2160p] [BluRay] [x265] [10bit] [5.1] [YTS.MX].torrent")
	if err != nil {
		log.Fatal(err)
	}
	_, err = io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	byteString := []byte{'5', ':', 'h', 'e', 'l', 'l', 'o', 'i', '4', '2', 'e'}
	res, remaining, err := stringParser(byteString)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res)
	fmt.Println(string(remaining))
}

// func bencodeUnmarshaller(b []byte) (string, []byte) {
//	data := map[string]string{}
// }

// func integerParser(b []byte) (string, []byte) {
//}

func stringParser(b []byte) (string, []byte, error) {
	for i, v := range b {
		if v == ':' {
			lenInt, err := strconv.Atoi(string(b[:i]))
			if err != nil {
				return "", nil, err
			}
			return string(b[i+1 : i+lenInt+1]), b[i+lenInt+1:], nil
		}
	}
	return "", nil, nil
}
