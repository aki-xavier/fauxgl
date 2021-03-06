package fauxgl

import (
	"math"
	"math/rand"
)

// Vector :
type Vector struct {
	X, Y, Z float64
}

// V :
func V(x, y, z float64) Vector {
	return Vector{x, y, z}
}

// RandomUnitVector :
func RandomUnitVector() Vector {
	for {
		x := rand.Float64()*2 - 1
		y := rand.Float64()*2 - 1
		z := rand.Float64()*2 - 1
		if x*x+y*y+z*z > 1 {
			continue
		}
		return Vector{x, y, z}.Normalize()
	}
}

// VectorW :
func (a Vector) VectorW() VectorW {
	return VectorW{a.X, a.Y, a.Z, 1}
}

// IsDegenerate :
func (a Vector) IsDegenerate() bool {
	nan := math.IsNaN(a.X) || math.IsNaN(a.Y) || math.IsNaN(a.Z)
	inf := math.IsInf(a.X, 0) || math.IsInf(a.Y, 0) || math.IsInf(a.Z, 0)
	return nan || inf
}

// Length :
func (a Vector) Length() float64 {
	return math.Sqrt(a.X*a.X + a.Y*a.Y + a.Z*a.Z)
}

// Less :
func (a Vector) Less(b Vector) bool {
	if a.X != b.X {
		return a.X < b.X
	}
	if a.Y != b.Y {
		return a.Y < b.Y
	}
	return a.Z < b.Z
}

// Distance :
func (a Vector) Distance(b Vector) float64 {
	return a.Sub(b).Length()
}

// LengthSquared :
func (a Vector) LengthSquared() float64 {
	return a.X*a.X + a.Y*a.Y + a.Z*a.Z
}

// DistanceSquared :
func (a Vector) DistanceSquared(b Vector) float64 {
	return a.Sub(b).LengthSquared()
}

// Lerp :
func (a Vector) Lerp(b Vector, t float64) Vector {
	return a.Add(b.Sub(a).MulScalar(t))
}

// LerpDistance :
func (a Vector) LerpDistance(b Vector, d float64) Vector {
	return a.Add(b.Sub(a).Normalize().MulScalar(d))
}

// Dot :
func (a Vector) Dot(b Vector) float64 {
	return a.X*b.X + a.Y*b.Y + a.Z*b.Z
}

// Cross :
func (a Vector) Cross(b Vector) Vector {
	x := a.Y*b.Z - a.Z*b.Y
	y := a.Z*b.X - a.X*b.Z
	z := a.X*b.Y - a.Y*b.X
	return Vector{x, y, z}
}

// Normalize :
func (a Vector) Normalize() Vector {
	r := 1 / math.Sqrt(a.X*a.X+a.Y*a.Y+a.Z*a.Z)
	return Vector{a.X * r, a.Y * r, a.Z * r}
}

// Negate :
func (a Vector) Negate() Vector {
	return Vector{-a.X, -a.Y, -a.Z}
}

// Abs :
func (a Vector) Abs() Vector {
	return Vector{math.Abs(a.X), math.Abs(a.Y), math.Abs(a.Z)}
}

// Add :
func (a Vector) Add(b Vector) Vector {
	return Vector{a.X + b.X, a.Y + b.Y, a.Z + b.Z}
}

// Sub :
func (a Vector) Sub(b Vector) Vector {
	return Vector{a.X - b.X, a.Y - b.Y, a.Z - b.Z}
}

// Mul :
func (a Vector) Mul(b Vector) Vector {
	return Vector{a.X * b.X, a.Y * b.Y, a.Z * b.Z}
}

// Div :
func (a Vector) Div(b Vector) Vector {
	return Vector{a.X / b.X, a.Y / b.Y, a.Z / b.Z}
}

// Mod :
func (a Vector) Mod(b Vector) Vector {
	// as implemented in GLSL
	x := a.X - b.X*math.Floor(a.X/b.X)
	y := a.Y - b.Y*math.Floor(a.Y/b.Y)
	z := a.Z - b.Z*math.Floor(a.Z/b.Z)
	return Vector{x, y, z}
}

