# binenc: Binary Encoding Library... NVM - found http://kaitai.io/ and instead of pursing this project will try to use Kaitai when it supports Go.
Easily implement binary encodings and JSON representation in Go using struct tags and marshaling.

# But Why

Embedded systems, IoT devices and other scenarios often employ binary encodings which require carefully crafted code to read, write and debug.  As the size of such information increases so does the tedium of implementing the parsing and encoding logic, as well as the likelyhood for bugs to arise.

Thus binenc aims to provide the following:
* A struct tag format capable of expressing common idioms in binary formats, such as bit flags, endianness, length+value pairs and more.
* An interface to override binenc's behavior should it fall short.
* A JSON encoder and decoder providing a human readable representation of the binary data, for debugging and/or storage.

# Usage

TODO: walk through the main features and provide examples - or possibly just one or two examples that each demonstrate mutliple features.

## Endianness

## Constant Values

...


# JSON Encoding/Decoding



# Documentation

More complete [documentation can be found on GoDoc.org](https://godoc.org/github.com/bradleypeabody/binenc).
