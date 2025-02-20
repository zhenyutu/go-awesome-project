package codec

import (
	"bytes"
	"encoding/json"
)

type JsonCodec struct {
}

func (c *JsonCodec) Encode(obj interface{}) ([]byte, error) {
	return json.Marshal(obj)
}

func (c *JsonCodec) Decode(data []byte, obj interface{}) error {
	buffer := bytes.NewBuffer(data)
	jsonDecoder := json.NewDecoder(buffer)
	return jsonDecoder.Decode(obj)
	//return json.Unmarshal(data, obj)
}
