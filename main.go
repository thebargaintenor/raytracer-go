package main

import (
	"fmt"
	"math"
	"math/rand"

	"github.com/thebargaintenor/raytracer-go/engine"
	"github.com/thebargaintenor/raytracer-go/worldbuilder"
)

func main() {
	fmt.Print(createPpm())
}

func createPpm() string {
	var (
		xres          = 300
		yres          = 150
		samples       = 100
		lookFrom      = engine.Vec3{13.0, 3.0, 2.0}
		lookAt        = engine.Vec3{0.0, 0.0, -1.0}
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

	c := engine.Color{0.0, 0.0, 0.0}
	for y := yres - 1; y >= 0; y-- {
		for x := 0; x < xres; x++ {
			for s := 0; s < samples; s++ {
				u := (float64(x) + rand.Float64()) / float64(xres)
				v := (float64(y) + rand.Float64()) / float64(yres)
				ray := camera.GetRay(u, v)

				c = c.Add(color(&ray, world, 0))
			}

			c = correctGamma(c.ScalarDiv(float64(samples)))

			ppm += fmt.Sprintln(formatPpmPixel(c))
		}
	}

	return ppm
}

func correctGamma(c engine.Color) engine.Color {
	return engine.Color{
		math.Sqrt(c[0]),
		math.Sqrt(c[1]),
		math.Sqrt(c[2]),
	}
}

func color(r *engine.Ray, scene *engine.Scene, depth uint8) engine.Color {
	if hit, success := scene.Hit(r, 0.001, math.MaxFloat64); success {
		attenuation, scatteredRay, scattered := hit.Material.Scatter(r, hit)
		if depth < 50 && scattered {
			return attenuation.Mult(color(scatteredRay, scene, depth+1))
		}

		return engine.Color{0.0, 0.0, 0.0}
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
