package gltf

//A camera's projection.  A node **MAY** reference a camera to apply a transform to place the camera in the scene.
type Camera struct {
	ChildOfRootProperty
	Orthographic *CameraOrthographic `json:"orthographic,omitempty"` //An orthographic camera containing properties to create an orthographic projection matrix. This property **MUST NOT** be defined when `perspective` is defined.
	Perspective  *CameraPerspective  `json:"perspective,omitempty"`  //Specifies if the camera uses a perspective or orthographic projection.  Based on this, either the camera's `perspective` or `orthographic` property **MUST** be defined.
	Type         CameraType          `json:"type"`                   //Specifies if the camera uses a perspective or orthographic projection.  Based on this, either the camera's `perspective` or `orthographic` property **MUST** be defined.
}

type CameraOrthographic struct {
	Property
	XMag  float64 `json:"xmag"`  //The floating-point horizontal magnification of the view. This value **MUST NOT** be equal to zero. This value **SHOULD NOT** be negative.
	YMag  float64 `json:"ymag"`  //The floating-point vertical magnification of the view. This value **MUST NOT** be equal to zero. This value **SHOULD NOT** be negative.
	ZFar  float64 `json:"zfar"`  //The floating-point distance to the far clipping plane. This value **MUST NOT** be equal to zero. `zfar` **MUST** be greater than `znear`.
	ZNear float64 `json:"znear"` //The floating-point distance to the near clipping plane.
}

type CameraPerspective struct {
	Property
	ASpectRatio float64 `json:"aspectRatio,omitempty"` //The floating-point aspect ratio of the field of view. When undefined, the aspect ratio of the rendering viewport **MUST** be used.
	YFov        float64 `json:"yfov"`                  //The floating-point vertical field of view in radians. This value **SHOULD** be less than π.
	ZFar        float64 `json:"zfar,omitempty"`        //The floating-point distance to the far clipping plane. When defined, `zfar` **MUST** be greater than `znear`. If `zfar` is undefined, client implementations **SHOULD** use infinite projection matrix.
	ZNear       float64 `json:"znear"`                 //The floating-point distance to the near clipping plane.
}
