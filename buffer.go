package gltf

import (
	"encoding/base64"
	"strings"
)

// A buffer points to binary geometry, animation, or skins.
type Buffer struct {
	ChildOfRootProperty
	URI        string `json:"uri,omitempty"` //The URI (or IRI) of the buffer.  Relative paths are relative to the current glTF asset.  Instead of referencing an external file, this field **MAY** contain a `data:`-URI.
	ByteLength uint32 `json:"byteLength"`    //The length of the buffer in bytes.
}

func (buf *Buffer) IsEmbedded() bool { return strings.HasPrefix(buf.URI, octet_stream_base64_prefix) }

func (buf *Buffer) Bytes() []byte {
	if buf.IsEmbedded() {
		b, _ := base64.StdEncoding.DecodeString(buf.URI[len(octet_stream_base64_prefix):])
		return b
	}
	return nil
}
