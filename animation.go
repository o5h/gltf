package gltf

//A keyframe animation.
type Animation struct {
	ChildOfRootProperty
	Channels []*AnimationChannel `json:"channels"` //An array of animation channels. An animation channel combines an animation sampler with a target property being animated. Different channels of the same animation **MUST NOT** have the same targets.
	Samplers []*AnimationSampler `json:"samplers"` //An array of animation samplers. An animation sampler combines timestamps with a sequence of output values and defines an interpolation algorithm.
}

//An animation channel combines an animation sampler with a target property being animated.
type AnimationChannel struct {
	Property
	Sampler uint32                 `json:"sampler"` //The index of a sampler in this animation used to compute the value for the target, e.g., a node's translation, rotation, or scale (TRS).
	Target  AnimationChannelTarget `json:"target"`  //The descriptor of the animated property.
}

//The descriptor of the animated property.
type AnimationChannelTarget struct {
	Property
	Node *uint32 `json:"node,omitempty"` //The index of the node to animate. When undefined, the animated object **MAY** be defined by an extension.
	Path Path    `json:"path"`           //The name of the node's TRS property to animate, or the `\"weights\"` of the Morph Targets it instantiates. For the `\"translation\"` property, the values that are provided by the sampler are the translation along the X, Y, and Z axes. For the `\"rotation\"` property, the values are a quaternion in the order (x, y, z, w), where w is the scalar. For the `\"scale\"` property, the values are the scaling factors along the X, Y, and Z axes.
}

type AnimationSampler struct {
	Property
	Input         uint32        `json:"input"`                   //The index of an accessor containing keyframe timestamps. The accessor **MUST** be of scalar type with floating-point components. The values represent time in seconds with `time[0] >= 0.0`, and strictly increasing values, i.e., `time[n + 1] > time[n]`.
	Interpolation Interpolation `json:"interpolation,omitempty"` //Interpolation algorithm.
	Output        uint32        `json:"output"`                  //The index of an accessor, containing keyframe output values.
}
