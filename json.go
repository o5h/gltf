package gltf

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"io"
)

func DecodeJson(r io.Reader) (doc *GLTF, err error) {
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

func decodeFrom(r io.Reader) *GLTF {
	d := newDecoder(r)
	root := &GLTF{}
	d.expectDelim('{')
	for d.more() {
		t := d.token()
		switch t {
		case "asset":
			root.Asset = decode[Asset](d, root)
		case "scene":
			root.DefaultScene = decode[uint32](d, root)
		case "scenes":
			root.Scenes = decodeArray[Scene](d, root)
		case "nodes":
			root.Nodes = decodeArray[Node](d, root)
		case "materials":
			root.Materials = decodeArray[Material](d, root)
		case "meshes":
			root.Meshes = decodeArray[Mesh](d, root)
		case "accessors":
			root.Accessors = decodeArray[Accessor](d, root)
		case "bufferViews":
			root.BufferViews = decodeArray[BufferView](d, root)
		case "buffers":
			root.Buffers = decodeArray[Buffer](d, root)
		case "animations":
			root.Animations = decodeArray[Animation](d, root)
		case "skins":
			root.Skins = decodeArray[Skin](d, root)
		default:
			panic(fmt.Sprintf("unsupported token `%v`", t))
		}
	}
	return root
}

type decoder struct{ d *json.Decoder }

func newDecoder(r io.Reader) *decoder {
	return &decoder{d: json.NewDecoder(bufio.NewReader(r))}
}

func decode[T any](d *decoder, root *GLTF) *T {
	var t T
	err := d.d.Decode(&t)
	if err != nil {
		panic(err)
	}
	if child, ok := any(&t).(ChildOfRoot); ok {
		child.SetRoot(root)
	}
	return &t
}

func decodeArray[T any](d *decoder, root *GLTF) []*T {
	d.expectDelim('[')
	var result []*T
	for d.d.More() {
		result = append(result, decode[T](d, root))
	}
	d.expectDelim(']')
	return result
}

func (d *decoder) more() bool { return d.d.More() }

func (d *decoder) token() json.Token {
	t, err := d.d.Token()
	if err != nil {
		panic(err)
	}
	return t
}

func (d *decoder) expectDelim(r rune) { d.expect(json.Delim(r)) }

func (d *decoder) expect(expected any) {
	t, err := d.d.Token()
	if err != nil {
		panic(err)
	}
	if t != expected {
		panic(fmt.Errorf("got token %v, want token %v", t, expected))
	}
}
