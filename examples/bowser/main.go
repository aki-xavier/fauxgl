package main

import (
	"fmt"
	"time"

	fauxgl "github.com/aki-xavier/fauxgl/src"
	"github.com/nfnt/resize"
)

const (
	scale  = 4
	width  = 1024
	height = 1024
	fovy   = 30
	near   = 1
	far    = 10
)

var (
	eye    = fauxgl.V(3, 1, 0.5)
	center = fauxgl.V(0, -0.1, 0)
	up     = fauxgl.V(0, 0, 1)
)

func main() {
	// load a mesh
	mesh, err := fauxgl.LoadSTL("examples/bowser/bowser.stl")
	if err != nil {
		panic(err)
	}

	// fit mesh in a bi-unit cube centered at the origin
	mesh.BiUnitCube()

	// smooth the normals
	mesh.SmoothNormalsThreshold(fauxgl.Radians(30))

	// create a rendering context
	context := fauxgl.NewContext(width*scale, height*scale)
	context.ClearColorBufferWith(fauxgl.Black)

	// create transformation matrix and light direction
	aspect := float64(width) / float64(height)
	matrix := fauxgl.LookAt(eye, center, up).Perspective(fovy, aspect, near, far)
	light := fauxgl.V(0.75, 0.25, 1).Normalize()

	// render
	shader := fauxgl.NewPhongShader(matrix, light, eye)
	shader.ObjectColor = fauxgl.HexColor("FFD34E")
	shader.DiffuseColor = fauxgl.Gray(0.9)
	shader.SpecularColor = fauxgl.Gray(0.25)
	shader.SpecularPower = 100
	context.Shader = shader
	start := time.Now()
	info := context.DrawMesh(mesh)
	fmt.Println(info)
	fmt.Println(time.Since(start))

	// save image
	image := context.Image()
	image = resize.Resize(width, height, image, resize.Bilinear)
	fauxgl.SavePNG("out.png", image)
}
