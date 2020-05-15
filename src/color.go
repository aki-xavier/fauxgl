package fauxgl

import (
	"fmt"
	"image/color"
	"math"
	"strings"
)

var (
	// Discard :
	Discard = Color{}
	// Transparent :
	Transparent = Color{}
	// Black :
	Black = Color{0, 0, 0, 1}
	// White :
	White = Color{1, 1, 1, 1}
)

// Color :
type Color struct {
	R, G, B, A float64
}

// Gray :
func Gray(x float64) Color {
	return Color{x, x, x, 1}
}

// MakeColor :
func MakeColor(c color.Color) Color {
	r, g, b, a := c.RGBA()
	const d = 0xffff
	return Color{float64(r) / d, float64(g) / d, float64(b) / d, float64(a) / d}
}

// HexColor :
func HexColor(x string) Color {
	x = strings.Trim(x, "#")
	var r, g, b, a int
	a = 255
	switch len(x) {
	case 3:
		fmt.Sscanf(x, "%1x%1x%1x", &r, &g, &b)
		r = (r << 4) | r
		g = (g << 4) | g
		b = (b << 4) | b
	case 4:
		fmt.Sscanf(x, "%1x%1x%1x%1x", &r, &g, &b, &a)
		r = (r << 4) | r
		g = (g << 4) | g
		b = (b << 4) | b
		a = (a << 4) | a
	case 6:
		fmt.Sscanf(x, "%02x%02x%02x", &r, &g, &b)
	case 8:
		fmt.Sscanf(x, "%02x%02x%02x%02x", &r, &g, &b, &a)
	}
	const d = 0xff
	return Color{float64(r) / d, float64(g) / d, float64(b) / d, float64(a) / d}
}

// NRGBA :
func (c Color) NRGBA() color.NRGBA {
	const d = 0xff
	r := Clamp(c.R, 0, 1)
	g := Clamp(c.G, 0, 1)
	b := Clamp(c.B, 0, 1)
	a := Clamp(c.A, 0, 1)
	return color.NRGBA{uint8(r * d), uint8(g * d), uint8(b * d), uint8(a * d)}
}

// Opaque :
func (c Color) Opaque() Color {
	return Color{c.R, c.G, c.B, 1}
}

// Alpha :
func (c Color) Alpha(alpha float64) Color {
	return Color{c.R, c.G, c.B, alpha}
}

// Lerp :
func (c Color) Lerp(b Color, t float64) Color {
	return c.Add(b.Sub(c).MulScalar(t))
}

// Add :
func (c Color) Add(b Color) Color {
	return Color{c.R + b.R, c.G + b.G, c.B + b.B, c.A + b.A}
}

// Sub :
func (c Color) Sub(b Color) Color {
	return Color{c.R - b.R, c.G - b.G, c.B - b.B, c.A - b.A}
}

// Mul :
func (c Color) Mul(b Color) Color {
	return Color{c.R * b.R, c.G * b.G, c.B * b.B, c.A * b.A}
}

// Div :
func (c Color) Div(b Color) Color {
	return Color{c.R / b.R, c.G / b.G, c.B / b.B, c.A / b.A}
}

// AddScalar :
func (c Color) AddScalar(b float64) Color {
	return Color{c.R + b, c.G + b, c.B + b, c.A + b}
}

// SubScalar :
func (c Color) SubScalar(b float64) Color {
	return Color{c.R - b, c.G - b, c.B - b, c.A - b}
}

// MulScalar :
func (c Color) MulScalar(b float64) Color {
	return Color{c.R * b, c.G * b, c.B * b, c.A * b}
}

// DivScalar :
func (c Color) DivScalar(b float64) Color {
	return Color{c.R / b, c.G / b, c.B / b, c.A / b}
}

// Pow :
func (c Color) Pow(b float64) Color {
	return Color{math.Pow(c.R, b), math.Pow(c.G, b), math.Pow(c.B, b), math.Pow(c.A, b)}
}

// Min :
func (c Color) Min(b Color) Color {
	return Color{math.Min(c.R, b.R), math.Min(c.G, b.G), math.Min(c.B, b.B), math.Min(c.A, b.A)}
}

// Max :
func (c Color) Max(b Color) Color {
	return Color{math.Max(c.R, b.R), math.Max(c.G, b.G), math.Max(c.B, b.B), math.Max(c.A, b.A)}
}
