package gameobjects

import "github.com/gen2brain/raylib-go/raylib"

// Platform is a debug area for the player to stand on
type Wall struct {
	Bounds raylib.Rectangle
	Color  raylib.Color
}

// NewPlatform returns a new platform object
func NewWall(x float32, y float32, w float32, h float32, c raylib.Color) Wall {
	wall := Wall{}

	wall.Bounds = raylib.NewRectangle(x, y, w, h)
	wall.Color = c

	return wall
}

// Draw the platform individually
func (w *Wall) Draw() {
	raylib.DrawRectanglePro(w.Bounds, raylib.Vector2{}, 0, w.Color)
}

// DrawPlatforms draws an array of platforms
func DrawPlatforms(walls ...Wall) {
	for _, w := range walls {
		w.Draw()
	}
}
