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
	doc := &gltf.GLTF{}
	d.expectDelim('{')
	for d.more() {
		t := d.token()
		switch t {
		case "asset":
			doc.Asset = decode[gltf.Asset](d)
		case "scene":
			doc.Scene = decode[uint32](d)
		case "scenes":
			doc.Scenes = decodeArray[gltf.Scene](d)
		case "nodes":
			doc.Nodes = decodeArray[gltf.Node](d)
		case "materials":
			doc.Materials = decodeArray[gltf.Material](d)
		case "meshes":
			doc.Meshes = decodeArray[gltf.Mesh](d)
		case "accessors":
			doc.Accessors = decodeArray[gltf.Accessor](d)
		case "bufferViews":
			doc.BufferViews = decodeArray[gltf.BufferView](d)
		case "buffers":
			doc.Buffers = decodeArray[gltf.Buffer](d)
		case "animations":
			doc.Animations = decodeArray[gltf.Animation](d)
		case "skins":
			doc.Skins = decodeArray[gltf.Skin](d)
		default:
			panic(fmt.Sprintf("unsupported token `%v`", t))
		}
	}
	return doc
}
