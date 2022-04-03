package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"math"
	"os"

	"github.com/schollz/progressbar/v3"
)

func main() {
	log.Print("Starting rendering...")

	const image_width = 256
	const image_height = 256

	img := image.NewNRGBA(image.Rect(0, 0, image_width, image_height))
	bar := progressbar.Default(image_height)
	for j := 0; j < image_height; j++ {
		for i := 0; i < image_width; i++ {
			r := float64(i) / (image_width - 1)
			g := float64(j) / (image_height - 1)
			const b = 0.25

			img.Set(i, image_height - j - 1, color.NRGBA{
				R: uint8(math.Round(math.MaxUint8 * r)),
				G: uint8(math.Round(math.MaxUint8 * g)),
				B: uint8(math.Round(math.MaxUint8 * b)),
				A: 255,
			})
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
