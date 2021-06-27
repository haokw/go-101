package main

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"io"
	"net"
)

const (
	RAWHEADER_LENGTH int = 16

	PACKET_OFFSET int = 0
	HEADER_OFFSET int = 4
	VER_OFFSET    int = 6
	OP_OFFSET     int = 8
	SEQ_OFFSET    int = 12
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:3101")
	if err != nil {
		fmt.Println("dial failed, err", err)
		return
	}
	defer conn.Close()

	reader := bufio.NewReader(conn)
	for {
		msg, err := Decode(reader)
		if err == io.EOF {
			return
		}
		if err != nil {
			fmt.Println("decode msg failed, err:", err)
			return
		}
		fmt.Println("decode msg:", msg)
	}
}

// Decode 解码消息
func Decode(reader *bufio.Reader) (string, error) {
	// 读取消息头
	rawLenBytes, _ := reader.Peek(RAWHEADER_LENGTH)

	// 解析消息头
	packetLen := binary.LittleEndian.Uint32(rawLenBytes[PACKET_OFFSET:])
	headerLen := binary.LittleEndian.Uint16(rawLenBytes[HEADER_OFFSET:])
	ver := binary.LittleEndian.Uint16(rawLenBytes[VER_OFFSET:])
	op := binary.LittleEndian.Uint32(rawLenBytes[OP_OFFSET:])
	seq := binary.LittleEndian.Uint32(rawLenBytes[SEQ_OFFSET:])

	fmt.Println("[receiveHeader][packetLen:", packetLen, "][headerLen:", headerLen, "][ver:", ver, "][op:", op, "][seq:", seq, "]")

	// Buffered返回缓冲中现有的可读取的字节数
	if reader.Buffered() < int(packetLen) {
		return "", nil
	}

	// 读取真正的消息数据
	pack := make([]byte, int(packetLen))
	_, err := reader.Read(pack)
	if err != nil {
		return "", err
	}
	return string(pack[RAWHEADER_LENGTH:]), nil
}
