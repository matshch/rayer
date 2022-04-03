package main

type Lambertian struct {
	Albedo Color
}

func (l Lambertian) Scatter(r Ray, hit Hit) (*Ray, *Color) {
	direction := hit.Normal.Add(RandomUnitVector())
	if direction.NearZero() {
		direction = hit.Normal
	}
	return &Ray{hit.Point, direction}, &l.Albedo
}
