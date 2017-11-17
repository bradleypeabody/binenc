package binenc

import (
	"fmt"
	"io"
	"reflect"
)

// TODO: in the same style as some of the stdlib packages we could provide global functions
// that operate on a shared config, making the simple cases simpler but still mantaining
// the ability to have as many configs as you want
type BinEncoder struct {
	w io.Writer
	c *Config
}

func (c *Config) NewBinEncoder(w io.Writer) *BinEncoder {
	return &BinEncoder{w: w, c: c}
}

func (e *BinEncoder) Encode(v interface{}) error {

	// if it implements the interface, use it
	if et, ok := v.(BinEncoderTo); ok {
		return et.BinEncodeTo(c, e.w)
	}

	name := TypeNameFor(v)
	if name == "" {
		return fmt.Errorf("unable to determine name for value provided (and it does not implement BinEncoderTo interface): %t", v)
	}
	tc := e.c.TypeMap[name]
	if tc == nil {
		return fmt.Errorf("unable to find TypeConfig for %q (did you forget to register it?)", name)
	}

	// get the "flat type" - struct without pointer
	// flatType := tc.Type
	// for flatType.Kind() == reflect.Ptr {
	// 	flatType = flatType.Elem()
	// }

	vv := reflect.ValueOf(v)
	for vv.Kind() == reflect.Ptr {
		vv = vv.Elem()
	}

	for _, fc := range tc.FieldConfigs {

		// fc.Name
		// fv := fc.Name
		// flatType.FieldByName(fc.Name)
		bo := fc.ByteOrder
		if bo == nil {
			bo = DEFAULT_BYTE_ORDER
		}

		fv := vv.FieldByName(fc.Name)
		// fvv := fv.Interface()

		switch fv.Kind() {

		case reflect.Uint8:

			bv := fv.Interface().(byte)
			_, err := e.w.Write([]byte{bv})
			if err != nil {
				return err
			}
			continue

		case reflect.Uint16:

			outv := fv.Interface().(uint16)
			outb := make([]byte, 2)
			bo.PutUint16(outb, outv)
			_, err := e.w.Write(outb)
			if err != nil {
				return err
			}
			continue

		}

		// fc.Name

	}
	// tc.FieldConfigs

	// e.c.TypeMap

	// TypeConfigForValue(v)

	// loop over fields
	//

	return nil
}

type BinDecoder struct {
	r io.Reader
	c *Config
}

func (c *Config) NewBinDecoder(r io.Reader) *BinDecoder {
	return &BinDecoder{r: r, c: c}
}

func (d *BinDecoder) Decode(v interface{}) error {
	return nil
}

type BinEncoderTo interface {
	BinEncodeTo(c *Config, w io.Writer) error
}
type BinDecoderFrom interface {
	BinDecodeFrom(c *Config, r io.Reader) error
}

// // BinencBytes is implemented by types that want to override the JSON encoding/decoding behavior.
// // hm this is wrong - needs to be structred as reader/writer, since the length of the data may
// // be unknown by caller
// type BinencBytes interface {
// 	BinencFromBytes(c *Config, b []byte) error
// 	BinencToBytes(c *Config) ([]byte, error)
// }
