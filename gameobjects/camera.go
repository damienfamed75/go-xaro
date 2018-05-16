package gameobjects

import (
	"github.com/gen2brain/raylib-go/raylib"
)

type XaroCamera raylib.Camera2D

var pBeforePos raylib.Vector2
var pAfterPos raylib.Vector2

// NewCamera creates and puts in default
// values for a new main camera object.
func NewCamera() *XaroCamera {
	cam := XaroCamera{}

	cam.Offset = raylib.Vector2{X: 0, Y: 0}
	cam.Rotation = 0
	cam.Zoom = 1

	return &cam
}

// Update for XaroCamera updates the offset and
// target of the camera according to the player.
func (c *XaroCamera) Update(diff raylib.Vector2, curr raylib.Vector2) {
	c.Offset.X -= diff.X
	c.Offset.Y -= diff.Y

	c.Target = raylib.Vector2{X: curr.X, Y: curr.Y}
}
