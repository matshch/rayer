package main

import (
	"fmt"
	"image/color"
	"math"
	"math/rand"
)

const gamma = 2.2

type Color struct {
	R float64
	G float64
	B float64

	samples uint
}

func NewColor(r, g, b float64) Color {
	return Color{R: r, G: g, B: b}
}

func RandomColor() Color {
	return Color{R: rand.Float64(), G: rand.Float64(), B: rand.Float64()}
}

func RandomRangeColor(min, max float64) Color {
	return Color{R: Random(min, max), G: Random(min, max), B: Random(min, max)}
}

func (c Color) NRGBA() color.NRGBA {
	c.blended()
	r, g, b := c.Uint8()
	return color.NRGBA{r, g, b, math.MaxUint8}
}

func (c Color) Uint8() (uint8, uint8, uint8) {
	c.blended()
	return uint8(math.Round(math.MaxUint8 * math.Pow(c.R, 1/gamma))),
		uint8(math.Round(math.MaxUint8 * math.Pow(c.G, 1/gamma))),
		uint8(math.Round(math.MaxUint8 * math.Pow(c.B, 1/gamma)))
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

func (c Color) Permute(c2 Color) Color {
	c.blended()
	return Color{
		R: c.R * c2.R,
		G: c.G * c2.G,
		B: c.B * c2.B,
	}
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
