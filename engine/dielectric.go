package engine

import "math/rand"

// Dielectric splits incident ray into reflected and transmitted
type Dielectric struct {
	RefractiveIndex float64
}

var _ Material = &Dielectric{}

// Scatter described scattering of a ray by a dielectric material
func (d Dielectric) Scatter(incident *Ray, h *Hit) (*Color, *Ray, bool) {
	var (
		outwardNormal      Vec3
		niOverNt           float64
		reflectProbability float64
		cosine             float64
		scattered          *Ray
	)

	reflected := Reflect(incident.Direction, h.Normal)
	attenuation := &Color{1.0, 1.0, 1.0}

	if incident.Direction.Dot(h.Normal) > 0.0 {
		outwardNormal = h.Normal.Inverse()
		niOverNt = d.RefractiveIndex
		cosine = d.RefractiveIndex * incident.Direction.Dot(h.Normal) / incident.Direction.Magnitude()
	} else {
		outwardNormal = h.Normal
		niOverNt = 1.0 / d.RefractiveIndex
		cosine = -incident.Direction.Dot(h.Normal) / incident.Direction.Magnitude()
	}

	refracted, mightReflect := Refract(incident.Direction, outwardNormal, niOverNt)

	if mightReflect {
		reflectProbability = Schlick(cosine, d.RefractiveIndex)
	} else {
		reflectProbability = 1.0
	}

	if rand.Float64() < reflectProbability {
		scattered = &Ray{Origin: h.Point, Direction: reflected}
	} else {
		scattered = &Ray{Origin: h.Point, Direction: refracted}
	}

	return attenuation, scattered, true
}