// AddScalar :
func (a Vector) AddScalar(b float64) Vector {
	return Vector{a.X + b, a.Y + b, a.Z + b}
}

// SubScalar :
func (a Vector) SubScalar(b float64) Vector {
	return Vector{a.X - b, a.Y - b, a.Z - b}
}

// MulScalar :
func (a Vector) MulScalar(b float64) Vector {
	return Vector{a.X * b, a.Y * b, a.Z * b}
}

// DivScalar :
func (a Vector) DivScalar(b float64) Vector {
	return Vector{a.X / b, a.Y / b, a.Z / b}
}

// Min :
func (a Vector) Min(b Vector) Vector {
	return Vector{math.Min(a.X, b.X), math.Min(a.Y, b.Y), math.Min(a.Z, b.Z)}
}

// Max :
func (a Vector) Max(b Vector) Vector {
	return Vector{math.Max(a.X, b.X), math.Max(a.Y, b.Y), math.Max(a.Z, b.Z)}
}

// Floor :
func (a Vector) Floor() Vector {
	return Vector{math.Floor(a.X), math.Floor(a.Y), math.Floor(a.Z)}
}

// Ceil :
func (a Vector) Ceil() Vector {
	return Vector{math.Ceil(a.X), math.Ceil(a.Y), math.Ceil(a.Z)}
}

// Round :
func (a Vector) Round() Vector {
	return a.RoundPlaces(0)
}

// RoundPlaces :
func (a Vector) RoundPlaces(n int) Vector {
	x := RoundPlaces(a.X, n)
	y := RoundPlaces(a.Y, n)
	z := RoundPlaces(a.Z, n)
	return Vector{x, y, z}
}

// MinComponent :
func (a Vector) MinComponent() float64 {
	return math.Min(math.Min(a.X, a.Y), a.Z)
}

// MaxComponent :
func (a Vector) MaxComponent() float64 {
	return math.Max(math.Max(a.X, a.Y), a.Z)
}

// Reflect :
func (a Vector) Reflect(n Vector) Vector {
	return a.Sub(n.MulScalar(2 * n.Dot(a)))
}

// Perpendicular :
func (a Vector) Perpendicular() Vector {
	if a.X == 0 && a.Y == 0 {
		if a.Z == 0 {
			return Vector{}
		}
		return Vector{0, 1, 0}
	}
	return Vector{-a.Y, a.X, 0}.Normalize()
}

// SegmentDistance :
func (a Vector) SegmentDistance(v Vector, w Vector) float64 {
	l2 := v.DistanceSquared(w)
	if l2 == 0 {
		return a.Distance(v)
	}
	t := a.Sub(v).Dot(w.Sub(v)) / l2
	if t < 0 {
		return a.Distance(v)
	}
	if t > 1 {
		return a.Distance(w)
	}
	return v.Add(w.Sub(v).MulScalar(t)).Distance(a)
}

// VectorW :
type VectorW struct {
	X, Y, Z, W float64
}

// Vector :
func (a VectorW) Vector() Vector {
	return Vector{a.X, a.Y, a.Z}
}

// Outside :
func (a VectorW) Outside() bool {
	x, y, z, w := a.X, a.Y, a.Z, a.W
	return x < -w || x > w || y < -w || y > w || z < -w || z > w
}

// Dot :
func (a VectorW) Dot(b VectorW) float64 {
	return a.X*b.X + a.Y*b.Y + a.Z*b.Z + a.W*b.W
}

// Add :
func (a VectorW) Add(b VectorW) VectorW {
	return VectorW{a.X + b.X, a.Y + b.Y, a.Z + b.Z, a.W + b.W}
}

// Sub :
func (a VectorW) Sub(b VectorW) VectorW {
	return VectorW{a.X - b.X, a.Y - b.Y, a.Z - b.Z, a.W - b.W}
}

// MulScalar :
func (a VectorW) MulScalar(b float64) VectorW {
	return VectorW{a.X * b, a.Y * b, a.Z * b, a.W * b}
}

// DivScalar :
func (a VectorW) DivScalar(b float64) VectorW {
	return VectorW{a.X / b, a.Y / b, a.Z / b, a.W / b}
}
