package engine

import "math"

// Camera abstracts scene viewport
type Camera struct {
	Origin     Vec3
	BottomLeft Vec3
	Horizontal Vec3
	Vertical   Vec3
}

// CreateCamera builds a new camera with the given vertical FOV and aspect ratio
func CreateCamera(lookFrom Vec3, lookAt Vec3, vup Vec3, vfov float64, aspect float64) Camera {
	theta := vfov * math.Pi / 180
	halfHeight := math.Tan(theta / 2)
	halfWidth := aspect * halfHeight

	origin := lookFrom
	w := lookFrom.Sub(lookAt).Unit()
	u := vup.Cross(w).Unit()
	v := w.Cross(u)

	return Camera{
		Origin:     origin,
		BottomLeft: origin.Sub(u.ScalarMult(halfWidth)).Sub(v.ScalarMult(halfHeight)).Sub(w),
		Horizontal: u.ScalarMult(2 * halfWidth),
		Vertical:   v.ScalarMult(2 * halfHeight),
	}
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
