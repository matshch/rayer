package main

import "math"

type Camera struct {
	origin          Point
	horizontal      Vector
	vertical        Vector
	lowerLeftCorner Point
}

type Degrees float64

func (d Degrees) Radians() float64 {
	return float64(d) * math.Pi / 180
}

func NewCamera(vFOV Degrees, aspectRatio float64) Camera {
	theta := vFOV.Radians()
	h := math.Tan(theta / 2)
	viewportHeight := 2 * h
	viewportWidth := aspectRatio * viewportHeight
	const focalLength = 1
	var c Camera
	c.horizontal.X = viewportWidth
	c.vertical.Y = viewportHeight
	c.lowerLeftCorner = c.origin.SubVector(c.horizontal.Scale(0.5)).
		SubVector(c.vertical.Scale(0.5)).
		SubVector(Vector{Z: focalLength})
	return c
}

func (c Camera) Ray(u, v float64) Ray {
	return Ray{
		c.origin,
		c.lowerLeftCorner.Add(c.horizontal.Scale(u)).Add(c.vertical.Scale(v)).SubPoint(c.origin),
	}
}
