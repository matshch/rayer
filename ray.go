package main

type Ray struct {
	Origin    Point
	Direction Vector
}

func (r Ray) At(t float64) Point {
	return r.Origin.Add(r.Direction.Scale(t))
}
