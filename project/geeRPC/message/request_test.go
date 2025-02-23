package message

import (
	"testing"
)

func TestEncodeRequest(t *testing.T) {
	req := &Request{
		MessageId:   1,
		Version:     0,
		Serializer:  0,
		Compressor:  0,
		Ping:        0,
		ServiceName: "test",
		MethodName:  "test",
	}
	req.CalcHeaderBodyLength()

	data, err := EncodeRequest(req)
	if err != nil {
		t.Error(err)
	}
	t.Log(data)
}

func TestDecodeRequest(t *testing.T) {
	req := &Request{
		MessageId:   1,
		Version:     0,
		Serializer:  0,
		Compressor:  0,
		Ping:        0,
		ServiceName: "test",
		MethodName:  "test",
	}
	req.CalcHeaderBodyLength()

	data, err := EncodeRequest(req)
	if err != nil {
		t.Error(err)
	}
	t.Log(data)

	drep, err := DecodeRequest(data)
	if err != nil {
		t.Error(err)
	}
	t.Log(drep)
}
