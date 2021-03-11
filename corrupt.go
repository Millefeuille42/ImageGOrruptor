package ImageGOrruptor

import (
	"image/color"
	"math/rand"
)

func (m *GOrruptor) dust(strength int) {
	start := rand.Intn(m.corrupted.Bounds().Max.Y)
	stop :=	start + rand.Intn(20)
	back := m.corrupted

	for x := 0; x < m.corrupted.Bounds().Max.X; x++ {
		for y := start; y < stop; y++ {
			m.corrupted.Set(x + rand.Intn(strength) * posOrNeg(), y, back.At(x, y))
		}
	}
}

func (m *GOrruptor) verticalTranspose(strength int) {
	start := rand.Intn(m.corrupted.Bounds().Max.Y)
	stop :=	start + rand.Intn(20)
	randomStrength := rand.Intn(strength) * -1
	back := m.corrupted

	for y := 0; y < m.corrupted.Bounds().Max.Y; y++ {
		for x := start; x < stop; x++ {
			m.corrupted.Set(x, y + randomStrength, back.At(x, y))
		}
	}
}

func (m *GOrruptor) horizontalTranspose(strength int) {
	start := rand.Intn(m.corrupted.Bounds().Max.Y)
	stop :=	start + rand.Intn(20)
	back := m.corrupted
	randomStrength := rand.Intn(strength) * posOrNeg()

	for x := 0; x < m.corrupted.Bounds().Max.X; x++ {
		for y := start; y < stop; y++ {
			m.corrupted.Set(x + randomStrength, y, back.At(x, y))
		}
	}
}

func (m *GOrruptor) colorShift(start, stop [2]int, colorInc color.RGBA) {
	back := m.corrupted

	if start[0] >= stop[0] || start[1] >= stop[1] {
		return
	}

	for x := start[0]; x < stop[0]; x++ {
		for y := start[1]; y < stop[1]; y++ {
			or, og, ob, _ := back.At(x, y).RGBA()

			newColor := color.RGBA {
				R: uint8(or / 255) + colorInc.R,
				G: uint8(og / 255) + colorInc.G,
				B: uint8(ob / 255) + colorInc.B,
				A: 255,
			}
			m.corrupted.Set(x, y, newColor)
		}
	}
}