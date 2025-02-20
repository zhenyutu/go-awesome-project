package server

import (
	"awesomeProject/project/geeRPC/codec"
	"awesomeProject/project/geeRPC/message"
	"encoding/binary"
	"errors"
	"log"
	"net"
)

func (s *Server) ListenAndHandle(addr string) error {
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}

		go func() {
			e := s.HandleRequest(conn)
			if e != nil {
				log.Println(e)
			}
		}()
	}
}

func (s *Server) HandleRequest(conn net.Conn) error {
	defer func() { _ = conn.Close() }()
	var err error

	// 读请求
	data, err := s.ReadRequestData(conn)
	if err != nil {
		return err
	}
	req, err := message.DecodeRequest(data)
	if err != nil {
		return err
	}
	log.Println(req)

	resp := &message.Response{
		MessageId:  req.MessageId,
		Version:    req.Version,
		Serializer: req.Serializer,
		Compressor: req.Compressor,
	}

	//确定编码方式，入参解码
	codecType := string(rune(int(req.Serializer)))
	codecTool := codec.GetCodec(codecType)
	argsIn := make([]interface{}, 0)
	err = codecTool.Decode(req.Data, &argsIn)
	if err != nil {
		return err
	}

	// 执行
	service, ok := s.services[req.ServiceName]
	if !ok {
		resp.ErrorInfo = []byte("service not found")
		resp.CalcHeaderBodyLength()
		err = s.SendResponseData(conn, resp)
		if err != nil {
			return err
		}
	}
	result, err := service.Handle(req.MethodName, argsIn)
	if err != nil {
		return err
	}

	//编码结果
	resultData, err := codecTool.Encode(result)
	if err != nil {
		return err
	}
	resp.Data = resultData
	// 写回响应
	err = s.SendResponseData(conn, resp)
	if err != nil {
		return err
	}
	return nil
}

/**
 * 从TCP链接中读取数据
 */
func (s *Server) ReadRequestData(conn net.Conn) ([]byte, error) {
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

func (s *Server) SendResponseData(conn net.Conn, resp *message.Response) error {
	//编码
	respData, err := message.EncodeResponse(resp)
	if err != nil {
		return err
	}

	//发送
	_, err = conn.Write(respData)
	if err != nil {
		return err
	}

	return nil
}
