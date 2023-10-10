package gltf

import (
	"github.com/o5h/glm/vec2"
	"github.com/o5h/glm/vec3"
	"github.com/o5h/glm/vec4"
)

// The root object for a glTF asset.
type GLTF struct {
	Property
	ExtensionsUsed     []string      `json:"extensionsUsed,omitempty"`     //Names of glTF extensions used in this asset.
	ExtensionsRequired []string      `json:"extensionsRequired,omitempty"` //Names of glTF extensions required to properly load this asset.
	Accessors          []*Accessor   `json:"accessors,omitempty"`          //An array of accessors.
	Animations         []*Animation  `json:"animations,omitempty"`         //An array of keyframe animations.
	Asset              *Asset        `json:"asset"`                        //Metadata about the glTF asset.
	Buffers            []*Buffer     `json:"buffers,omitempty"`            //An array of buffers. A buffer points to binary geometry, animation, or skins.
	BufferViews        []*BufferView `json:"bufferViews,omitempty"`        //An array of bufferViews.  A bufferView is a view into a buffer generally representing a subset of the buffer.
	Cameras            []*Camera     `json:"cameras,omitempty"`            //An array of cameras.  A camera defines a projection matrix.
	Images             []*Image      `json:"images,omitempty"`             //An array of images.  An image defines data used to create a texture.
	Materials          []*Material   `json:"materials,omitempty"`          //An array of materials.  A material defines the appearance of a primitive.
	Meshes             []*Mesh       `json:"meshes,omitempty"`             //An array of meshes.  A mesh is a set of primitives to be rendered.
	Nodes              []*Node       `json:"nodes,omitempty"`              //An array of nodes.
	Samplers           []*Sampler    `json:"samplers,omitempty"`           //An array of samplers.  A sampler contains properties for texture filtering and wrapping modes.
	DefaultScene       *uint32       `json:"scene,omitempty"`              //The index of the default scene.  This property **MUST NOT** be defined, when `scenes` is undefined.
	Scenes             []*Scene      `json:"scenes,omitempty"`             //An array of scenes.
	Skins              []*Skin       `json:"skins,omitempty"`              //An array of skins.  A skin is defined by joints and matrices.
	Textures           []*Texture    `json:"textures,omitempty"`           //An array of textures.
}

// Although `extras` **MAY** have any type, it is common for applications to store and access custom data as key/value pairs. Therefore, `extras` **SHOULD** be a JSON object rather than a primitive value for best portability.
type Extras map[string]any

type Vec2 struct{ X, Y float32 }
type Vec3 struct{ X, Y, Z float32 }
type Vec4 struct{ X, Y, Z, W float32 }
type Type interface {
	Component | vec2.Vec2 | vec3.Vec3 | vec4.Vec4
}
