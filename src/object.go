package fauxgl

// Object :
type Object struct {
	Mesh   *Mesh
	Color  Color
	Matrix Matrix
}

// CreateObject :
func CreateObject(mesh *Mesh, color Color) *Object {
	o := &Object{}
	o.Mesh = mesh
	o.Color = color
	o.Matrix = Identity()
	return o
}
