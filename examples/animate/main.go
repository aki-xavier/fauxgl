package main

import (
	"fmt"

	fauxgl "github.com/aki-xavier/fauxgl/src"
	"github.com/nfnt/resize"
)

const (
	scale  = 4   // optional supersampling
	width  = 800 // output width in pixels
	height = 800 // output height in pixels
	fovy   = 30  // vertical field of view in degrees
	near   = 1   // near clipping plane
	far    = 10  // far clipping plane
)

var (
	eye        = fauxgl.V(4, 4, 2)                  // camera position
	center     = fauxgl.V(0, 0, 0)                  // view center position
	up         = fauxgl.V(0, 0, 1)                  // up vector
	light      = fauxgl.V(0.25, 0.5, 1).Normalize() // light direction
	color      = fauxgl.HexColor("#FEB41C")         // object color
	background = fauxgl.HexColor("#24221F")         // background color
)

func main() {
	// load a mesh
	mesh, err := fauxgl.LoadSTL("examples/animate/cube.stl")
	if err != nil {
		panic(err)
	}

	// fit mesh in a bi-unit cube centered at the origin
	mesh.BiUnitCube()

	// smooth the normals
	// mesh.SmoothNormalsThreshold(Radians(30))

	// create a rendering context
	context := fauxgl.NewContext(width*scale, height*scale)

	// create transformation matrix and light direction
	aspect := float64(width) / float64(height)
	matrix := fauxgl.LookAt(eye, center, up).Perspective(fovy, aspect, near, far)

	for i := 0; i < 360; i += 20 {
		// render
		context.ClearDepthBuffer()
		context.ClearColorBufferWith(background)
		shader := fauxgl.NewPhongShader(matrix, light, eye)
		shader.ObjectColor = color
		shader.DiffuseColor = fauxgl.Gray(0.9)
		shader.SpecularColor = fauxgl.Gray(0.25)
		shader.SpecularPower = 100
		context.Shader = shader
		context.DrawMesh(mesh)

		// save image
		image := context.Image()
		image = resize.Resize(width, height, image, resize.Bilinear)
		fauxgl.SavePNG(fmt.Sprintf("out/out%03d.png", i), image)

		mesh.Transform(fauxgl.Rotate(up, fauxgl.Radians(5)))
	}
}
