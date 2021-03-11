package ImageGOrruptor

import (
	"image/color"
	"math"
	"math/rand"
)

// Horizontally Dust the image
func (m *GOrruptor) Dust(strength int) {
	max := rand.Intn(int(math.Abs(float64(strength / 10))) + 1)

	for i := 0; i <= max; i++ {
		m.dust(strength)
	}
}

// VerticalTranspose Perform a vertical transpose of certain columns of the image
func (m *GOrruptor) VerticalTranspose(strength int) {
	max := rand.Intn(int(math.Abs(float64(strength / 10))) + 1)

	for i := 0; i <= max; i++ {
		m.verticalTranspose(strength)
	}
}

// HorizontalTranspose Perform a horizontal transpose of certain rows of the image
func (m *GOrruptor) HorizontalTranspose(strength int) {
	max := rand.Intn(int(math.Abs(float64(strength / 10))) + 1)

	for i := 0; i <= max; i++ {
		m.horizontalTranspose(strength)
	}
}

// ColorShift Shift every pixels rgb
func (m *GOrruptor) ColorShift(strength int) {
	max := rand.Intn(int(math.Abs(float64(strength / 10))) + 1)

	for i := 0; i <= max; i++ {
		m.colorShift([2]int{0, 0}, [2]int{m.corrupted.Bounds().Max.X, m.corrupted.Bounds().Max.Y}, color.RGBA {
			R: uint8(rand.Intn(strength * 10)),
			G: uint8(rand.Intn(strength * 10)),
			B: uint8(rand.Intn(strength * 10)),
			A: 255,
		})
	}
}

