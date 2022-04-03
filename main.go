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

func randomScene() []Hitter {
	var world []Hitter
	groundMaterial := Lambertian{NewColor(.5, .5, .5)}
	world = append(world, Sphere{Point{0, -1000, 0}, 1000, groundMaterial})

	for a := -11; a < 11; a++ {
		for b := -11; b < 11; b++ {
			materialChooser := rand.Float64()
			center := Point{float64(a) + .9*rand.Float64(), .2, float64(b) + .9*rand.Float64()}

			if (center.SubPoint(Point{4, .2, 0})).Len() > .9 {
				if materialChooser < .8 {
					albedo := RandomColor().Permute(RandomColor())
					sphereMaterial := Lambertian{albedo}
					world = append(world, Sphere{center, .2, sphereMaterial})
				} else if materialChooser < .95 {
					albedo := RandomRangeColor(.5, 1)
					fuzz := Random(0, .5)
					sphereMaterial := Metal{albedo, fuzz}
					world = append(world, Sphere{center, .2, sphereMaterial})
				} else {
					sphereMaterial := Dielectric{1.5}
					world = append(world, Sphere{center, .2, sphereMaterial})
				}
			}
		}
	}

	material1 := Dielectric{1.5}
	world = append(world, Sphere{Point{0, 1, 0}, 1, material1})

	material2 := Lambertian{NewColor(.4, .2, .1)}
	world = append(world, Sphere{Point{-4, 1, 0}, 1, material2})

	material3 := Metal{NewColor(.7, .6, .5), 0}
	world = append(world, Sphere{Point{4, 1, 0}, 1, material3})

	return world
}

func main() {
	log.Print("Starting rendering...")

	const aspectRatio = 3. / 2.
	const imageWidth = 1200
	imageHeight := int(math.Round(imageWidth / aspectRatio))
	const samplesPerPixel = 500
	const maxDepth = 50

	world := randomScene()

	lookFrom := Point{13, 2, 3}
	lookAt := Point{0, 0, 0}
	camera := NewCamera(lookFrom, lookAt, Vector{0, 1, 0},
		20, aspectRatio, 0.1, 10)

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
