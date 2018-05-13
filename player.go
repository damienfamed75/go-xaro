package main

import (
	"github.com/damienfamed75/go-xaro/gameobjects"
	"github.com/damienfamed75/go-xaro/system"
	"github.com/gen2brain/raylib-go/raylib"
	"github.com/solarlune/GoAseprite"
)

// Player is the object of the played protagonist
type Player struct {
	Ase         goaseprite.AsepriteFile // The animation file and meta data
	Texture     raylib.Texture2D        // The current image on the player
	Velocity    raylib.Vector2          // Velocity applied to player every frame
	Position    raylib.Vector2          // Position (anchor top left)
	Scale       float32                 // Multiplier for how large the player is on the screen
	facingRight bool                    // Whether the player is facing right or left
	grounded    bool                    // If the player is on the ground
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
	p.facingRight = true
	p.grounded = false
	p.Scale = 1.5

	// Queues the Run animation
	p.Ase.Play("Run")

	return p
}

// Update runs every frame and should be used for inputs or effects on the player
func (p *Player) update(dt float32, plats []gameobjects.Platform) {
	p.Ase.Update(dt) // Run this every single frame to keep the animation going

	p.Velocity.X = 0
	p.Velocity.Y += gravity * dt // Apply gravity by default

	// Movement bindings for player movement
	if raylib.IsKeyDown(system.KeyBindings["left"]) {
		p.Velocity.X = -moveSpd
		p.facingRight = false
	}
	if raylib.IsKeyDown(system.KeyBindings["right"]) {
		p.Velocity.X = +moveSpd
		p.facingRight = true
	}

	// Finding if the player is touching the ground or not
	p.grounded = false
	if p.Velocity.Y >= 0 {
		for _, pl := range plats {
			if p.Position.X+float32(p.Ase.FrameWidth) <= pl.Bounds.X || p.Position.X >= pl.Bounds.X+float32(pl.Bounds.Width) {
				continue
			}
			if p.Position.Y+float32(p.Ase.FrameHeight) <= pl.Bounds.Y-(pl.Bounds.Height*2) || p.Position.Y > pl.Bounds.Y+pl.Bounds.Height {
				continue
			}
			if p.Position.Y+float32(p.Ase.FrameHeight) > pl.Bounds.Y-(pl.Bounds.Height*1.9) {
				p.Position.Y-- // push player up if they're sunken into the platform
			}
			p.Velocity.Y = 0
			p.grounded = true
		}
	}

	// If jumping and on the ground
	if p.grounded && raylib.IsKeyDown(system.KeyBindings["up"]) {
		p.Velocity.Y = jumpFrc
	}

	// Apply the velocity to player's position
	p.Position.X += p.Velocity.X * dt
	p.Position.Y += p.Velocity.Y * dt
}

// Draw runs every frame and should only be used for rendering
func (p *Player) draw() {
	srcX, srcY := p.Ase.GetFrameXY()

	var playerWidth float32
	if p.facingRight == true {
		playerWidth = float32(+p.Ase.FrameWidth)
	} else {
		playerWidth = float32(-p.Ase.FrameWidth)
	} // flips the character around whether or not the player is facing right

	// The entire hitbox of the player
	src := raylib.NewRectangle(srcX, srcY,
		playerWidth, float32(p.Ase.FrameHeight))

	// Drawing destination of the player
	dest := raylib.NewRectangle(float32(p.Position.X), float32(p.Position.Y),
		float32(p.Ase.FrameWidth)*p.Scale, float32(p.Ase.FrameHeight)*p.Scale)

	// Draw the player finally
	raylib.DrawTexturePro(p.Texture, src, dest, raylib.Vector2{}, 0, raylib.White)
}
