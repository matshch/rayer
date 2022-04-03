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
	return math.Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z)
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
