package gltf

type ComponentType uint32
type Component interface {
	int8 | uint8 | int16 | uint16 | uint32 | float32
}
type IndexComponent interface{ uint8 | uint16 | uint32 }

type AccessorType string
type Interpolation string //Interpolation algorithm

// The name of the node's TRS property to animate, or the `\"weights\"` of the Morph Targets it instantiates. For the `\"translation\"` property, the values that are provided by the sampler are the translation along the X, Y, and Z axes. For the `\"rotation\"` property, the values are a quaternion in the order (x, y, z, w), where w is the scalar. For the `\"scale\"` property, the values are the scaling factors along the X, Y, and Z axes.
type Path string

type BufferType uint32
type CameraType string
type AlphaMode string
type MeshPrimitiveMode uint8

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

	MeshPrimitiveMode_POINTS         = MeshPrimitiveMode(0)
	MeshPrimitiveMode_LINES          = MeshPrimitiveMode(1)
	MeshPrimitiveMode_LINE_LOOP      = MeshPrimitiveMode(2)
	MeshPrimitiveMode_LINE_STRIP     = MeshPrimitiveMode(3)
	MeshPrimitiveMode_TRIANGLES      = MeshPrimitiveMode(4)
	MeshPrimitiveMode_TRIANGLE_STRIP = MeshPrimitiveMode(5)
	MeshPrimitiveMode_TRIANGLE_FAN   = MeshPrimitiveMode(6)
)

var (
	octet_stream_base64_prefix = "data:application/octet-stream;base64,"
)

func (t AccessorType) Size() uint32 {
	switch t {
	case AccessorType_SCALAR:
		return 1
	case AccessorType_VEC2:
		return 2
	case AccessorType_VEC3:
		return 3
	case AccessorType_VEC4:
		return 4
	case AccessorType_MAT2:
		return 2 * 2
	case AccessorType_MAT3:
		return 3 * 3
	case AccessorType_MAT4:
		return 4 * 4
	}
	return 1
}

func ComponentTypeOf[T Component]() ComponentType {
	var t T
	switch any(t).(type) {
	case int8:
		return ComponentType_BYTE
	case uint8:
		return ComponentType_UNSIGNED_BYTE
	case int16:
		return ComponentType_SHORT
	case uint16:
		return ComponentType_UNSIGNED_SHORT
	case uint32:
		return ComponentType_UNSIGNED_INT
	case float32:
		return ComponentType_FLOAT
	}
	return ComponentType_BYTE
}

const (
	gl_BYTE           = 5120 //0x1400
	gl_UNSIGNED_BYTE  = 5121 //0x1401
	gl_SHORT          = 5122 //0x1402
	gl_UNSIGNED_SHORT = 5123 //0x1403
	gl_UNSIGNED_INT   = 5125 //0x1405
	gl_FLOAT          = 5126 //0x1406

	gl_NEAREST                = 0x2600
	gl_LINEAR                 = 0x2601
	gl_NEAREST_MIPMAP_NEAREST = 0x2700
	gl_LINEAR_MIPMAP_NEAREST  = 0x2701
	gl_NEAREST_MIPMAP_LINEAR  = 0x2702
	gl_LINEAR_MIPMAP_LINEAR   = 0x2703
	gl_REPEAT                 = 0x2901
	gl_CLAMP_TO_EDGE          = 0x812F
	gl_MIRRORED_REPEAT        = 0x8370

	gl_ARRAY_BUFFER         = 34962 //0x8892
	gl_ELEMENT_ARRAY_BUFFER = 34963 //0x8893
)
