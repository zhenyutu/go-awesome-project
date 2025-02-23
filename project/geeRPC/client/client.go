package client

import (
	"awesomeProject/project/geeRPC/codec"
	"awesomeProject/project/geeRPC/message"
	"encoding/binary"
	"errors"
	"log"
	"net"
	"time"
)

const header_length = 8

type Client struct {
	addr string
}

func NewClient(addr string) *Client {
	return &Client{addr: addr}
}

func (cl *Client) invoke(serviceName string, serviceMethod string, args interface{}) error {
	//初始化请求链接
	conn, err := net.Dial("tcp", cl.addr)
	if err != nil {
		return err
	}
	defer conn.Close()

	//入参编码
	codecTool := codec.GetCodec("0")
	argsIn := make([]interface{}, 0)
	argsIn = append(argsIn, args)
	argsData, err := codecTool.Encode(argsIn)
	if err != nil {
		return err
	}

	//实例化请求内容
	req := &message.Request{
		MessageId:   1,
		Version:     0,
		Serializer:  0,
		Compressor:  0,
		Ping:        0,
		ServiceName: serviceName,
		MethodName:  serviceMethod,
		Data:        argsData,
	}
	req.CalcHeaderBodyLength()

	data, err := message.EncodeRequest(req)
	if err != nil {
		return err
	}
	_, err = conn.Write(data)
	if err != nil {
		return err
	}
	time.Sleep(1000 * time.Millisecond)

	respData, err := cl.ReadResponseData(conn)
	if err != nil {
		return err
	}
	resp, err := message.DecodeResponse(respData)
	if err != nil {
		return err
	}
	log.Println(req, string(resp.Data))
	return nil
}

/**
 * 从TCP链接中读取数据
 */
func (s *Client) ReadResponseData(conn net.Conn) ([]byte, error) {
	headerBytes := make([]byte, header_length)
	length, err := conn.Read(headerBytes)
	if err != nil {
		return nil, err
	}
	if length != header_length {
		conn.Close()
		return nil, errors.New("invalid header data length")
	}

	headLength := binary.BigEndian.Uint32(headerBytes[:4])
	bodyLength := binary.BigEndian.Uint32(headerBytes[4:header_length])
	data := make([]byte, headLength+bodyLength)
	n, _ := conn.Read(data[header_length:])
	if n != int(headLength+bodyLength-header_length) {
		conn.Close()
		return nil, errors.New("tcp连接未读够全部数据")
	}
	copy(data, headerBytes)

	return data, nil
}
