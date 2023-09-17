package gltf

// glTF Property
type Property struct {
	Extensions Extensions `json:"extensions,omitempty"`
	Extras     Extras     `json:"extras,omitempty"`
}

// glTF Child of Root Property
type ChildOfRootProperty struct {
	Name string `json:"name,omitempty"` //The user-defined name of this object.  This is not necessarily unique, e.g., an accessor and a buffer could have the same name, or two accessors could even have the same name.
	Property
}
