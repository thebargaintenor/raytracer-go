package worldbuilder

import (
	"math/rand"

	"github.com/thebargaintenor/raytracer-go/engine"
)

// CreateWorld builds the original scene from tutorial with 3 spheres
func CreateWorld() *engine.Scene {
	world := engine.Scene{}

	world = append(world,
		engine.Sphere{
			Center:   engine.Vec3{0.0, 0.0, -1.0},
			Radius:   0.5,
			Material: engine.Lambertian{Albedo: &engine.Color{0.3, 0.3, 0.8}},
		},
		engine.Sphere{
			Center:   engine.Vec3{0.0, -100.5, -1.0},
			Radius:   100.0,
			Material: engine.Lambertian{Albedo: &engine.Color{0.8, 0.8, 0.0}},
		},
		engine.Sphere{
			Center: engine.Vec3{1.0, 0.0, -1.0},
			Radius: 0.5,
			Material: engine.Metal{
				Albedo:    &engine.Color{0.8, 0.6, 0.2},
				Fuzziness: 1.0,
			},
		},
		engine.Sphere{
			Center:   engine.Vec3{-1.0, 0.0, -1.0},
			Radius:   0.5,
			Material: engine.Dielectric{RefractiveIndex: 1.5},
		},
		engine.Sphere{
			Center:   engine.Vec3{-1.0, 0.0, -1.0},
			Radius:   -0.45,
			Material: engine.Dielectric{RefractiveIndex: 1.5},
		},
	)

	return &world
}

// RandomScene recreates scene from tutorial book cover
func RandomScene() *engine.Scene {
	scene := engine.Scene{}

	scene = append(scene,
		engine.Sphere{
			Center:   engine.Vec3{0.0, -1000.0, 0.0},
			Radius:   1000.0,
			Material: engine.Lambertian{Albedo: &engine.Color{0.5, 0.5, 0.5}},
		},
	)

	for a := -11; a < 11; a++ {
		for b := -11; b < 11; b++ {
			materialChoice := rand.Float64()
			center := engine.Vec3{
				float64(a) + 0.9*rand.Float64(),
				0.2,
				float64(b) + 0.9*rand.Float64(),
			}

			if center.Sub(engine.Vec3{4.0, 0.2, 0.0}).Magnitude() > 0.9 {
				if materialChoice < 0.8 { // diffuse
					scene = append(scene,
						engine.Sphere{
							Center: center,
							Radius: 0.2,
							Material: engine.Lambertian{
								Albedo: &engine.Color{
									rand.Float64() * rand.Float64(),
									rand.Float64() * rand.Float64(),
									rand.Float64() * rand.Float64(),
								},
							},
						},
					)
				} else if materialChoice < 0.95 { // metal
					scene = append(scene,
						engine.Sphere{
							Center: center,
							Radius: 0.2,
							Material: engine.Metal{
								Albedo: &engine.Color{
									0.5 * (1 + rand.Float64()),
									0.5 * (1 + rand.Float64()),
									0.5 * (1 + rand.Float64()),
								},
								Fuzziness: 0.5 * rand.Float64(),
							},
						},
					)
				} else { // glass
					scene = append(scene,
						engine.Sphere{
							Center: center,
							Radius: 0.2,
							Material: engine.Dielectric{
								RefractiveIndex: 1.5,
							},
						},
					)
				}
			}
		}
	}

	scene = append(scene,
		engine.Sphere{
			Center:   engine.Vec3{0.0, 1.0, 0.0},
			Radius:   1.0,
			Material: engine.Dielectric{RefractiveIndex: 1.5},
		},
		engine.Sphere{
			Center:   engine.Vec3{-4.0, 1.0, 0.0},
			Radius:   1.0,
			Material: engine.Lambertian{Albedo: &engine.Color{0.4, 0.2, 0.1}},
		},
		engine.Sphere{
			Center: engine.Vec3{4.0, 1.0, 0.0},
			Radius: 1.0,
			Material: engine.Metal{
				Albedo:    &engine.Color{0.7, 0.6, 0.5},
				Fuzziness: 0.5,
			},
		},
	)

	return &scene
}
