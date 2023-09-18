package gltf

//The root nodes of a scene.
type Scene struct {
	ChildOfRootProperty
	Nodes []Id `json:"nodes,omitempty"` //The indices of each root node.
}
