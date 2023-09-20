package json

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"

	"github.com/o5h/gltf"
)

type decoder struct{ d *json.Decoder }

func newDecoder(r io.Reader) *decoder {
	return &decoder{d: json.NewDecoder(bufio.NewReader(r))}
}

func decode[T any](d *decoder, root *gltf.GLTF) *T {
	var t T
	err := d.d.Decode(&t)
	if err != nil {
		panic(err)
	}
	if child, ok := any(&t).(gltf.ChildOfRoot); ok {
		child.SetRoot(root)
	}
	return &t
}

func decodeArray[T any](d *decoder, root *gltf.GLTF) []*T {
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
