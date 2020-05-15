package main

import (
	"fmt"
	"math/rand"

	fauxgl "github.com/aki-xavier/fauxgl/src"
	"github.com/nfnt/resize"
)

const (
	scale  = 4
	width  = 1920
	height = 1080
	fovy   = 40
	near   = 1
	far    = 100
)

var (
	eye    = fauxgl.V(0, 0, 4)
	center = fauxgl.V(0, 0, 0)
	up     = fauxgl.V(0, 1, 0)
	light  = fauxgl.V(0.5, 0.5, 2).Normalize()
)

// RandomColor :
func RandomColor() fauxgl.Color {
	r := rand.Float64()
	g := rand.Float64()
	b := rand.Float64()
	return fauxgl.Color{R: r, G: g, B: b, A: 1}
}

// Directions :
var Directions = []Cell{
	{-1, 0, 0},
	{1, 0, 0},
	{0, -1, 0},
	{0, 1, 0},
	{0, 0, -1},
	{0, 0, 1},
}

// Cell :
type Cell struct {
	X, Y, Z int
}

// Add :
func (c Cell) Add(d Cell) Cell {
	return Cell{c.X + d.X, c.Y + d.Y, c.Z + d.Z}
}

// Sub :
func (c Cell) Sub(d Cell) Cell {
	return Cell{c.X - d.X, c.Y - d.Y, c.Z - d.Z}
}

// Vector :
func (c Cell) Vector() fauxgl.Vector {
	return fauxgl.V(float64(c.X), float64(c.Y), float64(c.Z))
}

// Mesh :
func (c Cell) Mesh() *fauxgl.Mesh {
	const s = 0.125
	mesh := fauxgl.NewLatLngSphere(15, 15)
	mesh.Transform(fauxgl.Scale(fauxgl.V(s, s, s)).Translate(c.Vector()))
	return mesh
}

// Grid :
type Grid struct {
	W, H, D int
	Cells   map[Cell]bool
}

// NewGrid :
func NewGrid(w, h, d int) *Grid {
	cells := make(map[Cell]bool)
	return &Grid{w, h, d, cells}
}

// Load :
func (g *Grid) Load() float64 {
	capacity := g.W * g.H * g.D
	size := len(g.Cells)
	return float64(size) / float64(capacity)
}

// RandomEmptyCell :
func (g *Grid) RandomEmptyCell() Cell {
	for {
		c := Cell{rand.Intn(g.W), rand.Intn(g.H), rand.Intn(g.D)}
		if !g.Get(c) {
			g.Set(c)
			return c
		}
	}
}

// Get :
func (g *Grid) Get(c Cell) bool {
	if c.X < 0 || c.Y < 0 || c.Z < 0 {
		return true
	}
	if c.X >= g.W || c.Y >= g.H || c.Z >= g.D {
		return true
	}
	return g.Cells[c]
}

// Set :
func (g *Grid) Set(c Cell) {
	g.Cells[c] = true
}

// MakeSegment :
func MakeSegment(p0, p1 fauxgl.Vector, r float64, c fauxgl.Color) *fauxgl.Mesh {
	p := p0.Add(p1).MulScalar(0.5)
	h := p0.Distance(p1)
	up := p1.Sub(p0).Normalize()
	mesh := fauxgl.NewCylinder(15, false)
	mesh.Transform(fauxgl.Orient(p, fauxgl.V(r, r, h), up, 0))
	return mesh
}

// Pipe :
type Pipe struct {
	Cell      Cell
	Direction Cell
	Color     fauxgl.Color
	Done      bool
	Mesh      *fauxgl.Mesh
}

// NewPipe :
func NewPipe(cell Cell) *Pipe {
	direction := Cell{}
	color := RandomColor()
	return &Pipe{cell, direction, color, false, fauxgl.NewEmptyMesh()}
}

// Update :
func (pipe *Pipe) Update(grid *Grid) {
	if pipe.Done {
		return
	}
	cells := make([]Cell, 0, 6)
	for _, d := range Directions {
		c := pipe.Cell.Add(d)
		if !grid.Get(c) {
			cells = append(cells, c)
		}
	}
	if len(cells) == 0 {
		pipe.Done = true
		return
	}
	c := cells[rand.Intn(len(cells))]
	d := c.Sub(pipe.Cell)
	if d != pipe.Direction {
		pipe.Mesh.Add(pipe.Cell.Mesh())
	}
	p0 := pipe.Cell.Vector()
	pipe.Cell = c
	p1 := pipe.Cell.Vector()
	pipe.Mesh.Add(MakeSegment(p0, p1, 0.125, pipe.Color))
	grid.Set(pipe.Cell)
	pipe.Direction = d
}

// GetMesh :
func (pipe *Pipe) GetMesh() *fauxgl.Mesh {
	mesh := pipe.Mesh.Copy()
	mesh.Add(pipe.Cell.Mesh())
	for _, t := range mesh.Triangles {
		t.V1.Color = pipe.Color
		t.V2.Color = pipe.Color
		t.V3.Color = pipe.Color
	}
	return mesh
}

func main() {
	aspect := float64(width) / float64(height)
	matrix := fauxgl.LookAt(eye, center, up).Perspective(fovy, aspect, near, far)

	context := fauxgl.NewContext(width*scale, height*scale)
	context.ClearColor = fauxgl.Black
	context.Shader = fauxgl.NewPhongShader(matrix, light, eye)

	grid := NewGrid(19, 11, 11)
	pipes := make([]*Pipe, 8)
	for i := range pipes {
		pipes[i] = NewPipe(grid.RandomEmptyCell())
	}

	for grid.Load() < 0.9 {
		// mesh := NewEmptyMesh()
		dead := 0
		for _, pipe := range pipes {
			if pipe.Done {
				// mesh.Add(pipe.GetMesh())
				continue
			}
			pipe.Update(grid)
			// mesh.Add(pipe.GetMesh())
			if pipe.Done {
				dead++
			}
		}
		for j := 0; j < dead; j++ {
			pipes = append(pipes, NewPipe(grid.RandomEmptyCell()))
		}
		// mesh.Transform(Translate(V(-5, -5, -5)).Scale(V(0.2, 0.2, 0.2)))
		// mesh.SmoothNormals()

		// fmt.Println(i, len(pipes), len(mesh.Triangles))

		// context.ClearColorBuffer()
		// context.ClearDepthBuffer()
		// context.DrawMesh(mesh)

		// image := context.Image()
		// image = resize.Resize(width, height, image, resize.Bilinear)

		// SavePNG(fmt.Sprintf("frame%06d.png", i), image)
	}

	mesh := fauxgl.NewEmptyMesh()
	for _, pipe := range pipes {
		mesh.Add(pipe.GetMesh())
	}
	mesh.Transform(fauxgl.Translate(fauxgl.V(-9, -5, -5)).Scale(fauxgl.V(0.2, 0.2, 0.2)))
	mesh.SmoothNormals()

	fmt.Println(len(pipes), len(mesh.Triangles))

	context.ClearColorBuffer()
	context.ClearDepthBuffer()
	context.DrawMesh(mesh)

	image := context.Image()
	image = resize.Resize(width, height, image, resize.Bilinear)

	fauxgl.SavePNG("out.png", image)
}
