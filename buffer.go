package gltf

//A buffer points to binary geometry, animation, or skins.
type Buffer struct {
	ChildOfRootProperty
	URI        string `json:"uri,omitempty"` //The URI (or IRI) of the buffer.  Relative paths are relative to the current glTF asset.  Instead of referencing an external file, this field **MAY** contain a `data:`-URI.
	ByteLength uint32 `json:"byteLength"`    //The length of the buffer in bytes.
}
