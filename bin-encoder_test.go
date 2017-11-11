package binenc

import (
	"bytes"
	"log"
	"reflect"
	"testing"
)

type T1 struct {
	Header []byte `binenc:"const=0xBABECAFE"`
	Test1  uint16 `binenc:"le"`
	Test2  uint16 `binenc:"be"`
	// PayloadType uint16      `binenc:"mask=T1A:0x01,mask=T1B:0x02"`
	// PayloadData interface{} `binenc:"typemask=PayloadType"`
}

type T1A struct {
}

type T1B struct {
}

func TestBinenc(t *testing.T) {

	config := NewConfig()
	config.AddTypeConfig(NewTypeConfig(reflect.TypeOf(&T1{})))

	t1 := T1{}
	t1.Test1 = 0x1234
	t1.Test2 = 0x1234

	var buf bytes.Buffer

	enc := config.NewBinEncoder(&buf)
	err := enc.Encode(t1)
	if err != nil {
		t.Fatal(err)
	}

	log.Printf("b=%x", buf.Bytes())

}
