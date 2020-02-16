package engine

// Ray is a vector with a defined origin point
type Ray struct {
	Origin    Vec3
	Direction Vec3
}

// CreateRay returns a new Ray
func CreateRay(origin Vec3, direction Vec3) Ray {
	return Ray{
		Origin:    origin,
		Direction: direction}
}

// PointAt returns point along ray at given time
func (r Ray) PointAt(time float64) Vec3 {
	return r.Origin.Add(r.Direction.ScalarMult(time))
}
