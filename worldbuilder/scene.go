package worldbuilder

import (
	"github.com/thebargaintenor/raytracer-go/engine"
)

func createWorld() *engine.Scene {
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
