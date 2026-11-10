// Package udpconn implements the udp connection for announce with a udp scheme.
package udpconn

import (
	"bytes"
	"crypto/rand"
	"encoding/binary"
	"net"
	"net/url"
	"time"

	"github.com/official-taufiq/LeechFree/torrent"
)

func Connect(u *url.URL, t *torrent.MetaInfo) error {
	const protocolID uint64 = 0x41727101980

	udpAddr, err := net.ResolveUDPAddr("udp", u.String())
	if err != nil {
		return err
	}
	UDPconn, err := net.DialUDP("udp", nil, udpAddr)
	if err != nil {
		return err
	}
	defer UDPconn.Close()

	if err := UDPconn.SetDeadline(time.Now().Add(time.Second * 10)); err != nil {
		return err
	}

	txID := new(uint32)
	if err = binary.Read(rand.Reader, binary.BigEndian, txID); err != nil {
		return err
	}

	var req bytes.Buffer
	_ = binary.Write(&req, binary.BigEndian, protocolID)
	_ = binary.Write(&req, binary.BigEndian, uint32(0))
	_ = binary.Write(&req, binary.BigEndian, *txID)

	if _, err = UDPconn.Write(req.Bytes()); err != nil {
		return err
	}
	resp := make([]byte, 16)
	if _, err = UDPconn.Read(resp); err != nil {
		return err
	}

	return nil
}
