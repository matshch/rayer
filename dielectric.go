package main

import (
	"math"
	"math/rand"
)

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
	cannotRefract := refractionRatio*sinTheta > 1
	if cannotRefract || reflectance(cosTheta, refractionRatio) > rand.Float64() {
		direction = directionUnit.Reflect(hit.Normal)
	} else {
		direction = directionUnit.Refract(hit.Normal, refractionRatio)
	}
	return &Ray{hit.Point, direction}, &Color{R: 1, G: 1, B: 1}
}

func reflectance(cos, refractionRatio float64) float64 {
	r0 := (1 - refractionRatio) / (1 + refractionRatio)
	r0 = r0 * r0
	return r0 + (1-r0)*math.Pow((1-cos), 5)
}
