package fauxgl

import "math"

// Matrix :
type Matrix struct {
	X00, X01, X02, X03 float64
	X10, X11, X12, X13 float64
	X20, X21, X22, X23 float64
	X30, X31, X32, X33 float64
}

// Identity :
func Identity() Matrix {
	return Matrix{
		1, 0, 0, 0,
		0, 1, 0, 0,
		0, 0, 1, 0,
		0, 0, 0, 1}
}

// Translate :
func Translate(v Vector) Matrix {
	return Matrix{
		1, 0, 0, v.X,
		0, 1, 0, v.Y,
		0, 0, 1, v.Z,
		0, 0, 0, 1}
}

// Scale :
func Scale(v Vector) Matrix {
	return Matrix{
		v.X, 0, 0, 0,
		0, v.Y, 0, 0,
		0, 0, v.Z, 0,
		0, 0, 0, 1}
}

// Rotate :
func Rotate(v Vector, a float64) Matrix {
	v = v.Normalize()
	s := math.Sin(a)
	c := math.Cos(a)
	m := 1 - c
	return Matrix{
		m*v.X*v.X + c, m*v.X*v.Y + v.Z*s, m*v.Z*v.X - v.Y*s, 0,
		m*v.X*v.Y - v.Z*s, m*v.Y*v.Y + c, m*v.Y*v.Z + v.X*s, 0,
		m*v.Z*v.X + v.Y*s, m*v.Y*v.Z - v.X*s, m*v.Z*v.Z + c, 0,
		0, 0, 0, 1}
}

// RotateTo :
func RotateTo(a, b Vector) Matrix {
	dot := b.Dot(a)
	if dot == 1 {
		return Identity()
	} else if dot == -1 {
		return Rotate(a.Perpendicular(), math.Pi)
	} else {
		angle := math.Acos(dot)
		v := b.Cross(a).Normalize()
		return Rotate(v, angle)
	}
}

// Orient :
func Orient(position, size, up Vector, rotation float64) Matrix {
	m := Rotate(Vector{0, 0, 1}, rotation)
	m = m.Scale(size)
	m = m.RotateTo(Vector{0, 0, 1}, up)
	m = m.Translate(position)
	return m
}

// Frustum :
func Frustum(l, r, b, t, n, f float64) Matrix {
	t1 := 2 * n
	t2 := r - l
	t3 := t - b
	t4 := f - n
	return Matrix{
		t1 / t2, 0, (r + l) / t2, 0,
		0, t1 / t3, (t + b) / t3, 0,
		0, 0, (-f - n) / t4, (-t1 * f) / t4,
		0, 0, -1, 0}
}

// Orthographic :
func Orthographic(l, r, b, t, n, f float64) Matrix {
	return Matrix{
		2 / (r - l), 0, 0, -(r + l) / (r - l),
		0, 2 / (t - b), 0, -(t + b) / (t - b),
		0, 0, -2 / (f - n), -(f + n) / (f - n),
		0, 0, 0, 1}
}

// Perspective :
func Perspective(fovy, aspect, near, far float64) Matrix {
	ymax := near * math.Tan(fovy*math.Pi/360)
	xmax := ymax * aspect
	return Frustum(-xmax, xmax, -ymax, ymax, near, far)
}

// LookAt :
func LookAt(eye, center, up Vector) Matrix {
	z := eye.Sub(center).Normalize()
	x := up.Cross(z).Normalize()
	y := z.Cross(x)
	return Matrix{
		x.X, x.Y, x.Z, -x.Dot(eye),
		y.X, y.Y, y.Z, -y.Dot(eye),
		z.X, z.Y, z.Z, -z.Dot(eye),
		0, 0, 0, 1,
	}
}

// LookAtDirection :
func LookAtDirection(forward, up Vector) Matrix {
	z := forward.Normalize()
	x := up.Cross(z).Normalize()
	y := z.Cross(x)
	return Matrix{
		x.X, x.Y, x.Z, 0,
		y.X, y.Y, y.Z, 0,
		z.X, z.Y, z.Z, 0,
		0, 0, 0, 1,
	}
}

// Screen :
func Screen(w, h int) Matrix {
	w2 := float64(w) / 2
	h2 := float64(h) / 2
	return Matrix{
		w2, 0, 0, w2,
		0, -h2, 0, h2,
		0, 0, 0.5, 0.5,
		0, 0, 0, 1,
	}
}

