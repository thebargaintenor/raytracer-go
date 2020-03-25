package engine

import "math"

// Camera abstracts scene viewport
type Camera struct {
	Origin     Vec3
	BottomLeft Vec3
	Horizontal Vec3
	Vertical   Vec3
	U          Vec3
	V          Vec3
	W          Vec3
	LensRadius float64
}

// CreateCamera builds a new camera with the given vertical FOV and aspect ratio
func CreateCamera(
	lookFrom Vec3,
	lookAt Vec3,
	vup Vec3,
	vfov float64,
	aspect float64,
	aperture float64,
	focusDistance float64,
) Camera {
	lensRadius := aperture / 2
	theta := vfov * math.Pi / 180
	halfHeight := math.Tan(theta / 2)
	halfWidth := aspect * halfHeight

	origin := lookFrom
	w := lookFrom.Sub(lookAt).Unit()
	u := vup.Cross(w).Unit()
	v := w.Cross(u)
	bottomLeft := origin.
		Sub(u.ScalarMult(halfWidth * focusDistance)).
		Sub(v.ScalarMult(halfHeight * focusDistance)).
		Sub(w.ScalarMult(focusDistance))

	return Camera{
		Origin:     origin,
		BottomLeft: bottomLeft,
		Horizontal: u.ScalarMult(2 * halfWidth * focusDistance),
		Vertical:   v.ScalarMult(2 * halfHeight * focusDistance),
		U:          u,
		V:          v,
		W:          w,
		LensRadius: lensRadius,
	}
}

// GetRay casts rays from viewport into scene
func (c Camera) GetRay(u float64, v float64) Ray {
	pointOnLens := RandomInUnitDisk().ScalarMult(c.LensRadius)
	offset := c.U.ScalarMult(pointOnLens.X()).Add(c.V.ScalarMult(pointOnLens.Y()))

	return Ray{
		Origin: c.Origin.Add(offset),
		Direction: c.BottomLeft.
			Add(c.Horizontal.ScalarMult(u)).
			Add(c.Vertical.ScalarMult(v)).
			Sub(c.Origin).
			Sub(offset),
	}
}
