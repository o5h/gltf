package gltf

//A texture and its sampler.
type Texture struct {
	ChildOfRootProperty
	Sampler *uint32 `json:"sampler,omitempty"` //The index of the sampler used by this texture. When undefined, a sampler with repeat wrapping and auto filtering **SHOULD** be used.
	Source  *uint32 `json:"source,omitempty"`  //The index of the image used by this texture. When undefined, an extension or other mechanism **SHOULD** supply an alternate texture source, otherwise behavior is undefined.
}
