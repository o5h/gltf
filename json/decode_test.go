package json

import (
	"os"
	"testing"

	"github.com/o5h/gltf"
	"github.com/o5h/testing/assert"
)

func TestXxx(t *testing.T) {
	file, err := os.Open("testdata/minimal.gltf")
	assert.Nil(t, err)
	defer file.Close()
	doc, err := Decode(file)
	assert.Nil(t, err)
	assert.Eq(t, doc.Asset.Version, "2.0")
	assert.Eq(t, *doc.DefaultScene, 0)
	assert.Eq(t, doc.Scenes[0].Name, "")

	accessor := doc.GetAccessor(0).V
	buf := gltf.GetAccessorBuffer[uint16](accessor)
	_ = buf

	view := doc.GetBufferView(0).V
	viewBytes := view.Bytes()
	assert.Nil(t, err)
	_ = viewBytes
	buf0 := doc.GetBuffer(0)
	bytes := buf0.V.Bytes()
	assert.Nil(t, err)
	assert.Eq(t, len(bytes), int(buf0.V.ByteLength))
}
