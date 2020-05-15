package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"

	fauxgl "github.com/aki-xavier/fauxgl/src"
	"github.com/nfnt/resize"
)

const (
	scale  = 4    // optional supersampling
	width  = 1600 // output width in pixels
	height = 1600 // output height in pixels
	fovy   = 30   // vertical field of view in degrees
	near   = 1    // near clipping plane
	far    = 100  // far clipping plane
)

var (
	eye        = fauxgl.V(3*4, 3*4, 1.5*4)          // camera position
	center     = fauxgl.V(0, 0, 0)                  // view center position
	up         = fauxgl.V(0, 0, 1)                  // up vector
	light      = fauxgl.V(0.75, 0.5, 1).Normalize() // light direction
	color      = fauxgl.HexColor("#468966")         // object color
	background = fauxgl.HexColor("#FFF8E3")         // background color
)

func main() {
	mesh := fauxgl.NewEmptyMesh()
	for i := 0; i < 1500; i++ {
		var x, y, z float64
		for {
			x = rand.Float64()*2 - 1
			y = rand.Float64()*2 - 1
			z = rand.Float64()*2 - 1
			if x*x+y*y+z*z < 1 {
				break
			}
		}
		p := fauxgl.Vector{X: x, Y: y, Z: z}.MulScalar(4)
		s := fauxgl.V(0.2, 0.2, 0.2)
		u := fauxgl.RandomUnitVector()
		a := rand.Float64() * 2 * math.Pi
		c := fauxgl.NewCube()
		c.Transform(fauxgl.Orient(p, s, u, a))
		mesh.Add(c)
	}

	// create a rendering context
	context := fauxgl.NewContext(width*scale, height*scale)
	context.ClearColorBufferWith(fauxgl.Black)

	// create transformation matrix and light direction
	aspect := float64(width) / float64(height)
	matrix := fauxgl.LookAt(eye, center, up).Perspective(fovy, aspect, near, far)

	// render
	shader := fauxgl.NewPhongShader(matrix, light, eye)
	shader.ObjectColor = color
	context.Shader = shader
	start := time.Now()
	context.DrawMesh(mesh)
	fmt.Println(time.Since(start))

	mesh, _ = fauxgl.LoadSTL("examples/shapes/sphere.stl")
	mesh.SmoothNormals()
	mesh.Transform(fauxgl.Scale(fauxgl.V(2.5, 2.5, 2.5)))
	shader = fauxgl.NewPhongShader(matrix, light, eye)
	shader.ObjectColor = fauxgl.HexColor("FFFF9D").Alpha(0.65)
	shader.SpecularPower = 0
	context.Shader = shader
	context.DrawMesh(mesh)
	context.Wireframe = true
	context.DepthBias = -0.00001
	context.DrawMesh(mesh)

	// downsample image for antialiasing
	image := context.Image()
	image = resize.Resize(width, height, image, resize.Bilinear)

	// save image
	fauxgl.SavePNG("out.png", image)
}
