package engine

// Material describes scattering and attenuation of ray on surface
type Material interface {
	// outputs color attenuation, reflected ray, and whether ray was scattered
	Scatter(ray *Ray, hit *Hit) (*Color, *Ray, bool)
}