// Viewport :
func Viewport(x, y, w, h float64) Matrix {
	l := x
	b := y
	r := x + w
	t := y + h
	return Matrix{
		(r - l) / 2, 0, 0, (r + l) / 2,
		0, (t - b) / 2, 0, (t + b) / 2,
		0, 0, 0.5, 0.5,
		0, 0, 0, 1,
	}
}

// Translate :
func (m Matrix) Translate(v Vector) Matrix {
	return Translate(v).Mul(m)
}

// Scale :
func (m Matrix) Scale(v Vector) Matrix {
	return Scale(v).Mul(m)
}

// Rotate :
func (m Matrix) Rotate(v Vector, a float64) Matrix {
	return Rotate(v, a).Mul(m)
}

// RotateTo :
func (m Matrix) RotateTo(a, b Vector) Matrix {
	return RotateTo(a, b).Mul(m)
}

// Frustum :
func (m Matrix) Frustum(l, r, b, t, n, f float64) Matrix {
	return Frustum(l, r, b, t, n, f).Mul(m)
}

// Orthographic :
func (m Matrix) Orthographic(l, r, b, t, n, f float64) Matrix {
	return Orthographic(l, r, b, t, n, f).Mul(m)
}

// Perspective :
func (m Matrix) Perspective(fovy, aspect, near, far float64) Matrix {
	return Perspective(fovy, aspect, near, far).Mul(m)
}

// LookAt :
func (m Matrix) LookAt(eye, center, up Vector) Matrix {
	return LookAt(eye, center, up).Mul(m)
}

// Viewport :
func (m Matrix) Viewport(x, y, w, h float64) Matrix {
	return Viewport(x, y, w, h).Mul(m)
}

// MulScalar :
func (m Matrix) MulScalar(b float64) Matrix {
	return Matrix{
		m.X00 * b, m.X01 * b, m.X02 * b, m.X03 * b,
		m.X10 * b, m.X11 * b, m.X12 * b, m.X13 * b,
		m.X20 * b, m.X21 * b, m.X22 * b, m.X23 * b,
		m.X30 * b, m.X31 * b, m.X32 * b, m.X33 * b,
	}
}

// Mul :
func (m Matrix) Mul(b Matrix) Matrix {
	mm := Matrix{}
	mm.X00 = m.X00*b.X00 + m.X01*b.X10 + m.X02*b.X20 + m.X03*b.X30
	mm.X10 = m.X10*b.X00 + m.X11*b.X10 + m.X12*b.X20 + m.X13*b.X30
	mm.X20 = m.X20*b.X00 + m.X21*b.X10 + m.X22*b.X20 + m.X23*b.X30
	mm.X30 = m.X30*b.X00 + m.X31*b.X10 + m.X32*b.X20 + m.X33*b.X30
	mm.X01 = m.X00*b.X01 + m.X01*b.X11 + m.X02*b.X21 + m.X03*b.X31
	mm.X11 = m.X10*b.X01 + m.X11*b.X11 + m.X12*b.X21 + m.X13*b.X31
	mm.X21 = m.X20*b.X01 + m.X21*b.X11 + m.X22*b.X21 + m.X23*b.X31
	mm.X31 = m.X30*b.X01 + m.X31*b.X11 + m.X32*b.X21 + m.X33*b.X31
	mm.X02 = m.X00*b.X02 + m.X01*b.X12 + m.X02*b.X22 + m.X03*b.X32
	mm.X12 = m.X10*b.X02 + m.X11*b.X12 + m.X12*b.X22 + m.X13*b.X32
	mm.X22 = m.X20*b.X02 + m.X21*b.X12 + m.X22*b.X22 + m.X23*b.X32
	mm.X32 = m.X30*b.X02 + m.X31*b.X12 + m.X32*b.X22 + m.X33*b.X32
	mm.X03 = m.X00*b.X03 + m.X01*b.X13 + m.X02*b.X23 + m.X03*b.X33
	mm.X13 = m.X10*b.X03 + m.X11*b.X13 + m.X12*b.X23 + m.X13*b.X33
	mm.X23 = m.X20*b.X03 + m.X21*b.X13 + m.X22*b.X23 + m.X23*b.X33
	mm.X33 = m.X30*b.X03 + m.X31*b.X13 + m.X32*b.X23 + m.X33*b.X33
	return mm
}

