package engine

// Hit is a record of a ray hitting an object in scene
type Hit struct {
	T        float64
	Point    Vec3
	Normal   Vec3
	Material Material
}

// Hittable describes an object that implements ray collisions
type Hittable interface {
	Hit(r *Ray, tmin float64, tmax float64) (*Hit, bool)
}
