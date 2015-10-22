package main

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"net"
)

type PacketHeader struct {
	StunMessageType uint16
	MessageLength   uint16
	MagicCookie     uint32
	TransactionId   [12]byte
}

type Packet struct {
	header PacketHeader
	//body   []byte
}

func UnserializePacket(buf *bytes.Buffer) *Packet {
	p := new(Packet)
	headers := bytes.NewBuffer(buf.Bytes()[0:20])
	binary.Read(headers, binary.BigEndian, &p.header)

	return p
}

func CreateBindingPacket() *Packet {
	p := new(Packet)
	p.header.StunMessageType = 0x1
	p.header.MessageLength = 0x0
	p.header.MagicCookie = 0x2112A442
	//p.header.TransactionId =
	return p
}

func (p *Packet) Serialize() []byte {
	var bin_buf bytes.Buffer
	binary.Write(&bin_buf, binary.BigEndian, p.header)
	return bin_buf.Bytes()
}

func (p *Packet) Debug() {
	b := p.Serialize()
	fmt.Print("Binary: \n", hex.Dump(b), "\n")
}

func check_err(err error) {
	if err != nil {
		fmt.Println("Error: ", err)
	}
}

func main() {
	serverAddr, err := net.ResolveUDPAddr("udp", "stun.ideasip.com:3478")
	check_err(err)

	con, err := net.DialUDP("udp", nil, serverAddr)
	check_err(err)

	p := CreateBindingPacket()
	p.Debug()
	con.Write(p.Serialize())

	for {
		buff := make([]byte, 1024)
		ln, err := con.Read(buff)
		check_err(err)
		fmt.Print("Received: \n", hex.Dump(buff[0:ln]), "\n")
		q := UnserializePacket(bytes.NewBuffer(buff[0:ln]))
		q.Debug()
	}

	defer con.Close()
}
