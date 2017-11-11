/*
	Easily implement binary encodings and JSON representation in Go using struct tags and marshaling.

	The binenc package aims to provide: A struct tag format capable of expressing common idioms in binary formats, such as bit flags, endianness, length+value pairs and more. An interface to override binenc's behavior should it fall short. A JSON encoder and decoder providing a human readable representation of the binary data, for debugging and/or storage.

	Basic Usage

	TODO - need to show a type, registering it, encoding, parsing. (maybe JSON but not essential, could just mention name of encoder)

	Struct Tags

	Format:

	Struct tags for binenc have values that are separated by commas and can be a single value or a name-value pair
	separated by an equal sign.  Multiple values are allowed.  Examples:

		type Packet struct {
			NumA uint16 `binenc:"le,const=0xF00D"` // has the "le" option set and "const" is set to "0xF00D"
		}

	Endianness:

		type Packet struct {
			NumA uint64 `binenc:"le"` // little endian (e.g. x86, most ARM, etc.)
			NumB uint64 `binenc:"be"` // big endian (network byte order)
		}

	You can also set the default endianess for cases where not tagged by setting DEFAULT_BYTE_ORDER (defaults to
	big endian).

	Constants:

	Headers, footers or other demarcation bytes which have a constant value can be specified with the "const" option
	as a hex encoded string prefixed with "0x".

		type Packet struct {
			// Will always be written as big endian 0xD00D regardless of actual value and parsing will fail unless
			// 0xD00D is read from the stream for this field.
			Head uint16 `binenc:"be,const=0xD00D"`
		}

*/
package binenc
