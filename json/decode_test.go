package json

import (
	"os"
	"testing"

	"github.com/o5h/testing/assert"
)

func TestXxx(t *testing.T) {
	file, err := os.Open("testdata/animation.gltf")
	assert.Nil(t, err)
	defer file.Close()
	doc, err := Decode(file)
	assert.Nil(t, err)
	assert.Eq(t, doc.Asset.Version, "2.0")
	assert.Eq(t, *doc.Scene, 0)
	assert.Eq(t, doc.Scenes[0].Name, "Scene")
}
