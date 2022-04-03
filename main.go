package main

import (
	"fmt"
	"math"
)

func main() {
	const image_width = 256
	const image_height = 256
	const max_value = 255

	fmt.Println("P3")
	fmt.Println(image_width, image_height)
	fmt.Println(max_value)
	for j := image_height - 1; j > 0; j-- {
		for i := 0; i < image_width; i++ {
			r := float64(i) / (image_width - 1)
			g := float64(j) / (image_height - 1)
			const b = 0.25

			ir := uint8(math.Round(max_value * r))
			ig := uint8(math.Round(max_value * g))
			ib := uint8(math.Round(max_value * b))
			fmt.Println(ir, ig, ib)
		}
	}
}
