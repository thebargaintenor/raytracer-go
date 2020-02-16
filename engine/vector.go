package engine

import "math"

// Vec3 is an array of XYZ coordinates
type Vec3 [3]float64

// X coordinate of vector
func (v Vec3) X() float64 {
	return v[0]
}

// Y coordinate of vector
func (v Vec3) Y() float64 {
	return v[1]
}

// Z coordinate of vector
func (v Vec3) Z() float64 {
	return v[2]
}

// Inverse returns vector of equal magnitude in opposite direction
func (v Vec3) Inverse() Vec3 {
	return Vec3{-v[0], -v[1], -v[2]}
}

// Magnitude of the vector
func (v Vec3) Magnitude() float64 {
	return math.Sqrt(v[0]*v[0] + v[1]*v[1] + v[2]*v[2])
}

// SquareMagnitude is what it says on the tin
func (v Vec3) SquareMagnitude() float64 {
	return v[0]*v[0] + v[1]*v[1] + v[2]*v[2]
}

// ScalarMult returns vector multiplied by scalar
func (v Vec3) ScalarMult(scalar float64) Vec3 {
	return Vec3{v[0] * scalar, v[1] * scalar, v[2] * scalar}
}

// ScalarDiv returns vector divided by scalar
func (v Vec3) ScalarDiv(scalar float64) Vec3 {
	return Vec3{v[0] / scalar, v[1] / scalar, v[2] / scalar}
}

// Unit returns vector with same direction but magnitude 1
func (v Vec3) Unit() Vec3 {
	return v.ScalarDiv(v.Magnitude())
}

// Add returns sum (v + v2)
func (v Vec3) Add(v2 Vec3) Vec3 {
	return Vec3{v[0] + v2[0], v[1] + v2[1], v[2] + v2[2]}
}

// Sub returns difference (v - v2)
func (v Vec3) Sub(v2 Vec3) Vec3 {
	return Vec3{v[0] - v2[0], v[1] - v2[1], v[2] + v2[2]}
}

// Dot product of v, v2
func (v Vec3) Dot(v2 Vec3) float64 {
	return v[0]*v2[0] + v[1]*v2[1] + v[2]*v2[2]
}

// Cross product of v, v2
func (v Vec3) Cross(v2 Vec3) Vec3 {
	return Vec3{
		v[1]*v2[2] - v[2]*v2[1],
		v[2]*v2[0] - v[0]*v2[2],
		v[0]*v2[1] - v[1]*v2[0]}
}

// Will these be for somewhere else? 3D vectors don't really
// use elementwise products...

// Mult returns product (v * v2)
func (v Vec3) Mult(v2 Vec3) Vec3 {
	return Vec3{v[0] * v2[0], v[1] * v2[1], v[2] * v2[2]}
}

// Div returns quotient (v / v2)
func (v Vec3) Div(v2 Vec3) Vec3 {
	return Vec3{v[0] / v2[0], v[1] / v2[1], v[2] / v2[2]}
}
