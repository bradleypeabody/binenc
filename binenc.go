/*
	Easily implement binary encodings and JSON representation in Go using struct tags and marshaling.

	The binenc package aims to provide: A struct tag format capable of expressing common idioms in binary formats, such as bit flags, endianness, length+value pairs and more. An interface to override binenc's behavior should it fall short. A JSON encoder and decoder providing a human readable representation of the binary data, for debugging and/or storage.

	Struct Tags

	Endianness:

		type Packet struct {
			NumA uint64 `binenc: "le"` // littleendian
			NumB uint64 `binenc: "be"` // big endian
		}

*/
package binenc
