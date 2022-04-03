package main

import "math"

type Sphere struct {
	Center   Point
	Radius   float64
	Material Material
}

func (s Sphere) Hit(r Ray, tMin, tMax float64) *Hit {
	oc := r.Origin.SubPoint(s.Center)
	a := r.Direction.LenSq()
	h := r.Direction.Dot(oc)
	c := oc.LenSq() - s.Radius*s.Radius
	d := h*h - a*c
	if d < 0 {
		return nil
	}
	sqrtD := math.Sqrt(d)
	t := (-h - sqrtD) / a
	if t < tMin || tMax < t {
		t = (-h + sqrtD) / a
		if t < tMin || tMax < t {
			return nil
		}
	}
	p := r.At(t)
	hit := Hit{Point: p, Material: s.Material, T: t}
	hit.SetFrontNormal(r, p.SubPoint(s.Center).Scale(1/s.Radius))
	return &hit
}
