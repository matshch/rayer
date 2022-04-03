package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"math"
	"os"
)

func main() {
	const image_width = 256
	const image_height = 256
	const max_value = 255

	img := image.NewNRGBA(image.Rect(0, 0, image_width, image_height))
	for j := 0; j < image_height; j++ {
		for i := 0; i < image_width; i++ {
			r := float64(i) / (image_width - 1)
			g := float64(j) / (image_height - 1)
			const b = 0.25

			ir := uint8(math.Round(max_value * r))
			ig := uint8(math.Round(max_value * g))
			ib := uint8(math.Round(max_value * b))
			img.Set(i, image_height - j - 1, color.NRGBA{
				R: ir,
				G: ig,
				B: ib,
				A: 255,
			})
		}
	}

	f, err := os.Create("rays.png")
	if err != nil {
		log.Fatal(err)
	}

	if err := png.Encode(f, img); err != nil {
		f.Close()
		log.Fatal(err)
	}

	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}
