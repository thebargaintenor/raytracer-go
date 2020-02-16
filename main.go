package main

import (
	"fmt"

	"github.com/thebargaintenor/raytracer-go/engine"
)

func main() {
	fmt.Print(createPpm())
}

func createPpm() string {
	xres := 200
	yres := 100
	ppm := fmt.Sprintf("P3\n%d %d\n255\n", xres, yres)

	for y := yres - 1; y >= 0; y-- {
		for x := 0; x < xres; x++ {
			rgb := engine.Color{
				float64(x) / float64(xres),
				float64(y) / float64(yres),
				0.2}

			ppm += fmt.Sprintln(formatPpmPixel(rgb))
		}
	}

	return ppm
}

func formatPpmPixel(rgb engine.Color) string {
	r := uint8(255.99 * rgb[0])
	g := uint8(255.99 * rgb[1])
	b := uint8(255.99 * rgb[2])

	return fmt.Sprintf("%d %d %d\n", r, g, b)
}
