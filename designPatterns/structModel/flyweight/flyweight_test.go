package flyweight

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func ExampleFlyweight() {
	viewer := NewImageViewer("image1.png")
	viewer.Display()
	// Output:
	// Display: image data image1.png
}

func TestFlyweight(t *testing.T) {
	viewer1 := NewImageViewer("image1.png")
	viewer2 := NewImageViewer("image1.png")
	assert.Equal(t, viewer1.ImageFlyweight, viewer2.ImageFlyweight)
}