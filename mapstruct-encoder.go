package binenc

type MapStructEncoder struct {
	c *Config
}

func (c *Config) NewMapStructEncoder() *MapStructEncoder {
	return &MapStructEncoder{c: c}
}

func (e *MapStructEncoder) Encode(bindata interface{}, mapStruct map[string]interface{}) error {
	return nil
}

type MapStructDecoder struct {
	c *Config
}

func (c *Config) NewMapStructDecoder() *MapStructDecoder {
	return &MapStructDecoder{c: c}
}

func (d *MapStructDecoder) Decode(mapStruct map[string]interface{}, bindata interface{}) error {
	return nil
}
