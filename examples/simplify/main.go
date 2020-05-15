package main

import (
	"fmt"
	"image"

	fauxgl "github.com/aki-xavier/fauxgl/src"
	"github.com/nfnt/resize"
)

const (
	scale  = 4
	width  = 2048
	height = 2048
	fovy   = 30
	near   = 1
	far    = 10
)

var (
	eye    = fauxgl.V(-2, -4, 2)
	center = fauxgl.V(-0.1, 0, -0.1)
	up     = fauxgl.V(0, 0, 1)
)

func render(mesh *fauxgl.Mesh) image.Image {
	context := fauxgl.NewContext(width*scale, height*scale)
	context.ClearColorBufferWith(fauxgl.White)

	aspect := float64(width) / float64(height)
	matrix := fauxgl.LookAt(eye, center, up).Perspective(fovy, aspect, near, far)
	light := fauxgl.V(-0.75, -0.25, 1).Normalize()

	shader := fauxgl.NewPhongShader(matrix, light, eye)
	shader.ObjectColor = fauxgl.HexColor("FFD34E")
	shader.DiffuseColor = fauxgl.Gray(0.9)
	shader.SpecularColor = fauxgl.Gray(0.25)
	shader.SpecularPower = 100
	context.Shader = shader
	context.DrawMesh(mesh)

	context.Shader = fauxgl.NewSolidColorShader(matrix, fauxgl.Black)
	context.DepthBias = -1e-4
	context.Wireframe = true
	context.LineWidth = 4
	context.DrawMesh(mesh)

	image := context.Image()
	image = resize.Resize(width, height, image, resize.Bilinear)
	return image
}

func main() {
	mesh, err := fauxgl.LoadSTL("examples/simplify/bunny.stl")
	if err != nil {
		panic(err)
	}
	mesh.Transform(fauxgl.Matrix{X00: 0.023175793856519147, X01: 0, X02: 0, X03: 0, X10: 0, X11: 0.023175793856519147, X12: 0, X13: 0, X20: 0, X21: 0, X22: 0.023175793856519147, X23: -0.9704647076255632, X30: 0, X31: 0, X32: 0, X33: 1})
	fmt.Println(len(mesh.Triangles))
	f := 1.0
	for i := 1; ; i++ {
		m := mesh.Copy()
		m.Simplify(f)
		fmt.Println(i, f, len(m.Triangles))
		image := render(m)
		fauxgl.SavePNG(fmt.Sprintf("bunny/out%02d.png", i), image)
		f *= 0.75
	}
}
