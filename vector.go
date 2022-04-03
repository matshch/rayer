package main

import (
	"fmt"
	"math"
)

type Vector struct {
	X float64
	Y float64
	Z float64
}

func (v Vector) String() string {
	return fmt.Sprintf("<%g, %g, %g>", v.X, v.Y, v.Z)
}

func (v Vector) Len() float64 {
	return math.Sqrt(v.LenSq())
}

func (v Vector) LenSq() float64 {
	return v.Dot(v)
}

func (v Vector) Scale(t float64) Vector {
	return Vector{t * v.X, t * v.Y, t * v.Z}
}

func (v Vector) Unit() Vector {
	return v.Scale(1 / v.Len())
}

func (a Vector) Add(b Vector) Vector {
	return Vector{a.X + b.X, a.Y + b.Y, a.Z + b.Z}
}

func (a Vector) Sub(b Vector) Vector {
	return Vector{a.X - b.X, a.Y - b.Y, a.Z - b.Z}
}

func (v Vector) Dot(v2 Vector) float64 {
	return v.X*v2.X + v.Y*v2.Y + v.Z*v2.Z
}

func (v Vector) NearZero() bool {
	const s = 1e-8
	return (math.Abs(v.X) < s) && (math.Abs(v.Y) < s) && (math.Abs(v.Z) < s)
}

func (v Vector) Reflect(n Vector) Vector {
	return v.Sub(n.Scale(2 * v.Dot(n)))
}

func (v Vector) Refract(n Vector, refractionRatio float64) Vector {
	cosTheta := math.Min(v.Scale(-1).Dot(n), 1.)
	rPerp := v.Add(n.Scale(cosTheta)).Scale(refractionRatio)
	rPar := n.Scale(-math.Sqrt(math.Abs(1. - rPerp.LenSq())))
	return rPerp.Add(rPar)
}

func RandomVector() Vector {
	return Vector{Random(-1, 1), Random(-1, 1), Random(-1, 1)}
}

func RandomUnitSphereVector() Vector {
	for {
		v := RandomVector()
		if v.LenSq() < 1 {
			return v
		}
	}
}

func RandomUnitVector() Vector {
	return RandomUnitSphereVector().Unit()
}
