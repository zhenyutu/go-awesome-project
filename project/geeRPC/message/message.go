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
	return req, nil
}

func EncodeRequest(req *Request) ([]byte, error) {
	return nil, nil
}

type Response struct {
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
	ErrorInfo []byte
	// 消息体
	Data []byte // 不要用interface，interface不知道类型，所以序列化之后是一个map[string]interface类型
}

func EncodeResponse(resp *Response) ([]byte, error) {
	if resp.HeaderLength == 0 || resp.BodyLength == 0 || resp.MessageId == 0 || resp.Version == 0 {
		return nil, errors.New("invalid response header info")
	}
	data := make([]byte, resp.HeaderLength+resp.BodyLength)
	cur := data

	binary.BigEndian.PutUint32(cur, resp.HeaderLength)
	cur = cur[4:]

	binary.BigEndian.PutUint32(cur, resp.BodyLength)
	cur = cur[4:8]

	binary.BigEndian.PutUint32(cur, resp.MessageId)
	cur = cur[8:12]

	cur[0] = resp.Version
	cur = cur[1:]

	cur[0] = resp.Serializer
	cur = cur[1:]

	cur[0] = resp.Compressor
	cur = cur[1:]

	cur[0] = resp.Ping
	cur = cur[1:]

	return data, nil
}

func (resp *Response) CalculateHeaderLength() {
	if len(resp.ErrorInfo) == 0 {
		return
	}

	length := 16 + len(resp.ErrorInfo)
	resp.HeaderLength = uint32(length)
}
