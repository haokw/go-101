# week 09 作业

1. 总结几种 socket 粘包的解包方式: fix length / delimiter based / length field based frame decoder。尝试举例其应用

2. 实现一个从 socket connection 中解码出 goim 协议的解码器。

## 解答

### 1

TCP 粘包、拆包，常见的原因

1. 要发送的数据 大于 TCP 发送缓冲区剩余空间大小，将会 拆包
2. 待发送数据 小于 MSS（最大报文长度），TCP 在传输前将进行 拆包
3. 要发送的数据 小于 TCP 发送缓冲区大小，TCP 将多次写入缓冲区的数据一次发送出去，将会发生 粘包
4. 接受数据端的应用层没有及时读取缓冲区数据，将发生粘包

解决办法

1. fix length，发送端将每个数据包封装为固定长度（不够的通过补零填充），这样接收端每次从接收缓冲区中读取固定长度的数据进行拆包
2. delimiter based，在数据包之间设置边界，添加特殊符号，接收端通过这个边界将不同的数据包拆开
3. length field based，发送端将每个数据包添加包首部，首部中应该至少包含数据包的长度，这样接收端在接收到数据后，通过读取包首部的长度字段，通过长度拆包

应用场景

1. fix length，用于对消息实时性要求不高，或者网络比较好的内网环境
2. delimiter based，消息内容固定，用于传输内部指令等
3. length field based，消息长度和内容多变且多样的场景

### 2

参考 examples/javascript/client.js 实现

```go
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
```
