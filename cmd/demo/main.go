package main

import fauxgl "github.com/aki-xavier/fauxgl/src"

func main() {
	scene := fauxgl.CreateScene(1024, 768)
	cube := fauxgl.NewCube()
	obj := fauxgl.CreateObject(cube, fauxgl.Black)
	scene.AddObject(obj)

	scene.Render()
	fauxgl.SavePNG("out.png", scene.ColorBuffer)
}
