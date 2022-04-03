package main

import (
	"fmt"
	"image/color"
	"math"
)

type Color struct {
	R float64
	G float64
	B float64
}

func (c Color) NRGBA() color.NRGBA {
	r, g, b := c.Uint8()
	return color.NRGBA{r, g, b, math.MaxUint8}
}

func (c Color) Uint8() (uint8, uint8, uint8) {
	return uint8(math.Round(math.MaxUint8 * c.R)),
		uint8(math.Round(math.MaxUint8 * c.G)),
		uint8(math.Round(math.MaxUint8 * c.B))
}

func (c Color) String() string {
	r, g, b := c.Uint8()
	return fmt.Sprintf("#%02x%02x%02x (%g, %g, %g)", r, g, b, c.R, c.G, c.B)
}

func (c Color) Blend(t float64, c2 Color) Color {
	return Color{(1-t)*c.R + t*c2.R, (1-t)*c.G + t*c2.G, (1-t)*c.B + t*c2.B}
}
