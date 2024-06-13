package main

import (
	"bytes"
	"encoding/binary"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

var (
	timeout int64
	size    int
	count   int
	typ     uint8 = 8
	code    uint8 = 8
)

type ICMP struct {
	Type        uint8
	Code        uint8
	CheckSum    uint16
	ID          uint16
	SequenceNum uint16
}

func main() {
	fmt.Println("hello go ping")
	getCommandArgs()
	fmt.Println(timeout, size, count)

	desIp := os.Args[len(os.Args)-1]
	conn, err := net.DialTimeout("ip:icmp", desIp, time.Duration(timeout)*time.Millisecond)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer conn.Close()

	fmt.Printf("正在Ping %s [%s] 具有 %d 字节的数据:\n", desIp, conn.RemoteAddr(), size)
	for i := 0; i <= count; i++ {
		t1 := time.Now()
		icmp := &ICMP{
			Type:        typ,
			Code:        code,
			CheckSum:    0,
			ID:          1,
			SequenceNum: 1,
		}

		data := make([]byte, size)
		var buffer bytes.Buffer
		binary.Write(&buffer, binary.BigEndian, icmp)
		buffer.Write(data)
		data = buffer.Bytes()
		checkSum := checkSum(data)
		data[2] = byte(checkSum >> 8)
		data[3] = byte(checkSum)
		conn.SetDeadline(time.Now().Add(time.Duration(timeout) * time.Millisecond))
		n, err := conn.Write(data)
		if err != nil {
			log.Println(err)
			return
		}
		buf := make([]byte, 65535)
		n, err = conn.Read(buf)
		if err != nil {
			log.Println(err)
			return
		}
		ts := time.Since(t1).Milliseconds()
		fmt.Printf("来自%d.%d.%d.%d的回复: 字节=%d, 时间=%d ms TTL=%d\n",
			buf[12], buf[13], buf[14], buf[15], n-28, ts, buf[8])
		time.Sleep(time.Second)
	}
}

func getCommandArgs() {
	flag.Int64Var(&timeout, "w", 1000, "请求超时时长,单位毫秒")
	flag.IntVar(&size, "l", 32, "请求发送缓冲区大小，单位字节")
	flag.IntVar(&count, "n", 4, "发送请求数")
	flag.Parse()
}

func checkSum(data []byte) uint16 {
	length := len(data)
	index := 0
	var sum uint32 = 0
	for length > 1 {
		sum += uint32(data[index])<<8 + uint32(data[index+1])
		length -= 2
		index += 2
	}
	if length != 0 {
		sum += uint32(data[index])
	}
	hi16 := sum >> 16
	for hi16 != 0 {
		sum = hi16 + uint32(uint16(sum))
		hi16 = sum >> 16
	}
	return uint16(^sum)
}
