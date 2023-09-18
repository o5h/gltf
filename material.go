package gltf

type Material struct {
	ChildOfRootProperty
	PBRMetallicRoughness *MaterialPBRMetallicRoughness `json:"pbrMetallicRoughness,omitempty"` //A set of parameter values that are used to define the metallic-roughness material model from Physically Based Rendering (PBR) methodology. When undefined, all the default values of `pbrMetallicRoughness` **MUST** apply.
	NormalTexture        *MaterialNormalTextureInfo    `json:"normalTexture,omitempty"`        //The tangent space normal texture. The texture encodes RGB components with linear transfer function. Each texel represents the XYZ components of a normal vector in tangent space. The normal vectors use the convention +X is right and +Y is up. +Z points toward the viewer. If a fourth component (A) is present, it **MUST** be ignored. When undefined, the material does not have a tangent space normal texture.
	OcclusionTexture     *OcclusionTexture             `json:"occlusionTexture,omitempty"`     //The occlusion texture. The occlusion values are linearly sampled from the R channel. Higher values indicate areas that receive full indirect lighting and lower values indicate no indirect lighting. If other channels are present (GBA), they **MUST** be ignored for occlusion calculations. When undefined, the material does not have an occlusion texture.
	EmissiveTexture      *TextureInfo                  `json:"emissiveTexture,omitempty"`      //The emissive texture. It controls the color and intensity of the light being emitted by the material. This texture contains RGB components encoded with the sRGB transfer function. If a fourth component (A) is present, it **MUST** be ignored. When undefined, the texture **MUST** be sampled as having `1.0` in RGB components.
	EmissiveFactor       [3]float64                    `json:"emissiveFactor,omitempty"`       //The factors for the emissive color of the material. This value defines linear multipliers for the sampled texels of the emissive texture.
	AlphaMode            AlphaMode                     `json:"alphaMode,omitempty"`            //The material's alpha rendering mode enumeration specifying the interpretation of the alpha value of the base color.
	AlphaCutoff          float64                       `json:"alphaCutoff,omitempty"`          //Specifies the cutoff threshold when in `MASK` alpha mode. If the alpha value is greater than or equal to this value then it is rendered as fully opaque, otherwise, it is rendered as fully transparent. A value greater than `1.0` will render the entire material as fully transparent. This value **MUST** be ignored for other alpha modes. When `alphaMode` is not defined, this value **MUST NOT** be defined.
	DoubleSided          bool                          `json:"doubleSided,omitempty"`          //Specifies whether the material is double sided. When this value is false, back-face culling is enabled. When this value is true, back-face culling is disabled and double-sided lighting is enabled. The back-face **MUST** have its normals reversed before the lighting equation is evaluated.
}

// Material PBR Metallic Roughness
type MaterialPBRMetallicRoughness struct {
	Property
	BaseColorFactor          [4]float64   `json:"baseColorFactor"`          //The factors for the base color of the material. This value defines linear multipliers for the sampled texels of the base color texture.
	BaseColorTexture         *TextureInfo `json:"baseColorTexture"`         //The base color texture. The first three components (RGB) **MUST** be encoded with the sRGB transfer function. They specify the base color of the material. If the fourth component (A) is present, it represents the linear alpha coverage of the material. Otherwise, the alpha coverage is equal to `1.0`. The `material.alphaMode` property specifies how alpha is interpreted. The stored texels **MUST NOT** be premultiplied. When undefined, the texture **MUST** be sampled as having `1.0` in all components.
	MetallicFactor           *float64     `json:"metallicFactor"`           //The factor for the metalness of the material. This value defines a linear multiplier for the sampled metalness values of the metallic-roughness texture.
	RoughnessFactor          *float64     `json:"roughnessFactor"`          //The factor for the roughness of the material. This value defines a linear multiplier for the sampled roughness values of the metallic-roughness texture.
	MetallicRoughnessTexture *TextureInfo `json:"metallicRoughnessTexture"` //The metallic-roughness texture. The metalness values are sampled from the B channel. The roughness values are sampled from the G channel. These values **MUST** be encoded with a linear transfer function. If other channels are present (R or A), they **MUST** be ignored for metallic-roughness calculations. When undefined, the texture **MUST** be sampled as having `1.0` in G and B components.
}

// Material Normal Texture Info
type MaterialNormalTextureInfo struct {
	TextureInfo
	Scale float64 //The scalar parameter applied to each normal vector of the texture. This value scales the normal vector in X and Y directions using the formula: `scaledNormal =  normalize((<sampled normal texture value> * 2.0 - 1.0) * vec3(<normal scale>, <normal scale>, 1.0))`.
}

// Material Occlusion Texture Info
type OcclusionTexture struct {
	TextureInfo
	Strength float64 `json:"strength"` //A scalar parameter controlling the amount of occlusion applied. A value of `0.0` means no occlusion. A value of `1.0` means full occlusion. This value affects the final occlusion value as: `1.0 + strength * (<sampled occlusion texture value> - 1.0)`.
}
