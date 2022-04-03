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

func (p Point) Add(v Vector) Point {
	return Point{X: p.X + v.X, Y: p.Y + v.Y, Z: p.Z + v.Z}
}
