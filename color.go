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

	samples uint
}

func (c Color) NRGBA() color.NRGBA {
	c.blended()
	r, g, b := c.Uint8()
	return color.NRGBA{r, g, b, math.MaxUint8}
}

func (c Color) Uint8() (uint8, uint8, uint8) {
	c.blended()
	return uint8(math.Round(math.MaxUint8 * c.R)),
		uint8(math.Round(math.MaxUint8 * c.G)),
		uint8(math.Round(math.MaxUint8 * c.B))
}

func (c Color) String() string {
	c.blended()
	r, g, b := c.Uint8()
	return fmt.Sprintf("#%02x%02x%02x (%g, %g, %g)", r, g, b, c.R, c.G, c.B)
}

func (c Color) Blend(t float64, c2 Color) Color {
	c.blended()
	return Color{
		R: (1-t)*c.R + t*c2.R,
		G: (1-t)*c.G + t*c2.G,
		B: (1-t)*c.B + t*c2.B,
	}
}

func (c *Color) LazyBlend(c2 Color) {
	c.R += c2.R
	c.G += c2.G
	c.B += c2.B
	c.samples++
}

func (c *Color) blended() {
	if c.samples != 0 {
		ratio := 1 / float64(c.samples)
		c.R *= ratio
		c.G *= ratio
		c.B *= ratio
		c.samples = 0
	}
}

func (c Color) Scale(f float64) Color {
	c.blended()
	return Color{
		R: f * c.R,
		G: f * c.G,
		B: f * c.B,
	}
}
