package main

import (
	"fmt"
	"time"

	fauxgl "github.com/aki-xavier/fauxgl/src"
	"github.com/nfnt/resize"
)

// download dragon.obj from here:
// http://graphics.cs.williams.edu/data/meshes/dragon.zip

const (
	scale  = 4    // optional supersampling
	width  = 2048 // output width in pixels
	height = 2048 // output height in pixels
	fovy   = 30   // vertical field of view in degrees
	near   = 1    // near clipping plane
	far    = 10   // far clipping plane
)

var (
	eye        = fauxgl.V(-3, 1.25, -2)               // camera position
	center     = fauxgl.V(0, -0.1, -0.1)              // view center position
	up         = fauxgl.V(0, 1, 0)                    // up vector
	light      = fauxgl.V(-0.75, 1, 0.25).Normalize() // light direction
	color      = fauxgl.HexColor("#468966")           // object color
	background = fauxgl.HexColor("#FFF8E3")           // background color
)

func timed(name string) func() {
	if len(name) > 0 {
		fmt.Printf("%s... ", name)
	}
	start := time.Now()
	return func() {
		fmt.Println(time.Since(start))
	}
}

func main() {
	var done func()

	// load a mesh
	done = timed("loading mesh")
	mesh, err := fauxgl.LoadSTL("examples/dragon/stl/dragon.stl")
	if err != nil {
		panic(err)
	}
	done()

	// fit mesh in a bi-unit cube centered at the origin
	done = timed("transforming mesh")
	mesh.BiUnitCube()
	done()

	// smooth the normals
	done = timed("smoothing normals")
	mesh.SmoothNormalsThreshold(fauxgl.Radians(30))
	done()

	// create a rendering context
	context := fauxgl.NewContext(width*scale, height*scale)
	context.ClearColorBufferWith(background)

	// create transformation matrix and light direction
	aspect := float64(width) / float64(height)
	matrix := fauxgl.LookAt(eye, center, up).Perspective(fovy, aspect, near, far)

	// render
	shader := fauxgl.NewPhongShader(matrix, light, eye)
	shader.ObjectColor = color
	context.Shader = shader
	done = timed("rendering mesh")
	context.DrawMesh(mesh)
	done()

	// downsample image for antialiasing
	done = timed("downsampling image")
	image := context.Image()
	image = resize.Resize(width, height, image, resize.Bilinear)
	done()

	// save image
	done = timed("writing output")
	fauxgl.SavePNG("out.png", image)
	done()
}
