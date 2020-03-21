package engine

import (
	"math"
)

// Sphere is defined by center and radius
type Sphere struct {
	Center   Vec3
	Radius   float64
	Material Material
}

// a bit odd of a convention, but forces the compiler to check
// that the code that satisfies the interface exists somewhere
var _ Hittable = &Sphere{}

// Hit yields color where ray intersects sphere
func (s Sphere) Hit(r *Ray, tmin float64, tmax float64) (*Hit, bool) {
	oc := r.Origin.Sub(s.Center)
	a := r.Direction.Dot(r.Direction)
	b := oc.Dot(r.Direction)
	c := oc.Dot(oc) - s.Radius*s.Radius
	discriminant := b*b - a*c

	if discriminant < 0.0 {
		return nil, false
	}

	t := (-b - math.Sqrt(discriminant)) / a
	if t < tmax && t > tmin {
		return s.createHitRecord(r, t), true
	}

	t = (-b + math.Sqrt(discriminant)) / a
	if t < tmax && t > tmin {
		return s.createHitRecord(r, t), true
	}

	return nil, false
}

func (s Sphere) createHitRecord(r *Ray, t float64) *Hit {
	location := r.PointAt(t)
	normal := location.Sub(s.Center).Unit()

	return &Hit{
		T:        t,
		Point:    location,
		Normal:   normal,
		Material: s.Material}
}
