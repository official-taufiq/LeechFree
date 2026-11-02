package main

import (
	"errors"
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
	fmt.Println("result of string parser")
	fmt.Println(res)
	fmt.Println(string(remaining))

	byteInt := []byte{'i', '4', '2', 'e', '5', ':', 'h', 'e', 'l', 'l', 'o'}
	resInt, remainingInt, err := integerParser(byteInt)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("result integer Parser")
	fmt.Println(resInt)
	fmt.Println(string(remainingInt))
}

// func bencodeUnmarshaller(b []byte) (string, []byte) {
//	data := map[string]string{}
// }

func integerParser(b []byte) (string, []byte, error) {
	for i, v := range b {
		if v == 'e' {
			return string(b[1:i]), b[i+1:], nil
		}
	}

	return "", nil, errors.New("malformed bencode")
}

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
	return "", nil, errors.New("malformed bencode")
}
