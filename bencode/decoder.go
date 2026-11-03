// Package bencode unmarshalls bencoded data from .torrent files.
package bencode

import (
	"errors"
	"fmt"
	"strconv"
)

var errMalformed = errors.New("marlformed bencode")

func Unmarshal(b []byte) (interface{}, error) {
	if len(b) == 0 {
		return nil, fmt.Errorf("%w: empty torrent, no data found", errMalformed)
	}
	val, rest, err := parse(b)
	if err != nil {
		return nil, err
	}
	if len(rest) != 0 {
		return nil, fmt.Errorf("%w: trailing data", errMalformed)
	}
	return val, nil
}

func parse(b []byte) (interface{}, []byte, error) {
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
	return nil, nil, fmt.Errorf("%w: unexpected byte %q", errMalformed, b[0])
}

func dictParser(b []byte) (map[string]interface{}, []byte, error) {
	dict := make(map[string]interface{})
	rest := b[1:]
	for {
		if len(rest) == 0 {
			return nil, nil, fmt.Errorf("%w: unterminated dictionary", errMalformed)
		}
		if rest[0] == 'e' {
			return dict, rest[1:], nil
		}
		key, remaining1, err := stringParser(rest[1:])
		if err != nil {
			return nil, nil, err
		}
		val, remaining2, err := parse(remaining1)
		if err != nil {
			return nil, nil, err
		}
		dict[key] = val
		rest = remaining2
	}
}

func integerParser(b []byte) (int64, []byte, error) {
	for i, v := range b {
		if v == 'e' {
			if i < 2 {
				return 0, nil, fmt.Errorf("%w: bad integer", errMalformed)
			}
			res, err := strconv.ParseInt(string(b[1:i]), 10, 64)
			if err != nil {
				return 0, nil, fmt.Errorf("%w: %w", errMalformed, err)
			}
			return res, b[i+1:], nil
		}
	}

	return 0, nil, errMalformed
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
	return "", nil, errMalformed
}

func listParser(b []byte) ([]interface{}, []byte, error) {
	val := []interface{}{}
	rest := b[1:]
	for {
		if len(rest) == 0 {
			return nil, nil, fmt.Errorf("%w: unterminated list", errMalformed)
		}
		if rest[0] == 'e' {
			return val, rest[1:], nil
		}
		returnedVal, remaining, err := parse(rest[1:])
		if err != nil {
			return nil, nil, err
		}
		val = append(val, returnedVal)
		rest = remaining
	}
}
