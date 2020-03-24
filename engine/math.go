package engine

import (
	"math"
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

// Refract handles the bending of a ray at the interface of two materials
func Refract(vector, normal Vec3, refactiveIdxRatio float64) (Vec3, bool) {
	uv := vector.Unit()
	dtheta := uv.Dot(normal)
	discriminant := 1.0 - refactiveIdxRatio*refactiveIdxRatio*(1-dtheta*dtheta)

	if discriminant > 0.0 {
		refracted := uv.Sub(normal.ScalarMult(dtheta)).
			ScalarMult(refactiveIdxRatio).
			Sub(normal.ScalarMult(math.Sqrt(discriminant)))

		return refracted, true
	}

	return Vec3{}, false
}

// Schlick approximates dielectric material reflectivity
func Schlick(cosine, refractiveIndex float64) float64 {
	r0 := (1 - refractiveIndex) / (1 + refractiveIndex)
	r0 = r0 * r0
	return r0 + (1.0-r0)*math.Pow(1-cosine, 5)
}
