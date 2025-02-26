package client

import (
	"awesomeProject/project/geeRPC/codec"
	"awesomeProject/project/geeRPC/message"
	"awesomeProject/project/geeRPC/service"
	constant "awesomeProject/project/geeRPC/utils"
	"encoding/binary"
	"errors"
	"net"
	"reflect"
	"time"
)

type Client struct {
	addr string
}

func NewClient(addr string) *Client {
	return &Client{addr: addr}
}

func (cl *Client) invoke(serviceName string, serviceMethod string, args []interface{}) ([]byte, error) {
	//初始化请求链接
	conn, err := net.Dial("tcp", cl.addr)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	//入参编码
	codecTool := codec.GetCodec(constant.Codec_type)
	argsData, err := codecTool.Encode(args)
	if err != nil {
		return nil, err
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
		return nil, err
	}
	_, err = conn.Write(data)
	if err != nil {
		return nil, err
	}
	time.Sleep(1000 * time.Millisecond)

	respData, err := cl.ReadResponseData(conn)
	if err != nil {
		return nil, err
	}
	resp, err := message.DecodeResponse(respData)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}

/**
 * 从TCP链接中读取数据
 */
func (s *Client) ReadResponseData(conn net.Conn) ([]byte, error) {
	headerBytes := make([]byte, constant.Header_length)
	length, err := conn.Read(headerBytes)
	if err != nil {
		return nil, err
	}
	if length != constant.Header_length {
		conn.Close()
		return nil, errors.New("invalid header data length")
	}

	headLength := binary.BigEndian.Uint32(headerBytes[:4])
	bodyLength := binary.BigEndian.Uint32(headerBytes[4:constant.Header_length])
	data := make([]byte, headLength+bodyLength)
	n, _ := conn.Read(data[constant.Header_length:])
	if n != int(headLength+bodyLength-constant.Header_length) {
		conn.Close()
		return nil, errors.New("tcp连接未读够全部数据")
	}
	copy(data, headerBytes)

	return data, nil
}

func (c *Client) InitServiceStub(service service.Service) {

	tye := reflect.TypeOf(service).Elem()
	val := reflect.ValueOf(service).Elem()
	for i := 0; i < tye.NumField(); i++ {
		fieldType := tye.Field(i)
		fieldValue := val.Field(i)

		if fieldType.Type.Kind() != reflect.Func {
			continue
		}
		if !fieldValue.CanSet() {
			continue
		}

		fn := reflect.MakeFunc(fieldType.Type, func(args []reflect.Value) (results []reflect.Value) {
			//获取函数和方法名
			serviceName := service.Name()
			methodName := fieldType.Name

			// 处理输入参数
			inArgs := make([]interface{}, 0, len(args))
			for _, arg := range args {
				inArgs = append(inArgs, arg.Interface())
			}

			outType := fieldType.Type.Out(0)
			first := reflect.New(outType).Interface()
			data, _ := c.invoke(serviceName, methodName, inArgs)
			if len(data) > 0 {
				codecTool := codec.GetCodec(constant.Codec_type)
				err := codecTool.Decode(data, first)
				if err != nil {
					results = append(results, reflect.Zero(outType))
					//results = append(results, reflect.ValueOf(fmt.Sprintf("%s%v", "decode response body failed, err: ", err)))
					return
				}
			}

			fr := reflect.ValueOf(first).Elem().Interface()
			results = append(results, reflect.ValueOf(fr))

			//if err != nil {
			//	results = append(results, reflect.ValueOf(errors.New(err.Error())))
			//} else {
			//	results = append(results, reflect.Zero(reflect.TypeOf(new(error))))
			//}
			return results
		})

		fieldValue.Set(fn)
	}

}
