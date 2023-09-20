package gltf

type ComponentType uint32
type AccessorType string
type Interpolation string //Interpolation algorithm

// The name of the node's TRS property to animate, or the `\"weights\"` of the Morph Targets it instantiates. For the `\"translation\"` property, the values that are provided by the sampler are the translation along the X, Y, and Z axes. For the `\"rotation\"` property, the values are a quaternion in the order (x, y, z, w), where w is the scalar. For the `\"scale\"` property, the values are the scaling factors along the X, Y, and Z axes.
type Path string

type BufferType uint32
type CameraType string
type AlphaMode string
type DrawMode uint8

type Component interface {
	int8 | uint8 | int16 | uint16 | uint32 | float32
}

const (
	ComponentType_BYTE           = ComponentType(gl_BYTE)
	ComponentType_UNSIGNED_BYTE  = ComponentType(gl_UNSIGNED_BYTE)
	ComponentType_SHORT          = ComponentType(gl_SHORT)
	ComponentType_UNSIGNED_SHORT = ComponentType(gl_UNSIGNED_SHORT)
	ComponentType_UNSIGNED_INT   = ComponentType(gl_UNSIGNED_INT)
	ComponentType_FLOAT          = ComponentType(gl_FLOAT)

	AccessorType_SCALAR = AccessorType("SCALAR")
	AccessorType_VEC2   = AccessorType("VEC2")
	AccessorType_VEC3   = AccessorType("VEC3")
	AccessorType_VEC4   = AccessorType("VEC4")
	AccessorType_MAT2   = AccessorType("MAT2")
	AccessorType_MAT3   = AccessorType("MAT3")
	AccessorType_MAT4   = AccessorType("MAT4")

	Interpolation_LINEAR      = Interpolation("LINEAR")      //The animated values are linearly interpolated between keyframes. When targeting a rotation, spherical linear interpolation (slerp) **SHOULD** be used to interpolate quaternions. The number of output elements **MUST** equal the number of input elements.
	Interpolation_STEP        = Interpolation("STEP")        //The animated values remain constant to the output of the first keyframe, until the next keyframe. The number of output elements **MUST** equal the number of input elements.
	Interpolation_CUBICSPLINE = Interpolation("CUBICSPLINE") //The animation's interpolation is computed using a cubic spline with specified tangents. The number of output elements **MUST** equal three times the number of input elements. For each input element, the output stores three elements, an in-tangent, a spline vertex, and an out-tangent. There **MUST** be at least two keyframes when using this interpolation.

	Path_TRANSLATION = Path("translation")
	Path_ROTATION    = Path("rotation")
	Path_SCALE       = Path("scale")
	Path_WEIGHTS     = Path("weights")

	BufferType_ARRAY_BUFFER         = BufferType(gl_ARRAY_BUFFER)
	BufferType_ELEMENT_ARRAY_BUFFER = BufferType(gl_ELEMENT_ARRAY_BUFFER)

	PERSPECTIVE  = CameraType("perspective")
	ORTHOGRAPHIC = CameraType("orthographic")

	MimeType_Image_JPEG = "image/jpeg"
	MimeType_Image_PNG  = "image/png"

	AlphaMode_OPAQUE = AlphaMode("OPAQUE")
	AlphaMode_MASK   = AlphaMode("MASK")
	AlphaMode_BLEND  = AlphaMode("BLEND")

	DrawMode_POINTS         = DrawMode(0)
	DrawMode_LINES          = DrawMode(1)
	DrawMode_LINE_LOOP      = DrawMode(2)
	DrawMode_LINE_STRIP     = DrawMode(3)
	DrawMode_TRIANGLES      = DrawMode(4)
	DrawMode_TRIANGLE_STRIP = DrawMode(5)
	DrawMode_TRIANGLE_FAN   = DrawMode(6)
)

var (
	octet_stream_base64_prefix = "data:application/octet-stream;base64,"
)
