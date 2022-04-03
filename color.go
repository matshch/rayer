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

func (c Color) NRGBA() (color.NRGBA) {
	r, g, b := c.Uint8()
	return color.NRGBA{R: r, G: g, B: b, A: math.MaxUint8}
}

func (c Color) Uint8() (uint8, uint8, uint8) {
	return uint8(math.Round(math.MaxUint8 * c.R)), uint8(math.Round(math.MaxUint8 * c.G)), uint8(math.Round(math.MaxUint8 * c.B))
}

func (c Color) String() string {
	r, g, b := c.Uint8()
	return fmt.Sprintf("#%02x%02x%02x (%g, %g, %g)", r, g, b, c.R, c.G, c.B)
}
