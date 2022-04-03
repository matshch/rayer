package main

import "fmt"

type Vector struct {
	X float64
	Y float64
	Z float64
}

func (v Vector) String() string {
	return fmt.Sprintf("<%g, %g, %g>", v.X, v.Y, v.Z)
}

func (v Vector) Scale(t float64) Vector {
	return Vector{X: t * v.X, Y: t * v.Y, Z: t * v.Z}
}

func (a Vector) Add(b Vector) Vector {
	return Vector{X: a.X + b.X, Y: a.Y + b.Y, Z: a.Z + b.Z}
}

func (a Vector) Sub(b Vector) Vector {
	return Vector{X: a.X - b.X, Y: a.Y - b.Y, Z: a.Z - b.Z}
}
