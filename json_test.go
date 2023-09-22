package gltf

import (
	"os"
	"testing"

	"github.com/o5h/testing/assert"
)

func TestXxx(t *testing.T) {
	file, err := os.Open("testdata/animation.gltf")
	assert.Nil(t, err)
	defer file.Close()
	doc, err := DecodeJson(file)
	assert.Nil(t, err)

	// TODO: more tests
	_ = doc
}
