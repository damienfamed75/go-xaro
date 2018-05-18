package gameobjects

import (
	"github.com/gen2brain/raylib-go/raylib"
)

// Room is the structure to hold
// all the walls and items in a room.
type Room struct {
	Walls  []Wall
	Bounds raylib.Rectangle
}

// NewRoom will return a new full room for the player
// to explore. Including all the walls and obstacles.
func NewRoom(w []Wall) *Room {
	r := Room{}
	r.Walls = w

	mr := raylib.NewRectangle(w[0].Bounds.X, w[0].Bounds.Y, w[0].Bounds.Width, w[0].Bounds.Height)
	for _, m := range w {
		if m.Bounds.X < mr.X {
			mr.X = m.Bounds.X
		}
		if m.Bounds.Y < mr.Y {
			mr.Y = m.Bounds.Y
		}
		if m.Bounds.Width > mr.Width {
			mr.Width = m.Bounds.Width
		}
		if m.Bounds.Height > mr.Height {
			mr.Height = m.Bounds.Height
		}
	}

	r.Bounds = mr

	return &r
}

// Draw will render the entire room
func (r *Room) Draw() {
	for _, w := range r.Walls {
		w.Draw()
	}
}
