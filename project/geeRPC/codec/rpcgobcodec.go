package codec

import (
	"bytes"
	"encoding/gob"
)

type GobCodec struct {
	enc *gob.Encoder
	dec *gob.Decoder
}

func (c *GobCodec) Encode(obj interface{}) ([]byte, error) {
	var buffer bytes.Buffer
	encode := gob.NewEncoder(&buffer)
	err := encode.Encode(obj)
	if err != nil {
		return nil, err
	}
	return buffer.Bytes(), nil
}

func (c *GobCodec) Decode(data []byte, obj interface{}) error {
	buffer := bytes.NewBuffer(data)
	decode := gob.NewDecoder(buffer)
	return decode.Decode(obj)
}
