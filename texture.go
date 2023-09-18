package gltf

//A texture and its sampler.
type Texture struct {
	ChildOfRootProperty
	Sampler *Id `json:"sampler,omitempty"` //The index of the sampler used by this texture. When undefined, a sampler with repeat wrapping and auto filtering **SHOULD** be used.
	Source  *Id `json:"source,omitempty"`  //The index of the image used by this texture. When undefined, an extension or other mechanism **SHOULD** supply an alternate texture source, otherwise behavior is undefined.
}
