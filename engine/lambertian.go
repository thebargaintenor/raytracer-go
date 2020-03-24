package engine

// Lambertian describes a diffuse material
type Lambertian struct {
	Albedo *Color
}

// Lambertian implements Material
var _ Material = &Lambertian{}

// Scatter describes ray interacting with a diffuse material
func (l Lambertian) Scatter(incident *Ray, h *Hit) (*Color, *Ray, bool) {
	target := h.Point.Add(h.Normal).Add(RandomInUnitSphere())
	scattered := Ray{
		Origin:    h.Point,
		Direction: target.Sub(h.Point)}

	// Note: we could just as well only scatter with some probability
	// p and have attenuation be albedo/p. Your choice.
	return l.Albedo, &scattered, true
}
