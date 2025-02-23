package message

import (
	"encoding/binary"
	"errors"
	"strings"
)

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
	ErrorInfo string
	// 消息体
	Data []byte // 不要用interface，interface不知道类型，所以序列化之后是一个map[string]interface类型
}

func EncodeResponse(resp *Response) ([]byte, error) {
	if resp.HeaderLength == 0 {
		return nil, errors.New("invalid response header info")
	}
	data := make([]byte, resp.HeaderLength+resp.BodyLength)
	cur := data

	binary.BigEndian.PutUint32(cur, resp.HeaderLength)
	cur = cur[4:]

	binary.BigEndian.PutUint32(cur, resp.BodyLength)
	cur = cur[4:]

	binary.BigEndian.PutUint32(cur, resp.MessageId)
	cur = cur[4:]

	cur[0] = resp.Version
	cur = cur[1:]

	cur[0] = resp.Serializer
	cur = cur[1:]

	cur[0] = resp.Compressor
	cur = cur[1:]

	cur[0] = resp.Ping
	cur = cur[1:]

	if resp.ErrorInfo != "" {
		copy(cur, resp.ErrorInfo)
		cur = cur[len(resp.ErrorInfo):]
	}

	copy(cur, resp.Data)
	return data, nil
}

func DecodeResponse(data []byte) (*Response, error) {
	resp := &Response{}

	resp.HeaderLength = binary.BigEndian.Uint32(data[0:4])
	resp.BodyLength = binary.BigEndian.Uint32(data[4:8])
	resp.MessageId = binary.BigEndian.Uint32(data[8:12])
	resp.Version = data[12]
	resp.Serializer = data[13]
	resp.Compressor = data[14]
	resp.Ping = data[15]

	//剩余头部数据
	remainingHeader := data[16:resp.HeaderLength]
	errInfo := strings.TrimSpace(string(remainingHeader))
	resp.ErrorInfo = errInfo

	dataData := make([]byte, resp.BodyLength)
	copy(dataData, data[resp.HeaderLength:])
	resp.Data = dataData

	return resp, nil
}

func (resp *Response) CalcHeaderBodyLength() {
	if len(resp.ErrorInfo) != 0 {
		length := 16 + len(resp.ErrorInfo)
		resp.HeaderLength = uint32(length)
	} else {
		resp.HeaderLength = 16
	}

	if len(resp.Data) != 0 {
		resp.BodyLength = uint32(len(resp.Data))
	}
}
