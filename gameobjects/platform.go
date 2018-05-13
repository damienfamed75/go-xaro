package gameobjects

import "github.com/gen2brain/raylib-go/raylib"

// Platform is a debug area for the player to stand on
type Platform struct {
	Bounds raylib.Rectangle
	Color  raylib.Color
}

// NewPlatform returns a new platform object
func NewPlatform(x float32, y float32, w float32, h float32, c raylib.Color) Platform {
	plat := Platform{}

	plat.Bounds = raylib.NewRectangle(x, y, w, h)
	plat.Color = c

	return plat
}

// Draw the platform individually
func (p *Platform) Draw() {
	raylib.DrawRectanglePro(p.Bounds, raylib.Vector2{}, 0, p.Color)
}

// DrawPlatforms draws an array of platforms
func DrawPlatforms(plats ...Platform) {
	for _, p := range plats {
		p.Draw()
	}
}
