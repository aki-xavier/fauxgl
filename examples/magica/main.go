package main

import (
	"fmt"
	"os"
	"time"

	fauxgl "github.com/aki-xavier/fauxgl/src"
	"github.com/nfnt/resize"
)

const (
	scale  = 4    // optional supersampling
	width  = 1600 // output width in pixels
	height = 1600 // output height in pixels
	fovy   = 10   // vertical field of view in degrees
	near   = 1    // near clipping plane
	far    = 100  // far clipping plane
)

var (
	eye    = fauxgl.V(-10, -10, 10)                // camera position
	center = fauxgl.V(0, 0, -0.25)                 // view center position
	up     = fauxgl.V(0, 0, 1)                     // up vector
	light  = fauxgl.V(-0.25, -0.75, 1).Normalize() // light direction
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
	total := timed("")

	var done func()

	// load a mesh
	done = timed("loading vox file")
	voxels, err := fauxgl.LoadVOX(os.Args[1])
	if err != nil {
		panic(err)
	}
	done()

	done = timed("generating mesh")
	mesh := fauxgl.NewVoxelMesh(voxels)
	done()

	fmt.Println(len(voxels), "voxels")
	fmt.Println(len(mesh.Triangles), "triangles")
	fmt.Println(len(mesh.Lines), "lines")

	// fit mesh in a bi-unit cube centered at the origin
	mesh.BiUnitCube()

	// create a rendering context
	context := fauxgl.NewContext(width*scale, height*scale)
	context.ClearColorBufferWith(fauxgl.HexColor("323"))

	// create transformation matrix and light direction
	const s = 1.5
	matrix := fauxgl.LookAt(eye, center, up).Orthographic(-s, s, -s, s, -20, 20)

	// render
	shader := fauxgl.NewPhongShader(matrix, light, eye)
	shader.AmbientColor = fauxgl.Gray(0.4)
	shader.DiffuseColor = fauxgl.Gray(0.9)
	shader.SpecularPower = 0
	context.Shader = shader
	done = timed("rendering triangles")
	context.DrawTriangles(mesh.Triangles)
	done()

	context.Shader = fauxgl.NewSolidColorShader(matrix, fauxgl.HexColor("000"))
	context.Wireframe = true
	context.LineWidth = scale * 2
	context.DepthBias = -4e-5
	done = timed("rendering lines")
	context.DrawLines(mesh.Lines)
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

	total()
}
