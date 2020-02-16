package engine

// Color is an array of RGB values
type Color [3]float64

// R returns red channel
func (c Color) R() float64 {
	return c[0]
}

// G returns green channel
func (c Color) G() float64 {
	return c[1]
}

// B returns blue channel
func (c Color) B() float64 {
	return c[2]
}

// ScalarMult returns color multiplied by scalar
func (c Color) ScalarMult(scalar float64) Color {
	return Color{c[0] * scalar, c[1] * scalar, c[2] * scalar}
}

// ScalarDiv returns color divided by scalar
func (c Color) ScalarDiv(scalar float64) Color {
	return Color{c[0] / scalar, c[1] / scalar, c[2] / scalar}
}

// Add returns sum of colors
func (c Color) Add(c2 Color) Color {
	return Color{c[0] + c2[0], c[1] + c2[1], c[2] + c2[2]}
}

// Mult returns product of colors
func (c Color) Mult(c2 Color) Color {
	return Color{c[0] * c2[0], c[1] * c2[1], c[2] * c2[2]}
}
