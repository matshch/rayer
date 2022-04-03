package main

import (
	"image"
	"image/png"
	"log"
	"math"
	"os"

	"github.com/schollz/progressbar/v3"
)

func rayColor(r Ray, world Hitter) Color {
	hit := world.Hit(r, 0, math.Inf(1))
	if hit != nil {
		return Color{
			.5*hit.Normal.X + .5,
			.5*hit.Normal.Y + .5,
			.5*hit.Normal.Z + .5,
		}
	}
	unit := r.Direction.Unit()
	t := 0.5 * (unit.Y + 1.0)
	return Color{1, 1, 1}.Blend(t, Color{0.5, 0.7, 1})
}

func main() {
	log.Print("Starting rendering...")

	const aspectRatio = 16. / 9.
	const imageWidth = 400
	imageHeight := int(math.Round(imageWidth / aspectRatio))

	var world []Hitter
	world = append(world, Sphere{Point{0, 0, -1}, 0.5})
	world = append(world, Sphere{Point{0, -100.5, -1}, 100})

	camera := NewCamera(aspectRatio)

	img := image.NewNRGBA(image.Rect(0, 0, imageWidth, imageHeight))
	bar := progressbar.Default(int64(imageHeight))
	for j := 0; j < imageHeight; j++ {
		for i := 0; i < imageWidth; i++ {
			u := float64(i) / (imageWidth - 1)
			v := float64(j) / float64(imageHeight-1)
			ray := camera.Ray(u, v)
			img.Set(i, imageHeight-j-1, rayColor(ray, HitterSlice(world)).NRGBA())
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
