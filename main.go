package main

import (
	"fmt"
	"math"

	"github.com/thebargaintenor/raytracer-go/engine"
	"github.com/thebargaintenor/raytracer-go/worldbuilder"
)

func main() {
	fmt.Print(createPpm())
}

func createPpm() string {
	var (
		xres          = 640
		yres          = 480
		samples       = 100
		lookFrom      = engine.Vec3{13.0, 2.0, 3.0}
		lookAt        = engine.Vec3{0.0, 0.0, 0.0}
		vup           = engine.Vec3{0.0, 1.0, 0.0}
		fov           = 30.0
		aspect        = float64(xres) / float64(yres)
		focusDistance = lookFrom.Sub(lookAt).Magnitude()
		aperture      = 0.1
	)

	ppm := fmt.Sprintf("P3\n%d %d\n255\n", xres, yres)

	world := worldbuilder.RandomScene()
	camera := engine.CreateCamera(
		lookFrom,
		lookAt,
		vup,
		fov,
		aspect,
		aperture,
		focusDistance,
	)

	renderContext := &RenderContext{
		Width:      xres,
		Height:     yres,
		SampleSize: samples,
		Camera:     &camera,
		Scene:      world,
	}

	bitmap := renderParallel(renderContext)

	for _, c := range *bitmap.Data {
		ppm += fmt.Sprintln(formatPpmPixel(c))
	}

	return ppm
}

func formatPpmPixel(rgb engine.Color) string {
	// conversion needs protection against overflows
	r := uint8(math.Min(255.99*rgb[0], 255.0))
	g := uint8(math.Min(255.99*rgb[1], 255.0))
	b := uint8(math.Min(255.99*rgb[2], 255.0))

	return fmt.Sprintf("%d %d %d", r, g, b)
}
