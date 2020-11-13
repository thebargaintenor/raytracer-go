package engine

import (
	"math"
	"math/rand"
)

// RandomInUnitSphere returns a point within a unit sphere centered at origin
func RandomInUnitSphere() Vec3 {
	var (
		rho   = rand.Float64()
		theta = 2.0 * math.Pi * rand.Float64()
		phi   = math.Pi * rand.Float64()
		x     = rho * math.Sin(phi) * math.Cos(theta)
		y     = rho * math.Sin(phi) * math.Sin(theta)
		z     = rho * math.Cos(phi)
	)

	return Vec3{x, y, z}
}

// RandomInUnitDisk returns random point in unit circle about origin
func RandomInUnitDisk() Vec3 {
	theta := 2.0 * math.Pi * rand.Float64()
	r := rand.Float64()

	return Vec3{r * math.Cos(theta), r * math.Sin(theta), 0.0}
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
