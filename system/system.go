package system

import (
	"github.com/gen2brain/raylib-go/raylib"
)

var (
	KeyBindings = make(map[string]int32)
)

// EscapeCheck is used to close the window and gets called
// int the update portion of the game loop.
func EscapeCheck() {
	if raylib.IsKeyDown(raylib.KeyEscape) {
		raylib.CloseWindow()
	}
}

// HalfH returns half of the screen height
func HalfH() int32 { return raylib.GetScreenHeight() / 2 }

// HalfW returns half of the screen width
func HalfW() int32 { return raylib.GetScreenWidth() / 2 }

// SystemInit initializes the controls
func SystemInit() {
	KeyBindings["left"] = raylib.KeyLeft
	KeyBindings["right"] = raylib.KeyRight
	KeyBindings["up"] = raylib.KeyUp
	KeyBindings["action"] = raylib.KeyX
	KeyBindings["menu"] = raylib.KeyEnter
}
