package gltf

//A node in the node hierarchy.  When the node contains `skin`, all `mesh.primitives` **MUST** contain `JOINTS_0` and `WEIGHTS_0` attributes.  A node **MAY** have either a `matrix` or any combination of `translation`/`rotation`/`scale` (TRS) properties. TRS properties are converted to matrices and postmultiplied in the `T * R * S` order to compose the transformation matrix; first the scale is applied to the vertices, then the rotation, and then the translation. If none are provided, the transform is the identity. When a node is targeted for animation (referenced by an animation.channel.target), `matrix` **MUST NOT** be present.
type Node struct {
	ChildOfRootProperty
	Camera      *uint32     `json:"camera,omitempty"`      //The index of the camera referenced by this node.
	Children    []uint32    `json:"children,omitempty"`    //The indices of this node's children.
	Skin        *uint32     `json:"skin,omitempty"`        //The index of the skin referenced by this node. When a skin is referenced by a node within a scene, all joints used by the skin **MUST** belong to the same scene. When defined, `mesh` **MUST** also be defined.
	Matrix      [16]float64 `json:"matrix,omitempty"`      //A floating-point 4x4 transformation matrix stored in column-major order.
	Mesh        *uint32     `json:"mesh,omitempty"`        //The index of the mesh in this node.
	Rotation    [4]float64  `json:"rotation,omitempty"`    //The node's unit quaternion rotation in the order (x, y, z, w), where w is the scalar.
	Scale       [3]float64  `json:"scale,omitempty"`       //The node's non-uniform scale, given as the scaling factors along the x, y, and z axes.
	Translation [3]float64  `json:"translation,omitempty"` //The node's translation along the x, y, and z axes.
	Weights     []float64   `json:"weights,omitempty"`     //The weights of the instantiated morph target. The number of array elements **MUST** match the number of morph targets of the referenced mesh. When defined, `mesh` **MUST** also be defined.
}
