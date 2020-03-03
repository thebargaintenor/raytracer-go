package engine

// Sphere is defined by center and radius
type Sphere struct {
	Center Vec3
	Radius float64
}

// Hit yields color where ray intersects sphere
func (s Sphere) Hit(r *Ray) (*Color, bool) {
	oc := r.Origin.Sub(s.Center)
	a := r.Direction.Dot(r.Direction)
	b := 2.0 * oc.Dot(r.Direction)
	c := oc.Dot(oc) - s.Radius*s.Radius
	discriminant := b*b - 4.0*a*c

	if discriminant > 0 {
		return &Color{1.0, 0.0, 0.0}, true
	}

	return nil, false
}
