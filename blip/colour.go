package blip

import (
	"github.com/faiface/pixel"
)

type Colour uint8

func NewColour(r uint8, g uint8, b uint8) Colour {
	return Colour(r | g | b)
}

func scaleColour(b Colour) uint8 {
	return (64 * uint8(b))
}

func scaleAlpha(a Colour) uint8 {
	if a == 1 {
		return 255
	}
	return 0
}

func (c Colour) ExtractComponents() (uint8, uint8, uint8, uint8) {
	a := (c & 0b01000000) >> 6
	r := (c & 0b00110000) >> 4
	g := (c & 0b00001100) >> 2
	b := (c & 0b00000011) >> 0
	return scaleColour(r), scaleColour(g), scaleColour(b), scaleAlpha(a)
}

func (c Colour) ToRGB() pixel.RGBA {
	r, g, b, a := c.ExtractComponents()
	return pixel.RGBA{
		R: float64(r) / 255,
		G: float64(g) / 255,
		B: float64(b) / 255,
		A: float64(a) / 255,
	}
}
