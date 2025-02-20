package codec

import "testing"

type Person struct {
	Name string
	Age  int
}

func TestCodec(test *testing.T) {
	codec := GetCodec("0")
	p := Person{"john", 20}

	val, err := codec.Encode(p)
	if err != nil {
		test.Error("encode error:", err)
	}
	test.Log(val)

	dp := Person{}
	e := codec.Decode(val, &dp)
	if e != nil {
		test.Error("decode error:", e)
	}
	test.Log(dp)
}
