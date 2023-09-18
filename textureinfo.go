package gltf

//Reference to a texture.
type TextureInfo struct {
	Index    *Id    `json:"index"`    //The index of the texture.
	TexCoord uint16 `json:"texCoord"` //This integer value is used to construct a string in the format `TEXCOORD_<set index>` which is a reference to a key in `mesh.primitives.attributes` (e.g. a value of `0` corresponds to `TEXCOORD_0`). A mesh primitive **MUST** have the corresponding texture coordinate attributes for the material to be applicable to it.
}
