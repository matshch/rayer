package main

type Dielectric struct {
	RefractionIndex float64
}

func (d Dielectric) Scatter(r Ray, hit Hit) (*Ray, *Color) {
	refractionRatio := d.RefractionIndex
	if hit.Front {
		refractionRatio = 1 / d.RefractionIndex
	}
	return &Ray{
		hit.Point,
		r.Direction.Unit().Refract(hit.Normal, refractionRatio),
	}, &Color{R: 1, G: 1, B: 1}
}
