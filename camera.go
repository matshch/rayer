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

func NewCamera(lookFrom, lookAt Point, up Vector, vFOV Degrees, aspectRatio float64) Camera {
	theta := vFOV.Radians()
	h := math.Tan(theta / 2)
	viewportHeight := 2 * h
	viewportWidth := aspectRatio * viewportHeight

	w := lookFrom.SubPoint(lookAt).Unit()
	u := up.Cross(w).Unit()
	v := w.Cross(u)

	var c Camera
	c.origin = lookFrom
	c.horizontal = u.Scale(viewportWidth)
	c.vertical = v.Scale(viewportHeight)
	c.lowerLeftCorner = c.origin.SubVector(c.horizontal.Scale(0.5)).
		SubVector(c.vertical.Scale(0.5)).
		SubVector(w)
	return c
}

func (c Camera) Ray(s, t float64) Ray {
	return Ray{
		c.origin,
		c.lowerLeftCorner.Add(c.horizontal.Scale(s)).Add(c.vertical.Scale(t)).SubPoint(c.origin),
	}
}
