package main

type Hit struct {
	Point    Point
	Normal   Vector
	Material Material
	T        float64
	Front    bool
}

func (h *Hit) SetFrontNormal(r Ray, outwardNormal Vector) {
	h.Front = r.Direction.Dot(outwardNormal) < 0
	if h.Front {
		h.Normal = outwardNormal
	} else {
		h.Normal = outwardNormal.Scale(-1)
	}
}

type Hitter interface {
	Hit(r Ray, tMin, tMax float64) *Hit
}

type HitterSlice []Hitter

func (h HitterSlice) Hit(r Ray, tMin, tMax float64) *Hit {
	var hit *Hit
	minT := tMax
	for _, hitter := range h {
		tmpHit := hitter.Hit(r, tMin, minT)
		if tmpHit != nil {
			hit = tmpHit
			minT = hit.T
		}
	}
	return hit
}
