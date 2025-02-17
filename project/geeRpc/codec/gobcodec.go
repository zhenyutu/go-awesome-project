package codec

import "encoding/gob"

type GobCodec struct {
	enc *gob.Encoder
	dec *gob.Decoder
}

func (c *GobCodec) Encode(clazz interface{}) error {
	return c.enc.Encode(clazz)
}

func (c *GobCodec) Decode(obj interface{}) error {
	return c.dec.Decode(obj)
}
