package main

import "math"

type Dielectric struct {
	RefractionIndex float64
}

func (d Dielectric) Scatter(r Ray, hit Hit) (*Ray, *Color) {
	refractionRatio := d.RefractionIndex
	if hit.Front {
		refractionRatio = 1 / d.RefractionIndex
	}
	directionUnit := r.Direction.Unit()
	cosTheta := math.Min(directionUnit.Scale(-1).Dot(hit.Normal), 1.)
	sinTheta := math.Sqrt(1 - cosTheta*cosTheta)

	var direction Vector
	if cannotRefract := refractionRatio*sinTheta > 1; cannotRefract {
		direction = directionUnit.Reflect(hit.Normal)
	} else {
		direction = directionUnit.Refract(hit.Normal, refractionRatio)
	}
	return &Ray{hit.Point, direction}, &Color{R: 1, G: 1, B: 1}
}
