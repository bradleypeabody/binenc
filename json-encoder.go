package binenc

import "io"

type JSONEncoder struct {
	w io.Writer
	c *Config
}

func (c *Config) NewJSONEncoder(w io.Writer) *JSONEncoder {
	return &JSONEncoder{w: w, c: c}
}

func (e *JSONEncoder) Encode(v interface{}) error {

	// TypeConfigForValue(v)

	// loop over fields
	//

	return nil
}

type JSONDecoder struct {
	r io.Reader
	c *Config
}

func (c *Config) NewJSONDecoder(r io.Reader) *JSONDecoder {
	return &JSONDecoder{r: r, c: c}
}

func (d *JSONDecoder) Decode(v interface{}) error {
	return nil
}

// BinencJSON is implemented by types that want to override the JSON encoding/decoding behavior.
type BinencJSON interface {
	BinencFromJSON(c *Config, b []byte) error
	BinencToJSON(c *Config) ([]byte, error)
}
