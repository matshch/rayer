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

func rayColor(r Ray, world Hitter, depth uint) Color {
	if depth == 0 {
		return Color{}
	}
	hit := world.Hit(r, 0.001, math.Inf(1))
	if hit != nil {
		newRay, attenuation := hit.Material.Scatter(r, *hit)
		if newRay != nil {
			return attenuation.Permute(rayColor(*newRay, world, depth-1))
		}
		return Color{}
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
	const maxDepth = 50

	var world []Hitter
	materialGround := Lambertian{Color{R: 0.8, G: 0.8, B: 0.0}}
	materialCenter := Lambertian{Color{R: 0.1, G: 0.2, B: 0.5}}
	materialLeft := Dielectric{1.5}
	materialRight := Metal{Color{R: 0.8, G: 0.6, B: 0.2}, 1.0}
	world = append(world, Sphere{Point{0.0, -100.5, -1.0}, 100.0, materialGround})
	world = append(world, Sphere{Point{0.0, 0.0, -1.0}, 0.5, materialCenter})
	world = append(world, Sphere{Point{-1.0, 0.0, -1.0}, 0.5, materialLeft})
	world = append(world, Sphere{Point{-1.0, 0.0, -1.0}, -0.4, materialLeft})
	world = append(world, Sphere{Point{1.0, 0.0, -1.0}, 0.5, materialRight})

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
				color.LazyBlend(rayColor(ray, HitterSlice(world), maxDepth))
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
