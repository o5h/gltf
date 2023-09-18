package gltf

//Joints and matrices defining a skin.
type Skin struct {
	InverseBindMatrices *Id  `json:"inverseBindMatrices,omitempty"` //The index of the accessor containing the floating-point 4x4 inverse-bind matrices. Its `accessor.count` property **MUST** be greater than or equal to the number of elements of the `joints` array. When undefined, each matrix is a 4x4 identity matrix.
	Skeleton            *Id  `json:"skeleton,omitempty"`            //The index of the node used as a skeleton root. The node **MUST** be the closest common root of the joints hierarchy or a direct or indirect parent node of the closest common root.
	Joints              []Id `json:"joints"`                        //Indices of skeleton nodes, used as joints in this skin.
}
