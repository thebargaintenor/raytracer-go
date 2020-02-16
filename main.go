package main

import (
	"fmt"
	"github.com/thebargaintenor/raytracer-go/engine/color"
)

func main() {
	fmt.Print(createPpm())
}

type ppmPixel [3]unit8

func createPpm() string {
	xres := 200
	yres := 100
	ppm := fmt.Sprintf("P3\n%d %d\n255\n", xres, yres)

	for y := yres - 1; y >= 0; y-- {
		for x := 0; x < xres; x++ {
			rgb := Color{
				float64(x) / float64(xres),
				float64(y) / float64(yres),
				0.2
			}

			pixel := ppmPixel{
				unit8(255.99 * rgb[0]),
				unit8(255.99 * rgb[1]),
				unit8(255.99 * rgb[2])
			}

			ppm += fmt.Sprintf("%d %d %d\n", pixel)
		}
	}

	return ppm
}
