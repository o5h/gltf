package gltf

type Asset struct {
	Property
	Copyright  string `json:"copyright,omitempty"`  //A copyright message suitable for display to credit the content creator.
	Generator  string `json:"generator,omitempty"`  //Tool that generated this glTF model.  Useful for debugging.
	Version    string `json:"version"`              //The glTF version in the form of `<major>.<minor>` that this asset targets.
	MinVersion string `json:"minVersion,omitempty"` //The minimum glTF version in the form of `<major>.<minor>` that this asset targets. This property **MUST NOT** be greater than the asset version.
}
