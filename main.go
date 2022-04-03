package main

import (
	"image"
	"image/png"
	"log"
	"math"
	"os"

	"github.com/schollz/progressbar/v3"
)

func hitSphere(center Point, radius float64, r Ray) float64 {
	oc := r.Origin.SubPoint(center)
	a := r.Direction.LenSq()
	h := r.Direction.Dot(oc)
	c := oc.LenSq() - radius*radius
	d := h*h - a*c
	if d < 0 {
		return -1
	}
	return (-h - math.Sqrt(d)) / a
}

func rayColor(r Ray) Color {
	sphereCenter := Point{0, 0, -1}
	t := hitSphere(sphereCenter, 0.5, r)
	if t > 0 {
		n := r.At(t).SubPoint(sphereCenter).Unit()
		return Color{.5*n.X + .5, .5*n.Y + .5, .5*n.Z + .5}
	}
	unit := r.Direction.Unit()
	t = 0.5 * (unit.Y + 1.0)
	return Color{1, 1, 1}.Blend(t, Color{0.5, 0.7, 1})
}

func main() {
	log.Print("Starting rendering...")

	const aspectRatio = 16. / 9.
	const imageWidth = 400
	imageHeight := int(math.Round(imageWidth / aspectRatio))

	const viewportHeight = 2
	const viewportWidth = 2 * aspectRatio
	const focalLength = 1.

	origin := Point{}
	horizontal := Vector{viewportWidth, 0, 0}
	vertical := Vector{0, viewportHeight, 0}
	lowerLeftCorner := origin.SubVector(horizontal.Scale(0.5)).
		SubVector(vertical.Scale(0.5)).
		SubVector(Vector{0, 0, focalLength})

	img := image.NewNRGBA(image.Rect(0, 0, imageWidth, imageHeight))
	bar := progressbar.Default(int64(imageHeight))
	for j := 0; j < imageHeight; j++ {
		for i := 0; i < imageWidth; i++ {
			u := float64(i) / (imageWidth - 1)
			v := float64(j) / float64(imageHeight-1)
			ray := Ray{origin, lowerLeftCorner.Add(horizontal.Scale(u)).Add(vertical.Scale(v)).SubPoint(origin)}
			img.Set(i, imageHeight-j-1, rayColor(ray).NRGBA())
		}
		bar.Add(1)
	}

	log.Print("Rendered, encoding and saving...")

	f, err := os.Create("rays.png")
	if err != nil {
		log.Fatal(err)
	}

	encoder := png.Encoder{CompressionLevel: png.BestCompression}
	if err := encoder.Encode(f, img); err != nil {
		f.Close()
		log.Fatal(err)
	}

	if err := f.Close(); err != nil {
		log.Fatal(err)
	}

	log.Print("Done!")
}
