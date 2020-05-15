package main

import fauxgl "github.com/aki-xavier/fauxgl/src"

const (
	width  = 2000
	height = 1000
	fovy   = 20
	near   = 1
	far    = 50
)

var (
	eye    = fauxgl.V(-1, -2, 2)
	center = fauxgl.V(-0.07, 0, 0)
	up     = fauxgl.V(0, 0, 1)
)

func main() {
	mesh, err := fauxgl.LoadSTL("examples/hello/hello.stl")
	if err != nil {
		panic(err)
	}
	mesh.BiUnitCube()
	mesh.SmoothNormalsThreshold(fauxgl.Radians(30))

	context := fauxgl.NewContext(width, height)
	context.ClearColor = fauxgl.Black
	context.ClearColorBuffer()

	aspect := float64(width) / float64(height)
	matrix := fauxgl.LookAt(eye, center, up).Perspective(fovy, aspect, near, far)
	light := fauxgl.V(-2, 0, 1).Normalize()
	color := fauxgl.Color{R: 0.5, G: 1, B: 0.65, A: 1}

	shader := fauxgl.NewPhongShader(matrix, light, eye)
	shader.ObjectColor = color
	context.Shader = shader
	context.DrawMesh(mesh)

	fauxgl.SavePNG("out.png", context.Image())
}
