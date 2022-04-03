package main

type Camera struct {
	origin          Point
	horizontal      Vector
	vertical        Vector
	lowerLeftCorner Point
}

func NewCamera(aspectRatio float64) Camera {
	const viewportHeight = 2.
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
