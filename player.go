package main

import (
	"github.com/damienfamed75/go-xaro/gameobjects"
	"github.com/damienfamed75/go-xaro/system"
	"github.com/gen2brain/raylib-go/raylib"
	"github.com/solarlune/GoAseprite"
)

// Player is the object of the played protagonist
type Player struct {
	Ase       goaseprite.AsepriteFile // The animation file and meta data
	Texture   raylib.Texture2D        // The current image on the player
	Velocity  raylib.Vector2          // Velocity applied to player every frame
	Position  raylib.Vector2          // Position (anchor top left)
	Scale     float32                 // Multiplier for how large the player is on the screen
	direction string                  // Whether the player is facing right or left
}

const gravity = 400

var (
	moveSpd float32 = 120
	jumpFrc float32 = -250
)

// NewPlayer generates a new player object
func newPlayer() Player {
	p := Player{}

	p.Ase = goaseprite.New("assets/graphics/samus.json")
	p.Texture = system.GetTexture(p.Ase.ImagePath)
	p.Velocity = raylib.NewVector2(0, 0)
	p.Position = raylib.NewVector2(float32(system.HalfW()), float32(system.HalfH()-p.Ase.FrameHeight))
	p.direction = "right"
	p.Scale = 1.5

	// Queues the Run animation
	p.Ase.Play("Run")

	return p
}

// Update runs every frame and should be used for inputs or effects on the player
func (p *Player) update(dt float32, walls []gameobjects.Wall) {
	p.Ase.Update(dt) // Run this every single frame to keep the animation going

	p.Velocity.X = 0
	p.Velocity.Y = 0

	// Movement bindings for player movement
	if raylib.IsKeyDown(system.KeyBindings["left"]) {
		p.Velocity.X = -moveSpd
		p.direction = "left"
	}
	if raylib.IsKeyDown(system.KeyBindings["right"]) {
		p.Velocity.X = +moveSpd
		p.direction = "right"
	}
	if raylib.IsKeyDown(system.KeyBindings["up"]) {
		p.Velocity.Y = -moveSpd
		p.direction = "up"
	}
	if raylib.IsKeyDown(system.KeyBindings["down"]) {
		p.Velocity.Y = +moveSpd
		p.direction = "down"
	}

	// Apply the velocity to player's position
	p.Position.X += p.Velocity.X * dt
	p.Position.Y += p.Velocity.Y * dt
}

// Draw runs every frame and should only be used for rendering
func (p *Player) draw() {
	srcX, srcY := p.Ase.GetFrameXY()

	playerWidth := float32(+p.Ase.FrameWidth)   // default value is upright
	playerHeight := float32(+p.Ase.FrameHeight) // default value is facing right

	if p.direction == "left" {
		playerWidth = float32(-p.Ase.FrameWidth)
	} else if p.direction == "down" {
		playerHeight = float32(-p.Ase.FrameHeight)
	}

	// The entire hitbox of the player
	src := raylib.NewRectangle(srcX, srcY,
		playerWidth, playerHeight)

	// Drawing destination of the player
	dest := raylib.NewRectangle(float32(p.Position.X), float32(p.Position.Y),
		float32(p.Ase.FrameWidth)*p.Scale, float32(p.Ase.FrameHeight)*p.Scale)

	// Draw the player finally
	raylib.DrawTexturePro(p.Texture, src, dest, raylib.Vector2{}, 0, raylib.White)
}
