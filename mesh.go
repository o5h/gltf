package gltf

//A set of primitives to be rendered.  Its global transform is defined by a node that references it.
type Mesh struct {
	ChildOfRootProperty
	Primitives []MeshPrimitive `json:"primitives"`
	Weights    []float64       `json:"weights"` //Array of weights to be applied to the morph targets. The number of array elements **MUST** match the number of morph targets.
}

//Geometry to be rendered with the given material.
type MeshPrimitive struct {
	Property
	Attributes MeshPrimitiveAttributes `json:"attributes"` //A plain JSON object, where each key corresponds to a mesh attribute semantic and each value is the index of the accessor containing attribute's data.
	Indices    Id                      `json:"indices"`    //The index of the accessor that contains the vertex indices.  When this is undefined, the primitive defines non-indexed geometry.  When defined, the accessor **MUST** have `SCALAR` type and an unsigned integer component type.
	Material   Id                      `json:"material"`   //The index of the material to apply to this primitive when rendering.
	Mode       MeshPrimitiveMode       `json:"mode"`       //The topology type of primitives to render.
	Targets    MeshPrimitiveAttributes `json:"targets"`    //A plain JSON object specifying attributes displacements in a morph target, where each key corresponds to one of the three supported attribute semantic (`POSITION`, `NORMAL`, or `TANGENT`) and each value is the index of the accessor containing the attribute displacements' data.
}

type MeshPrimitiveAttribute string
type MeshPrimitiveAttributes map[MeshPrimitiveAttribute]uint32
