package main

type Metal struct {
	Albedo Color
}

func (m Metal) Scatter(r Ray, hit Hit) (*Ray, *Color) {
	reflected := r.Direction.Unit().Reflect(hit.Normal)
	if reflected.Dot(hit.Normal) <= 0 {
		return nil, nil
	}
	return &Ray{hit.Point, reflected}, &m.Albedo
}
