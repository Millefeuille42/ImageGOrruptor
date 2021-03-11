package ImageGOrruptor

import (
	"image"
)

// GOrruptor Contains the original image as a presumably unmodifiable image, and the newly corrupted image
type GOrruptor struct {
	original  image.Image
	corrupted *image.RGBA
}

// NewGOrruptor Creates a new object and copy the contents of the original image to the corrupted one
func NewGOrruptor(img image.Image) *GOrruptor {
	ret := &GOrruptor{img, image.NewRGBA(img.Bounds())}

	for x := 0; x < img.Bounds().Max.X; x++ {
		for y := 0; y < img.Bounds().Max.Y; y++ {
			ret.corrupted.Set(x, y, ret.original.At(x, y))
		}
	}

	return ret
}

// Returns the Original image
func (m *GOrruptor) Original() image.Image {
	return m.original
}

// Returns the Corrupted image
func (m *GOrruptor) Corrupted() image.Image {
	return m.corrupted
}

// Change the original image without updating the corrupted, allowing another layer of corruption in some extent
func (m *GOrruptor) Change(new image.Image) {
	m.original = new
}