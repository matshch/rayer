package main

type Camera struct {
	aspectRatio    float64
	viewportHeight float64
	viewportWidth  float64
	focalLength    float64

	origin          Point
	horizontal      Vector
	vertical        Vector
	lowerLeftCorner Point
}

func NewCamera(aspectRatio float64) Camera {
	c := Camera{aspectRatio: aspectRatio, viewportHeight: 2., focalLength: 1}
	c.viewportWidth = c.aspectRatio * c.viewportHeight
	c.horizontal.X = c.viewportWidth
	c.vertical.Y = c.viewportHeight
	c.lowerLeftCorner = c.origin.SubVector(c.horizontal.Scale(0.5)).
		SubVector(c.vertical.Scale(0.5)).
		SubVector(Vector{Z: c.focalLength})
	return c
}

func (c Camera) Ray(u, v float64) Ray {
	return Ray{
		c.origin,
		c.lowerLeftCorner.Add(c.horizontal.Scale(u)).Add(c.vertical.Scale(v)).SubPoint(c.origin),
	}
}
