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
	byteString := []byte("5:helloi42e")
	res, remaining, err := stringParser(byteString)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("result of string parser")
	fmt.Println(res)
	fmt.Println(string(remaining))

	byteInt := []byte("i42e5:hello")
	resInt, remainingInt, err := integerParser(byteInt)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("result integer Parser")
	fmt.Println(resInt)
	fmt.Println(string(remainingInt))
}

func bencodeUnmarshaller(b []byte) (interface{}, []byte, error) {
	switch firstChar := b[0]; {
	case firstChar == 'l':
		return listParser(b)
	case firstChar >= '0' && firstChar <= '9':
		return stringParser(b)
	case firstChar == 'i':
		return integerParser(b)
	case firstChar == 'd':
		return dictParser(b)
	}
	return nil, nil, errors.New("malformed benconde")
}

func dictParser(b []byte) (interface{}, []byte, error) {
	return nil, nil, errors.New("malformed bencode")
}

func integerParser(b []byte) (int, []byte, error) {
	for i, v := range b {
		if v == 'e' {
			res, _ := strconv.Atoi(string(b[1:i]))
			return res, b[i+1:], nil
		}
	}

	return 0, nil, errors.New("malformed bencode")
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

func listParser(b []byte) ([]interface{}, []byte, error) {
	res := []interface{}{}

	for {
		returnedRes, remaining, err := bencodeUnmarshaller(b)
		if err != nil {
			return nil, nil, err
		}
		res = append(res, returnedRes)
		if remaining[0] == 'e' {
			return res, remaining[1:], nil
		}
	}
}
