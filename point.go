package main

import "fmt"

type Point struct {
	X float64
	Y float64
	Z float64
}

func (p Point) String() string {
	return fmt.Sprintf("(%g, %g, %g)", p.X, p.Y, p.Z)
}
