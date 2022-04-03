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
