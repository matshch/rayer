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
	return Point{p.X + v.X, p.Y + v.Y, p.Z + v.Z}
}

func (p Point) SubVector(v Vector) Point {
	return Point{p.X - v.X, p.Y - v.Y, p.Z - v.Z}
}

func (p Point) SubPoint(p2 Point) Vector {
	return Vector{p.X - p2.X, p.Y - p2.Y, p.Z - p2.Z}
}