// MulPosition :
func (m Matrix) MulPosition(b Vector) Vector {
	x := m.X00*b.X + m.X01*b.Y + m.X02*b.Z + m.X03
	y := m.X10*b.X + m.X11*b.Y + m.X12*b.Z + m.X13
	z := m.X20*b.X + m.X21*b.Y + m.X22*b.Z + m.X23
	return Vector{x, y, z}
}

// MulPositionW :
func (m Matrix) MulPositionW(b Vector) VectorW {
	x := m.X00*b.X + m.X01*b.Y + m.X02*b.Z + m.X03
	y := m.X10*b.X + m.X11*b.Y + m.X12*b.Z + m.X13
	z := m.X20*b.X + m.X21*b.Y + m.X22*b.Z + m.X23
	w := m.X30*b.X + m.X31*b.Y + m.X32*b.Z + m.X33
	return VectorW{x, y, z, w}
}

// MulDirection :
func (m Matrix) MulDirection(b Vector) Vector {
	x := m.X00*b.X + m.X01*b.Y + m.X02*b.Z
	y := m.X10*b.X + m.X11*b.Y + m.X12*b.Z
	z := m.X20*b.X + m.X21*b.Y + m.X22*b.Z
	return Vector{x, y, z}.Normalize()
}

// MulBox :
func (m Matrix) MulBox(box Box) Box {
	// http://dev.theomader.com/transform-bounding-boxes/
	r := Vector{m.X00, m.X10, m.X20}
	u := Vector{m.X01, m.X11, m.X21}
	b := Vector{m.X02, m.X12, m.X22}
	t := Vector{m.X03, m.X13, m.X23}
	xa := r.MulScalar(box.Min.X)
	xb := r.MulScalar(box.Max.X)
	ya := u.MulScalar(box.Min.Y)
	yb := u.MulScalar(box.Max.Y)
	za := b.MulScalar(box.Min.Z)
	zb := b.MulScalar(box.Max.Z)
	xa, xb = xa.Min(xb), xa.Max(xb)
	ya, yb = ya.Min(yb), ya.Max(yb)
	za, zb = za.Min(zb), za.Max(zb)
	min := xa.Add(ya).Add(za).Add(t)
	max := xb.Add(yb).Add(zb).Add(t)
	return Box{min, max}
}

// Transpose :
func (m Matrix) Transpose() Matrix {
	return Matrix{
		m.X00, m.X10, m.X20, m.X30,
		m.X01, m.X11, m.X21, m.X31,
		m.X02, m.X12, m.X22, m.X32,
		m.X03, m.X13, m.X23, m.X33}
}

// Determinant :
func (m Matrix) Determinant() float64 {
	return (m.X00*m.X11*m.X22*m.X33 - m.X00*m.X11*m.X23*m.X32 +
		m.X00*m.X12*m.X23*m.X31 - m.X00*m.X12*m.X21*m.X33 +
		m.X00*m.X13*m.X21*m.X32 - m.X00*m.X13*m.X22*m.X31 -
		m.X01*m.X12*m.X23*m.X30 + m.X01*m.X12*m.X20*m.X33 -
		m.X01*m.X13*m.X20*m.X32 + m.X01*m.X13*m.X22*m.X30 -
		m.X01*m.X10*m.X22*m.X33 + m.X01*m.X10*m.X23*m.X32 +
		m.X02*m.X13*m.X20*m.X31 - m.X02*m.X13*m.X21*m.X30 +
		m.X02*m.X10*m.X21*m.X33 - m.X02*m.X10*m.X23*m.X31 +
		m.X02*m.X11*m.X23*m.X30 - m.X02*m.X11*m.X20*m.X33 -
		m.X03*m.X10*m.X21*m.X32 + m.X03*m.X10*m.X22*m.X31 -
		m.X03*m.X11*m.X22*m.X30 + m.X03*m.X11*m.X20*m.X32 -
		m.X03*m.X12*m.X20*m.X31 + m.X03*m.X12*m.X21*m.X30)
}

