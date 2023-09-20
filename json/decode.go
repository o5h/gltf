package json

import (
	"errors"
	"fmt"
	"io"

	"github.com/o5h/gltf"
)

func Decode(r io.Reader) (doc *gltf.GLTF, err error) {
	defer func() {
		if r := recover(); r != nil {
			doc, err = nil, convertRecoverToError(r)
		}
	}()
	doc = decodeFrom(r)
	return
}

func convertRecoverToError(r interface{}) error {
	switch x := r.(type) {
	case string:
		return errors.New(x)
	case error:
		return x
	default:
		return errors.New(fmt.Sprint(x))
	}
}

func decodeFrom(r io.Reader) *gltf.GLTF {
	d := newDecoder(r)
	root := &gltf.GLTF{}
	d.expectDelim('{')
	for d.more() {
		t := d.token()
		switch t {
		case "asset":
			root.Asset = decode[gltf.Asset](d, root)
		case "scene":
			root.DefaultScene = decode[uint32](d, root)
		case "scenes":
			root.Scenes = decodeArray[gltf.Scene](d, root)
		case "nodes":
			root.Nodes = decodeArray[gltf.Node](d, root)
		case "materials":
			root.Materials = decodeArray[gltf.Material](d, root)
		case "meshes":
			root.Meshes = decodeArray[gltf.Mesh](d, root)
		case "accessors":
			root.Accessors = decodeArray[gltf.Accessor](d, root)
		case "bufferViews":
			root.BufferViews = decodeArray[gltf.BufferView](d, root)
		case "buffers":
			root.Buffers = decodeArray[gltf.Buffer](d, root)
		case "animations":
			root.Animations = decodeArray[gltf.Animation](d, root)
		case "skins":
			root.Skins = decodeArray[gltf.Skin](d, root)
		default:
			panic(fmt.Sprintf("unsupported token `%v`", t))
		}
	}
	return root
}
