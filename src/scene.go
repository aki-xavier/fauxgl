package fauxgl

import "image"

// Scene :
type Scene struct {
	Objects         []*Object
	Camera          *Camera
	BackgroundColor Color
	Width           int
	Height          int
	ColorBuffer     *image.NRGBA
}

// CreateScene :
func CreateScene(width, height int) *Scene {
	s := &Scene{}
	s.Objects = make([]*Object, 0)
	s.BackgroundColor = White
	s.Camera = CreateCamera()
	s.Width = width
	s.Height = height
	s.ColorBuffer = image.NewNRGBA(image.Rect(0, 0, s.Width, s.Height))
	return s
}

// Render :
func (s *Scene) Render() {
	s.fillColorBufferWith(s.BackgroundColor)
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

func (s *Scene) fillColorBufferWith(color Color) {
	c := color.NRGBA()
	for y := 0; y < s.Height; y++ {
		i := s.ColorBuffer.PixOffset(0, y)
		for x := 0; x < s.Width; x++ {
			s.ColorBuffer.Pix[i+0] = c.R
			s.ColorBuffer.Pix[i+1] = c.G
			s.ColorBuffer.Pix[i+2] = c.B
			s.ColorBuffer.Pix[i+3] = c.A
			i += 4
		}
	}
}
