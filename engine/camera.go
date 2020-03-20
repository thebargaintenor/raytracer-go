package engine

// Camera abstracts scene viewport
type Camera struct {
	Origin     Vec3
	BottomLeft Vec3
	Horizontal Vec3
	Vertical   Vec3
}

// GetRay casts rays from viewport into scene
func (c Camera) GetRay(u float64, v float64) Ray {
	return Ray{
		Origin: c.Origin,
		Direction: c.BottomLeft.
			Add(c.Horizontal.ScalarMult(u)).
			Add(c.Vertical.ScalarMult(v)).
			Sub(c.Origin),
	}
}
