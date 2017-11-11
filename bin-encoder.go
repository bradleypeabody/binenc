package binenc

import (
	"fmt"
	"io"
)

type BinEncoder struct {
	w io.Writer
	c *Config
}

func (c *Config) NewBinEncoder(w io.Writer) *BinEncoder {
	return &BinEncoder{w: w, c: c}
}

func (e *BinEncoder) Encode(v interface{}) error {

	// if it implements the interface, use it
	if bb, ok := v.(BinencBytes); ok {
		b, err := bb.BinencToBytes(e.c)
		if err != nil {
			return err
		}
		_, err = e.w.Write(b)
		return err
	}

	name := TypeNameFor(v)
	if name == "" {
		return fmt.Errorf("unable to determine name for value provided")
	}
	tc := e.c.TypeMap[name]
	if tc == nil {
		return fmt.Errorf("unable to find TypeConfig for %q (did you forget to register it?)", name)
	}

	for _, fc := range tc.FieldConfigs {
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

// BinencBytes is implemented by types that want to override the JSON encoding/decoding behavior.
type BinencBytes interface {
	BinencFromBytes(c *Config, b []byte) error
	BinencToBytes(c *Config) ([]byte, error)
}
