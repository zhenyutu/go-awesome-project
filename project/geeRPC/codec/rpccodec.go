package codec

type Type string

const (
	Json_Codec = "0"
	Gob_Codec  = "1"
)

var codecs map[string]Codec

type Header struct {
	seq    int64
	method string
}

type Codec interface {
	//序列化
	Encode(obj interface{}) ([]byte, error)
	//反序列化
	Decode(data []byte, obj interface{}) error
}

func init() {
	codecs = make(map[string]Codec)
	codecs[Gob_Codec] = &GobCodec{}
	codecs[Json_Codec] = &JsonCodec{}
}

func GetCodec(name string) Codec {
	if f, ok := codecs[name]; ok {
		return f
	}
	return nil
}
