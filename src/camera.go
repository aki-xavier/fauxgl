package fauxgl

// Camera :
type Camera struct {
	Eye       Vector
	Direction Vector
	Distance  float64
}

// CreateCamera :
func CreateCamera() *Camera {
	c := &Camera{}
	c.Distance = 5
	c.Eye = Vector{0, 0, 10}
	c.Direction = Vector{0, 0, -1}
	return c
}
