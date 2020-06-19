package fauxgl

var (
	eye    = V(-1, -2, 2)
	center = V(-0.07, 0, 0)
	up     = V(0, 0, 1)
)

// Scene :
type Scene struct {
	Objects         []*Object
	Camera          *Camera
	BackgroundColor Color
	Width           int
	Height          int
	Context         *Context
}

// CreateScene :
func CreateScene(width, height int) *Scene {
	s := &Scene{}
	s.Objects = make([]*Object, 0)
	s.BackgroundColor = White
	s.Camera = CreateCamera()
	s.Width = width
	s.Height = height
	s.Context = NewContext(s.Width, s.Height)
	s.Context.ClearColor = s.BackgroundColor

	aspect := float64(width) / float64(height)
	matrix := LookAt(eye, center, up).Perspective(20, aspect, 1, 50)
	light := V(-2, 0, 1).Normalize()
	color := Color{R: 0.5, G: 1, B: 0.65, A: 1}
	shader := NewPhongShader(matrix, light, eye)
	shader.ObjectColor = color

	s.Context.Shader = shader
	return s
}

// Render :
func (s *Scene) Render() {
	s.Context.ClearColorBuffer()
	for _, obj := range s.Objects {
		s.Context.DrawMesh(obj.Mesh)
	}
}

// AddObject :
func (s *Scene) AddObject(o *Object) {
	if s.ContainsObject(o) {
		return
	}
	s.Objects = append(s.Objects, o)
}

// ContainsObject :
func (s *Scene) ContainsObject(o *Object) bool {
	for _, oo := range s.Objects {
		if oo == o {
			return true
		}
	}
	return false
}

// func (s *Scene) fillColorBufferWith(color Color) {
// 	c := color.NRGBA()
// 	for y := 0; y < s.Height; y++ {
// 		i := s.ColorBuffer.PixOffset(0, y)
// 		for x := 0; x < s.Width; x++ {
// 			s.ColorBuffer.Pix[i+0] = c.R
// 			s.ColorBuffer.Pix[i+1] = c.G
// 			s.ColorBuffer.Pix[i+2] = c.B
// 			s.ColorBuffer.Pix[i+3] = c.A
// 			i += 4
// 		}
// 	}
// }
