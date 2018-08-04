package system

import (
	"github.com/gen2brain/raylib-go/raylib"
)

var (
	KeyBindings = make(map[string]int32)
)

// EscapeCheck is used to close the window and gets called
// int the update portion of the game loop.

// HalfH returns half of the screen height
func HalfH() int32 { return raylib.GetScreenHeight() / 2 }

// HalfW returns half of the screen width
func HalfW() int32 { return raylib.GetScreenWidth() / 2 }

// Init initializes the controls
func Init() {
	KeyBindings["left"] = raylib.KeyA
	KeyBindings["right"] = raylib.KeyD
	KeyBindings["up"] = raylib.KeyW
	KeyBindings["down"] = raylib.KeyS
	KeyBindings["action"] = raylib.MouseLeftButton
	KeyBindings["menu"] = raylib.KeyTab
}
