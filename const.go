package gltf

type ComponentType uint32
type AccessorType string
type Interpolation string //Interpolation algorithm

// The name of the node's TRS property to animate, or the `\"weights\"` of the Morph Targets it instantiates. For the `\"translation\"` property, the values that are provided by the sampler are the translation along the X, Y, and Z axes. For the `\"rotation\"` property, the values are a quaternion in the order (x, y, z, w), where w is the scalar. For the `\"scale\"` property, the values are the scaling factors along the X, Y, and Z axes.
type Path string

const (
	GL_BYTE           = 0x1400
	GL_UNSIGNED_BYTE  = 0x1401
	GL_SHORT          = 0x1402
	GL_UNSIGNED_SHORT = 0x1403
	GL_UNSIGNED_INT   = 0x1405
	GL_FLOAT          = 0x1406

	GL_ARRAY_BUFFER         = 0x8892
	GL_ELEMENT_ARRAY_BUFFER = 0x8893

	BYTE           = ComponentType(GL_BYTE)
	UNSIGNED_BYTE  = ComponentType(GL_UNSIGNED_BYTE)
	SHORT          = ComponentType(GL_SHORT)
	UNSIGNED_SHORT = ComponentType(GL_UNSIGNED_SHORT)
	UNSIGNED_INT   = ComponentType(GL_UNSIGNED_INT)
	FLOAT          = ComponentType(GL_FLOAT)

	SCALAR = AccessorType("SCALAR")
	VEC2   = AccessorType("VEC2")
	VEC3   = AccessorType("VEC3")
	VEC4   = AccessorType("VEC4")
	MAT2   = AccessorType("MAT2")
	MAT3   = AccessorType("MAT3")
	MAT4   = AccessorType("MAT4")

	LINEAR      = Interpolation("LINEAR")      //The animated values are linearly interpolated between keyframes. When targeting a rotation, spherical linear interpolation (slerp) **SHOULD** be used to interpolate quaternions. The number of output elements **MUST** equal the number of input elements.
	STEP        = Interpolation("STEP")        //The animated values remain constant to the output of the first keyframe, until the next keyframe. The number of output elements **MUST** equal the number of input elements.
	CUBICSPLINE = Interpolation("CUBICSPLINE") //The animation's interpolation is computed using a cubic spline with specified tangents. The number of output elements **MUST** equal three times the number of input elements. For each input element, the output stores three elements, an in-tangent, a spline vertex, and an out-tangent. There **MUST** be at least two keyframes when using this interpolation.

	Translation = Path("translation")
	Rotation    = Path("rotation")
	Scale       = Path("scale")
	Weights     = Path("weights")
)
