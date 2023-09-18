package gltf

//Image data used to create a texture. Image **MAY** be referenced by an URI (or IRI) or a buffer view index.
type Image struct {
	ChildOfRootProperty
	URI        string `json:"uri"`                //The URI (or IRI) of the image.  Relative paths are relative to the current glTF asset.  Instead of referencing an external file, this field **MAY** contain a `data:`-URI. This field **MUST NOT** be defined when `bufferView` is defined.
	MimeType   string `json:"mimeType,omitempty"` //The image's media type. This field **MUST** be defined when `bufferView` is defined.
	BufferView Id     `json:"bufferView"`         //The index of the bufferView that contains the image. This field **MUST NOT** be defined when `uri` is defined.
}
