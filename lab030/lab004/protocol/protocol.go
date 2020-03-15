package protocol

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
)

const (
	HEADER_LENGTH = 4
)

func Enpack(message []byte) []byte {
	fmt.Println("发送消息长度:", len(message))
	header := IntToBytes(int32(len(message)))
	return append(header, message...)
}

func Depack(message []byte) ([]byte, error) {
	if len(message) < HEADER_LENGTH {
		return nil, errors.New("包体长度小于理论最小长度")
	}
	fmt.Println("收到协议长度:", len(message))

	mLen := int(BytesToInt(message[:HEADER_LENGTH]))
	fmt.Println("协议头:", mLen)
	if mLen+HEADER_LENGTH != len(message) {
		return nil, errors.New("包体长度错误")
	}
	return message[HEADER_LENGTH:], nil
}

func IntToBytes(i int32) []byte { //这里用int32类型，是因为binary.Write里面是不对int类型操作的
	byteBuffer := bytes.NewBuffer([]byte{})
	binary.Write(byteBuffer, binary.BigEndian, i)

	return byteBuffer.Bytes()
}

func BytesToInt(b []byte) int32 {
	var i int32

	bytesBuffer := bytes.NewBuffer(b)
	binary.Read(bytesBuffer, binary.BigEndian, &i)

	return i
}
