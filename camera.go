package main

import "math"

type Camera struct {
	origin          Point
	horizontal      Vector
	vertical        Vector
	lowerLeftCorner Point

	u Vector
	v Vector

	lensRadius float64
}

type Degrees float64

func (d Degrees) Radians() float64 {
	return float64(d) * math.Pi / 180
}

func NewCamera(lookFrom, lookAt Point, up Vector,
	vFOV Degrees, aspectRatio, aperture, focusDistance float64) Camera {
	theta := vFOV.Radians()
	h := math.Tan(theta / 2)
	viewportHeight := 2 * h
	viewportWidth := aspectRatio * viewportHeight

	w := lookFrom.SubPoint(lookAt).Unit()
	var c Camera
	c.u = up.Cross(w).Unit()
	c.v = w.Cross(c.u)
	c.origin = lookFrom
	c.horizontal = c.u.Scale(focusDistance * viewportWidth)
	c.vertical = c.v.Scale(focusDistance * viewportHeight)
	c.lowerLeftCorner = c.origin.SubVector(c.horizontal.Scale(0.5)).
		SubVector(c.vertical.Scale(0.5)).
		SubVector(w.Scale(focusDistance))
	c.lensRadius = aperture / 2
	return c
}

func (c Camera) Ray(s, t float64) Ray {
	rd := RandomUnitDiscVector().Scale(c.lensRadius)
	offset := c.u.Scale(rd.X).Add(c.v.Scale(rd.Y))
	origin := c.origin.Add(offset)
	return Ray{
		origin,
		c.lowerLeftCorner.Add(c.horizontal.Scale(s)).
			Add(c.vertical.Scale(t)).
			SubPoint(origin),
	}
}
