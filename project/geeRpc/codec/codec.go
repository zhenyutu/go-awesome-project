package codec

type Header struct {
	seq    int64
	method string
}

type Codec interface {
	//序列化
	Encode(interface{}) error
	//反序列化
	Decode(interface{}) error
}
