package message

import (
	"bytes"
	"encoding/binary"
	"errors"
)

const (
	splitter = '\n'
)

type Request struct {
	// 头部长度
	HeaderLength uint32
	// 消息体长度
	BodyLength uint32
	// 消息ID
	MessageId uint32
	//消息版本
	Version byte
	// 序列化协议
	Serializer byte
	// 压缩算法
	Compressor byte
	// 探活
	Ping byte
	// 服务名
	ServiceName string
	// 方法名
	MethodName string
	//拓展信息
	Extra []byte
	// 消息体
	Data []byte // 不要用interface，interface不知道类型，所以序列化之后是一个map[string]interface类型
}

func DecodeRequest(data []byte) (*Request, error) {
	req := &Request{}

	req.HeaderLength = binary.BigEndian.Uint32(data[0:4])
	req.BodyLength = binary.BigEndian.Uint32(data[4:8])
	req.MessageId = binary.BigEndian.Uint32(data[8:12])
	req.Version = data[12]
	req.Serializer = data[13]
	req.Compressor = data[14]
	req.Ping = data[15]

	//剩余头部数据
	remainingHeader := data[16:req.HeaderLength]
	idx := bytes.IndexByte(remainingHeader, splitter)
	req.ServiceName = string(remainingHeader[:idx])

	remainingHeader = remainingHeader[idx+1:]
	idx = bytes.IndexByte(remainingHeader, splitter)
	req.MethodName = string(remainingHeader[:idx])
	remainingHeader = remainingHeader[idx+1:]

	extra := make([]byte, len(remainingHeader))
	copy(extra, remainingHeader)
	req.Extra = extra

	dataData := make([]byte, req.BodyLength)
	copy(dataData, data[req.HeaderLength:])
	req.Data = dataData
	return req, nil
}

func EncodeRequest(req *Request) ([]byte, error) {
	if req.HeaderLength == 0 {
		return nil, errors.New("header length is zero")
	}
	bytes := make([]byte, req.HeaderLength+req.BodyLength)
	cur := bytes

	binary.BigEndian.PutUint32(cur, req.HeaderLength)
	cur = cur[4:]

	binary.BigEndian.PutUint32(cur, req.BodyLength)
	cur = cur[4:]

	binary.BigEndian.PutUint32(cur, req.MessageId)
	cur = cur[4:]

	cur[0] = req.Version
	cur = cur[1:]

	cur[0] = req.Serializer
	cur = cur[1:]

	cur[0] = req.Compressor
	cur = cur[1:]

	cur[0] = req.Ping
	cur = cur[1:]

	copy(cur, req.ServiceName)
	cur[len(req.ServiceName)] = splitter
	cur = cur[len(req.ServiceName)+1:]

	copy(cur, req.MethodName)
	cur[len(req.MethodName)] = splitter
	cur = cur[len(req.MethodName)+1:]

	if len(req.Extra) > 0 {
		copy(cur, req.Extra)
		cur = cur[len(req.Extra):]
	}

	if len(req.Data) > 0 {
		copy(cur, req.Data)
	}

	return bytes, nil
}

func (req *Request) CalcHeaderBodyLength() {
	if req.ServiceName == "" || req.MethodName == "" {
		panic("service name or method is empty")
	}
	headerLength := 16
	headerLength = headerLength + len(req.ServiceName) + 1
	headerLength = headerLength + len(req.MethodName) + 1
	if len(req.Extra) != 0 {
		headerLength = headerLength + len(req.Extra)
	}
	req.HeaderLength = uint32(headerLength)

	if len(req.Data) != 0 {
		req.BodyLength = uint32(len(req.Data))
	}
}
