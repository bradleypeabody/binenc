package binenc

import (
	"encoding/binary"
	"encoding/hex"
	"net/url"
	"reflect"
	"strings"
)

/*
NOTES:

binenc: "le"
binenc: "be"
binenc: "le,const=0xFFFF"
binenc: "strz"
binenc: "mask=WHAT:0xFF,mask=..."
binenc: "mul=0.1,unit=kg" (maybe "optsuffix" or something - the kg doesn't mean anything really it's just for readability)
binenc: "count=FieldName"
binenc: "len=16,fillright=0x20" // or fillleft=0x00 or whatever

binenc: "typemask=FieldName"
binenc: "typemap=CustomMapper"

TODO: we need some sort of pluggable behavior for when something totally custom must happen for a field

--

And can we code generate additional methods for these types for conversion, setting from string, etc?
Could be used to generate JSON marshal stuff also (actually probably we should implmeent the mapstructure
marshaling...)

Probably we should break out the field stuff into some sort of configuration and auto populate
from the struct tags - this way it can be overrridden for cases where there are like 100 flags
and you don't want to put them all in the struct tag - or to allow complex behaviors with things
like callback functions that cannot be expressed in struct tags.

*/

var DEFAULT_BYTE_ORDER binary.ByteOrder = binary.BigEndian

type Config struct {
	TypeMap map[string]*TypeConfig
}

type TypeConfig struct {
	Name         string
	Type         reflect.Type
	FieldConfigs []*FieldConfig
}

type FieldConfig struct {
	Name string // name of the field, matching the Go struct
	// TODO: need to add the kind - probably as reflect.Kind and it needs to correspond to uint16 or whatever

	// Ignore    bool             // if true then field is skipped during all processing
	ByteOrder binary.ByteOrder // whenever endianness is needed, use this or default to DEFAULT_BYTE_ORDER
	Const     []byte           // if non-nil indicates a constant byte sequence
	Strz      bool             // if true indicates value is zero-terminated string
	MaskList  MaskList         // if non-zero the bit masks to parse/apply
}

func NewConfig() *Config {
	return &Config{
		TypeMap: make(map[string]*TypeConfig),
	}
}

func TypeNameFor(v interface{}) string {
	vt := reflect.TypeOf(v)
	for vt.Kind() == reflect.Ptr {
		vt = vt.Elem()
	}
	return vt.Name()
}

func NewTypeConfig(origvt reflect.Type) *TypeConfig {
	vt := origvt
	for vt.Kind() == reflect.Ptr {
		vt = vt.Elem()
	}
	tc := &TypeConfig{
		Name: vt.Name(),
		Type: origvt,
	}

	for i := 0; i < vt.NumField(); i++ {
		f := vt.Field(i)
		fc := &FieldConfig{}
		fc.Name = f.Name

		tagVals := ParseTag(f.Tag.Get("binenc"))

		if len(tagVals["-"]) > 0 {
			// fc.Ignore = true
			continue // skip without adding
		}
		fc.ByteOrder = DEFAULT_BYTE_ORDER
		if len(tagVals["le"]) > 0 {
			fc.ByteOrder = binary.LittleEndian
		}
		if len(tagVals["be"]) > 0 {
			fc.ByteOrder = binary.BigEndian
		}
		if tagVals.Get("const") != "" {
			conststr := strings.TrimPrefix(tagVals.Get("const"), "0x")
			constbytes, _ := hex.DecodeString(conststr)
			fc.Const = constbytes
		}
		if len(tagVals["strz"]) > 0 {
			fc.Strz = true
		}
		// if len(tagVals["mask"]) > 0 {
		// }

		tc.FieldConfigs = append(tc.FieldConfigs, fc)
	}

	return tc
}

// func (c *Config) AddTypes(vs ...interface{}) {
// 	for _, v := range vs {
// 		c.AddType(v)
// 	}
// }

// TODO: TypeConfigForValue(v interface{}) *TypeConfig

// hm, poor name - add "Type" but then it's an interface value...
// AddFromValue, and AddType could take a reflect.Type
// func (c *Config) AddType(v interface{}) {
// 	c.AddTypeConfig(NewTypeConfig(v))
// }

func (c *Config) AddTypeConfig(tc *TypeConfig) {
	c.TypeMap[tc.Name] = tc
}

type Mask struct {
	Name  string
	Value uint64
}

type MaskList []Mask

// func (ml MaskList) jfskdl() uint64 {
// }

// ParseTag splits a struct tag into a url.Values, using commas to separate entries,
// and an equal sign to separate name and value.
func ParseTag(t string) url.Values {

	ps := strings.Split(t, ",")
	v := make(url.Values, len(ps))
	for _, p := range ps {
		kv := strings.SplitN(p, "=", 2)
		if len(kv) < 2 {
			kv = append(kv, "")
		}
		v.Add(kv[0], kv[1])
	}

	return v
}
