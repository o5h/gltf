package gltf

//A view into a buffer generally representing a subset of the buffer.
type BufferView struct {
	ChildOfRootProperty
	Buffer uint32     `json:"buffer"`               //The index of the buffer.
	Offset uint32     `json:"byteOffset,omitempty"` //The offset into the buffer in bytes.
	Length uint32     `json:"byteLength"`           //The length of the bufferView in bytes.
	Stride uint8      `json:"byteStride,omitempty"` //The stride, in bytes, between vertex attributes.  When this is not defined, data is tightly packed. When two or more accessors use the same buffer view, this field **MUST** be defined.
	Target BufferType `json:"target,omitempty"`     //The hint representing the intended GPU buffer type to use with this buffer view.
}

func (view *BufferView) Bytes() (bytes []byte) {
	view.Root.GetBuffer(int(view.Buffer)).IfOk(func(b *Buffer) {
		bytes = b.Bytes()
		bytes = bytes[view.Offset : view.Offset+view.Length]
	})
	return
}
