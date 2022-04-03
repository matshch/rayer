package main

import (
	"image"
	"image/png"
	"log"
	"math"
	"math/rand"
	"os"

	"github.com/schollz/progressbar/v3"
)

func rayColor(r Ray, world Hitter) Color {
	hit := world.Hit(r, 0, math.Inf(1))
	if hit != nil {
		return Color{
			R: .5*hit.Normal.X + .5,
			G: .5*hit.Normal.Y + .5,
			B: .5*hit.Normal.Z + .5,
		}
	}
	unit := r.Direction.Unit()
	t := 0.5 * (unit.Y + 1.0)
	return Color{R: 1, G: 1, B: 1}.Blend(t, Color{R: 0.5, G: 0.7, B: 1})
}

func main() {
	log.Print("Starting rendering...")

	const aspectRatio = 16. / 9.
	const imageWidth = 400
	imageHeight := int(math.Round(imageWidth / aspectRatio))
	const samplesPerPixel = 100

	var world []Hitter
	world = append(world, Sphere{Point{0, 0, -1}, 0.5})
	world = append(world, Sphere{Point{0, -100.5, -1}, 100})

	camera := NewCamera(aspectRatio)

	img := image.NewNRGBA(image.Rect(0, 0, imageWidth, imageHeight))
	bar := progressbar.Default(int64(imageHeight))
	for j := 0; j < imageHeight; j++ {
		for i := 0; i < imageWidth; i++ {
			var color Color
			for s := 0; s < samplesPerPixel; s++ {
				u := (float64(i) + rand.Float64()) / (imageWidth - 1)
				v := (float64(j) + rand.Float64()) / float64(imageHeight-1)
				ray := camera.Ray(u, v)
				color.LazyBlend(rayColor(ray, HitterSlice(world)))
			}
			img.Set(i, imageHeight-j-1, color.NRGBA())
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
