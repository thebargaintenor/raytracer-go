package main

import (
	"fmt"
	"math"
	"math/rand"

	"github.com/thebargaintenor/raytracer-go/engine"
)

func main() {
	fmt.Print(createPpm())
}

func createPpm() string {
	var (
		xres            = 200
		yres            = 100
		samples         = 100
		lowerLeftCorner = engine.Vec3{-2.0, -1.0, -1.0}
		horizontal      = engine.Vec3{4.0, 0.0, 0.0}
		vertical        = engine.Vec3{0.0, 2.0, 0.0}
		origin          = engine.Vec3{0.0, 0.0, 0.0}
	)

	ppm := fmt.Sprintf("P3\n%d %d\n255\n", xres, yres)

	world := engine.Scene{
		engine.Sphere{
			Center: engine.Vec3{0.0, 0.0, -1.0},
			Radius: 0.5},
		engine.Sphere{
			Center: engine.Vec3{0.0, -100.5, -1.0},
			Radius: 100.0},
	}

	camera := engine.Camera{
		BottomLeft: lowerLeftCorner,
		Horizontal: horizontal,
		Vertical:   vertical,
		Origin:     origin,
	}

	c := engine.Color{0.0, 0.0, 0.0}
	for y := yres - 1; y >= 0; y-- {
		for x := 0; x < xres; x++ {
			for s := 0; s < samples; s++ {
				u := (float64(x) + rand.Float64()) / float64(xres)
				v := (float64(y) + rand.Float64()) / float64(yres)
				ray := camera.GetRay(u, v)

				c = c.Add(color(&ray, &world))
			}

			c = c.ScalarDiv(float64(samples))

			ppm += fmt.Sprintln(formatPpmPixel(c))
		}
	}

	return ppm
}

func color(r *engine.Ray, scene *engine.Scene) engine.Color {
	if hit, success := scene.Hit(r, 0.0, math.MaxFloat64); success {
		target := hit.Point.Add(hit.Normal).Add(engine.RandomInUnitSphere())

		return color(&engine.Ray{
			Origin:    hit.Point,
			Direction: target.Sub(hit.Point),
		}, scene).ScalarMult(0.5)
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
	// conversion needs protection against overflows
	r := uint8(math.Min(255.99*rgb[0], 255.0))
	g := uint8(math.Min(255.99*rgb[1], 255.0))
	b := uint8(math.Min(255.99*rgb[2], 255.0))

	return fmt.Sprintf("%d %d %d", r, g, b)
}
