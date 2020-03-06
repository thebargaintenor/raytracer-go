package engine

// Scene is a collection of ray-hittable objects
type Scene []Hittable

var _ Hittable = &Scene{}

// Hit determines the nearest collision for a scene
func (s Scene) Hit(r *Ray, tmin float64, tmax float64) (*Hit, bool) {
	var (
		hit     *Hit
		isHit   = false
		maxDist = tmax
	)

	for _, object := range s {
		if h, ok := object.Hit(r, tmin, maxDist); ok {
			hit = h
			isHit = true
			maxDist = hit.T
		}
	}

	return hit, isHit
}
