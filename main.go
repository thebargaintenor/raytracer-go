package main

import (
	"fmt"
	"math"

	"github.com/thebargaintenor/raytracer-go/engine"
)

func main() {
	fmt.Print(createPpm())
}

func createPpm() string {
	xres := 200
	yres := 100
	ppm := fmt.Sprintf("P3\n%d %d\n255\n", xres, yres)

	lowerLeftCorner := engine.Vec3{-2.0, -1.0, -1.0}
	horizontal := engine.Vec3{4.0, 0.0, 0.0}
	vertical := engine.Vec3{0.0, 2.0, 0.0}
	origin := engine.Vec3{0.0, 0.0, 0.0}

	scene := engine.Scene{
		engine.Sphere{
			Center: engine.Vec3{0.0, 0.0, -1.0},
			Radius: 0.5},
		engine.Sphere{
			Center: engine.Vec3{0.0, -100.5, -1.0},
			Radius: 100.0},
	}

	for y := yres - 1; y >= 0; y-- {
		for x := 0; x < xres; x++ {
			u := float64(x) / float64(xres)
			v := float64(y) / float64(yres)

			ray := engine.Ray{
				Origin:    origin,
				Direction: (lowerLeftCorner.Add(horizontal.ScalarMult(u)).Add(vertical.ScalarMult(v)))}

			rgb := color(&ray, &scene)

			ppm += fmt.Sprintln(formatPpmPixel(rgb))
		}
	}

	return ppm
}

func color(r *engine.Ray, scene *engine.Scene) engine.Color {
	if hit, success := scene.Hit(r, 0.0, math.MaxFloat64); success {
		return engine.Color{
			hit.Normal.X() + 1.0,
			hit.Normal.Y() + 1.0,
			hit.Normal.Z() + 1.0,
		}.ScalarMult(0.5)
	}

	// scene background
	unitDirection := r.Direction.Unit()
	t := 0.5 * (unitDirection.Y() + 1.0)
	startColor := engine.Color{1.0, 1.0, 1.0}
	endColor := engine.Color{0.5, 0.7, 1.0}
	// linear interpolation being color stops
	return startColor.ScalarMult(1.0 - t).Add(endColor.ScalarMult(t))
}

func formatPpmPixel(rgb engine.Color) string {
	r := uint8(255.99 * rgb[0])
	g := uint8(255.99 * rgb[1])
	b := uint8(255.99 * rgb[2])

	return fmt.Sprintf("%d %d %d", r, g, b)
}
