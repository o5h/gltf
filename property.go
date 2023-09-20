package gltf

// glTF Property
type Property struct {
	Extensions Extensions `json:"extensions,omitempty"`
	Extras     Extras     `json:"extras,omitempty"`
}

// glTF Child of Root Property
type ChildOfRootProperty struct {
	Property
	Name string `json:"name,omitempty"` //The user-defined name of this object.  This is not necessarily unique, e.g., an accessor and a buffer could have the same name, or two accessors could even have the same name.
	Root *GLTF  `json:"-"`
}

type ChildOfRoot interface {
	SetRoot(*GLTF)
	GetRoot() *GLTF
}

//GetRoot returns the root of GLTF document
func (c *ChildOfRootProperty) GetRoot() *GLTF { return c.Root }

//SetRoot sets the root of GLTF document
func (c *ChildOfRootProperty) SetRoot(root *GLTF) { c.Root = root }
