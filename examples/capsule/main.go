package main

import (
	"fmt"
	"time"

	fauxgl "github.com/aki-xavier/fauxgl/src"
	"github.com/nfnt/resize"
)

const (
	scale  = 4    // optional supersampling
	width  = 1024 // output width in pixels
	height = 1024 // output height in pixels
	fovy   = 40   // vertical field of view in degrees
	near   = 1    // near clipping plane
	far    = 10   // far clipping plane
)

var (
	eye    = fauxgl.V(-3, 1, 2)                // camera position
	center = fauxgl.V(0, 0, 0)                 // view center position
	up     = fauxgl.V(0, 0, 1)                 // up vector
	light  = fauxgl.V(-1, 1, 0.25).Normalize() // light direction
)

func main() {
	// load the mesh
	mesh, err := fauxgl.LoadOBJ("examples/capsule/capsule.obj")
	if err != nil {
		panic(err)
	}

	// load the texture
	texture, err := fauxgl.LoadTexture("examples/capsule/capsule.jpg")
	if err != nil {
		panic(err)
	}

	// fit mesh in a bi-unit cube centered at the origin
	mesh.BiUnitCube()

	// create a rendering context
	context := fauxgl.NewContext(width*scale, height*scale)

	// create transformation matrix and light direction
	aspect := float64(width) / float64(height)
	matrix := fauxgl.LookAt(eye, center, up).Perspective(fovy, aspect, near, far)

	// render
	shader := fauxgl.NewPhongShader(matrix, light, eye)
	shader.Texture = texture
	context.Shader = shader
	start := time.Now()
	context.DrawMesh(mesh)
	fmt.Println(time.Since(start))

	// downsample image for antialiasing
	image := context.Image()
	image = resize.Resize(width, height, image, resize.Bilinear)

	// save image
	fauxgl.SavePNG("out.png", image)
}
