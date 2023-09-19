package gltf

// A typed view into a buffer view that contains raw binary data.
type Accessor struct {
	ChildOfRootProperty
	BufferView    *uint32         `json:"bufferView,omitempty"` //The index of the buffer view. When undefined, the accessor **MUST** be initialized with zeros; `sparse` property or extensions **MAY** override zeros with actual values.
	ByteOffset    uint32          `json:"byteOffset,omitempty"` //The offset relative to the start of the buffer view in bytes.
	ComponentType ComponentType   `json:"componentType"`        //The datatype of the accessor's components.  UNSIGNED_INT type **MUST NOT** be used for any accessor that is not referenced by `mesh.primitive.indices`.
	Normalized    bool            `json:"normalized,omitempty"` //Specifies whether integer data values are normalized (`true`) to [0, 1] (for unsigned types) or to [-1, 1] (for signed types) when they are accessed. This property **MUST NOT** be set to `true` for accessors with `FLOAT` or `UNSIGNED_INT` component type.
	Count         uint32          `json:"count"`                //The number of elements referenced by this accessor, not to be confused with the number of bytes or number of components.
	Type          AccessorType    `json:"type"`                 //Specifies if the accessor's elements are scalars, vectors, or matrices.
	Max           []float64       `json:"max,omitempty"`        //Maximum value of each component in this accessor.  Array elements **MUST** be treated as having the same data type as accessor's `componentType`. Both `min` and `max` arrays have the same length.  The length is determined by the value of the `type` property; it can be 1, 2, 3, 4, 9, or 16.\n\n`normalized` property has no effect on array values: they always correspond to the actual values stored in the buffer. When the accessor is sparse, this property **MUST** contain maximum values of accessor data with sparse substitution applied.
	Min           []float64       `json:"min,omitempty"`        //Minimum value of each component in this accessor.  Array elements **MUST** be treated as having the same data type as accessor's `componentType`. Both `min` and `max` arrays have the same length.  The length is determined by the value of the `type` property; it can be 1, 2, 3, 4, 9, or 16.\n\n`normalized` property has no effect on array values: they always correspond to the actual values stored in the buffer. When the accessor is sparse, this property **MUST** contain minimum values of accessor data with sparse substitution applied.
	Sparse        *AccessorSparse `json:"sparse,omitempty"`     //Sparse storage of elements that deviate from their initialization value.
}

// Sparse storage of accessor values that deviate from their initialization value.
type AccessorSparse struct {
	Property
	Count   uint32                `json:"count"` //Number of deviating accessor values stored in the sparse array.
	Indices AccessorSparseIndices `json:"indices"`
	Values  AccessorSparseValues  `json:"values"`
}

// An object pointing to a buffer view containing the indices of deviating accessor values. The number of indices is equal to `accessor.sparse.count`. Indices **MUST** strictly increase.
type AccessorSparseIndices struct {
	Property
	BufferView    uint32        `json:"bufferView"`           // The index of the buffer view with sparse indices. The referenced buffer view **MUST NOT** have its `target` or `byteStride` properties defined. The buffer view and the optional `byteOffset` **MUST** be aligned to the `componentType` byte length.
	ByteOffset    uint32        `json:"byteOffset,omitempty"` //The offset relative to the start of the buffer view in bytes.
	ComponentType ComponentType `json:"componentType"`        //The indices data type.
}

// An object pointing to a buffer view containing the deviating accessor values. The number of elements is equal to `accessor.sparse.count` times number of components. The elements have the same component type as the base accessor. The elements are tightly packed. Data **MUST** be aligned following the same rules as the base accessor.
type AccessorSparseValues struct {
	Property
	BufferView uint32 `json:"bufferView,omitempty"` //The index of the bufferView with sparse values. The referenced buffer view **MUST NOT** have its `target` or `byteStride` properties defined.
	ByteOffset uint32 `json:"byteOffset,omitempty"` //The offset relative to the start of the bufferView in bytes.
}
