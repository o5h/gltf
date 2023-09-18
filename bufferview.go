package gltf

//A view into a buffer generally representing a subset of the buffer.
type BufferView struct {
	ChildOfRootProperty
	Buffer     Id         `json:"buffer"`               //The index of the buffer.
	ByteOffset uint32     `json:"byteOffset,omitempty"` //The offset into the buffer in bytes.
	ByteLength uint32     `json:"byteLength"`           //The length of the bufferView in bytes.
	ByteStride uint8      `json:"byteStride,omitempty"` //The stride, in bytes, between vertex attributes.  When this is not defined, data is tightly packed. When two or more accessors use the same buffer view, this field **MUST** be defined.
	Target     BufferType `json:"target,omitempty"`     //The hint representing the intended GPU buffer type to use with this buffer view.
}
