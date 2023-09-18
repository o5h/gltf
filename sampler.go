package gltf

type SamplerParameter uint32

const (
	SamplerParameter_NEAREST                = SamplerParameter(GL_NEAREST)
	SamplerParameter_LINEAR                 = SamplerParameter(GL_LINEAR)
	SamplerParameter_NEAREST_MIPMAP_NEAREST = SamplerParameter(GL_NEAREST_MIPMAP_NEAREST)
	SamplerParameter_LINEAR_MIPMAP_NEAREST  = SamplerParameter(GL_LINEAR_MIPMAP_NEAREST)
	SamplerParameter_NEAREST_MIPMAP_LINEAR  = SamplerParameter(GL_NEAREST_MIPMAP_LINEAR)
	SamplerParameter_LINEAR_MIPMAP_LINEAR   = SamplerParameter(GL_LINEAR_MIPMAP_LINEAR)
	SamplerParameter_CLAMP_TO_EDGE          = SamplerParameter(GL_CLAMP_TO_EDGE)
	SamplerParameter_MIRRORED_REPEAT        = SamplerParameter(GL_MIRRORED_REPEAT)
	SamplerParameter_REPEAT                 = SamplerParameter(GL_REPEAT)
)

// Texture sampler properties for filtering and wrapping modes.
type Sampler struct {
	MagFilter SamplerParameter `json:"magFilter"` //Magnification filter.
	MinFilter SamplerParameter `json:"minFilter"` //Minification filter.
	WrapS     SamplerParameter `json:"wrapS"`     //S (U) wrapping mode.  All valid values correspond to WebGL enums.
	WrapT     SamplerParameter `json:"wrapT"`     //T (V) wrapping mode.
}