// Inverse :
func (m Matrix) Inverse() Matrix {
	mm := Matrix{}
	d := m.Determinant()
	mm.X00 = (m.X12*m.X23*m.X31 - m.X13*m.X22*m.X31 + m.X13*m.X21*m.X32 - m.X11*m.X23*m.X32 - m.X12*m.X21*m.X33 + m.X11*m.X22*m.X33) / d
	mm.X01 = (m.X03*m.X22*m.X31 - m.X02*m.X23*m.X31 - m.X03*m.X21*m.X32 + m.X01*m.X23*m.X32 + m.X02*m.X21*m.X33 - m.X01*m.X22*m.X33) / d
	mm.X02 = (m.X02*m.X13*m.X31 - m.X03*m.X12*m.X31 + m.X03*m.X11*m.X32 - m.X01*m.X13*m.X32 - m.X02*m.X11*m.X33 + m.X01*m.X12*m.X33) / d
	mm.X03 = (m.X03*m.X12*m.X21 - m.X02*m.X13*m.X21 - m.X03*m.X11*m.X22 + m.X01*m.X13*m.X22 + m.X02*m.X11*m.X23 - m.X01*m.X12*m.X23) / d
	mm.X10 = (m.X13*m.X22*m.X30 - m.X12*m.X23*m.X30 - m.X13*m.X20*m.X32 + m.X10*m.X23*m.X32 + m.X12*m.X20*m.X33 - m.X10*m.X22*m.X33) / d
	mm.X11 = (m.X02*m.X23*m.X30 - m.X03*m.X22*m.X30 + m.X03*m.X20*m.X32 - m.X00*m.X23*m.X32 - m.X02*m.X20*m.X33 + m.X00*m.X22*m.X33) / d
	mm.X12 = (m.X03*m.X12*m.X30 - m.X02*m.X13*m.X30 - m.X03*m.X10*m.X32 + m.X00*m.X13*m.X32 + m.X02*m.X10*m.X33 - m.X00*m.X12*m.X33) / d
	mm.X13 = (m.X02*m.X13*m.X20 - m.X03*m.X12*m.X20 + m.X03*m.X10*m.X22 - m.X00*m.X13*m.X22 - m.X02*m.X10*m.X23 + m.X00*m.X12*m.X23) / d
	mm.X20 = (m.X11*m.X23*m.X30 - m.X13*m.X21*m.X30 + m.X13*m.X20*m.X31 - m.X10*m.X23*m.X31 - m.X11*m.X20*m.X33 + m.X10*m.X21*m.X33) / d
	mm.X21 = (m.X03*m.X21*m.X30 - m.X01*m.X23*m.X30 - m.X03*m.X20*m.X31 + m.X00*m.X23*m.X31 + m.X01*m.X20*m.X33 - m.X00*m.X21*m.X33) / d
	mm.X22 = (m.X01*m.X13*m.X30 - m.X03*m.X11*m.X30 + m.X03*m.X10*m.X31 - m.X00*m.X13*m.X31 - m.X01*m.X10*m.X33 + m.X00*m.X11*m.X33) / d
	mm.X23 = (m.X03*m.X11*m.X20 - m.X01*m.X13*m.X20 - m.X03*m.X10*m.X21 + m.X00*m.X13*m.X21 + m.X01*m.X10*m.X23 - m.X00*m.X11*m.X23) / d
	mm.X30 = (m.X12*m.X21*m.X30 - m.X11*m.X22*m.X30 - m.X12*m.X20*m.X31 + m.X10*m.X22*m.X31 + m.X11*m.X20*m.X32 - m.X10*m.X21*m.X32) / d
	mm.X31 = (m.X01*m.X22*m.X30 - m.X02*m.X21*m.X30 + m.X02*m.X20*m.X31 - m.X00*m.X22*m.X31 - m.X01*m.X20*m.X32 + m.X00*m.X21*m.X32) / d
	mm.X32 = (m.X02*m.X11*m.X30 - m.X01*m.X12*m.X30 - m.X02*m.X10*m.X31 + m.X00*m.X12*m.X31 + m.X01*m.X10*m.X32 - m.X00*m.X11*m.X32) / d
	mm.X33 = (m.X01*m.X12*m.X20 - m.X02*m.X11*m.X20 + m.X02*m.X10*m.X21 - m.X00*m.X12*m.X21 - m.X01*m.X10*m.X22 + m.X00*m.X11*m.X22) / d
	return mm
}
