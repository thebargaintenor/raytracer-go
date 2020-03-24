package engine

// Metal describes a reflective material
type Metal struct {
	Albedo    *Color
	Fuzziness float64
}

// Metal implements Material
var _ Material = &Metal{}

// Scatter describes ray interacting with a reflective material
func (m Metal) Scatter(incident *Ray, h *Hit) (*Color, *Ray, bool) {
	fuzz := m.Fuzziness
	if fuzz >= 1.0 {
		fuzz = 1.0
	} else if fuzz < 0.0 {
		fuzz = 0.0
	}

	reflected := Reflect(incident.Direction.Unit(), h.Normal)
	scattered := Ray{
		Origin:    h.Point,
		Direction: reflected.Add(RandomInUnitSphere().ScalarMult(fuzz))}

	// dot product determines if reflection was internal to surface
	return m.Albedo, &scattered, scattered.Direction.Dot(h.Normal) > 0.0
}
