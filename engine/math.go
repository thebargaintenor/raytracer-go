package engine

import (
	"math/rand"
)

// RandomInUnitSphere returns a point within a unit sphere centered at origin
func RandomInUnitSphere() Vec3 {
	var point Vec3
	// this concerns me because it only returns probabilistically
	for {
		// select point in unit cube about origin where each axis
		// in (-1, 1)
		point = Vec3{rand.Float64(), rand.Float64(), rand.Float64()}.
			ScalarMult(2.0).
			Sub(Vec3{1.0, 1.0, 1.0})

		// point is in sphere if distance to center < 1
		if point.SquareMagnitude() < 1.0 {
			return point
		}
	}
}

// Reflect a vector relative to plane normal
func Reflect(vector, normal Vec3) Vec3 {
	return vector.Sub(normal.ScalarMult(vector.Dot(normal) * 2))
}
