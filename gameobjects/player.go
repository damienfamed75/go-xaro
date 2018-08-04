package gameobjects

import (
	"github.com/damienfamed75/GoAseprite"
	"github.com/damienfamed75/go-xaro/system"
	"github.com/gen2brain/raylib-go/raylib"
)

// Player is the object of the played protagonist
type Player struct {
	// --- Public --- //
	Ase        goaseprite.AsepriteFile // The animation file and meta data
	Texture    raylib.Texture2D        // The current image on the player
	Velocity   raylib.Vector2          // Velocity applied to player every frame
	Position   raylib.Vector2          // Position (anchor top left)
	Scale      float32                 // Multiplier for how large the player is on the screen
	ShootSpeed float32                 // Shooting speed that will be upgraded
	MoveSpeed  float32                 // Moving speed that will be upgraded

	// --- Private --- //
	shot bool
}

var (
	raylibIsKeyDown = raylib.IsKeyDown
)

// NewPlayer generates a new player object
func NewPlayer() *Player {
	p := &Player{}

	p.Ase = goaseprite.New("D:\\gocode\\src\\github.com\\damienfamed75\\go-xaro\\assets\\graphics\\player.json", "player")
	p.Texture = system.GetTexture(p.Ase.ImagePath)
	p.Velocity = raylib.NewVector2(0, 0)
	p.Position = raylib.NewVector2(float32(system.HalfW())-float32((p.Ase.FrameWidth/2)), float32(system.HalfH())-float32((p.Ase.FrameHeight/2)))
	p.Scale = 1.0
	p.ShootSpeed = 0.5
	p.MoveSpeed = 120

	p.shot = false

	// Queues an animation
	p.Ase.Play("right")

	return p
}

// Update runs every frame and should be used for inputs or effects on the player
func (p *Player) Update(dt float32) (raylib.Vector2, raylib.Vector2) {
	p.Ase.Update(dt) // Run this every single frame to keep the animation going

	p.Velocity.X, p.Velocity.Y = 0, 0

	p.updateMovement()
	p.updateAction()
	p.updateAnimation()
	p.updateIdleAnimation()
	p.updateActionAnimationSpeed()

	// Apply the velocity to player's position
	difference := raylib.NewVector2(p.Velocity.X*dt, p.Velocity.Y*dt)

	p.Position.X += difference.X
	p.Position.Y += difference.Y

	return difference, raylib.Vector2{X: p.Position.X, Y: p.Position.Y}
}

// Draw runs every frame and should only be used for rendering
func (p *Player) Draw() {
	srcX, srcY := p.Ase.GetFrameXY()
	playerW, playerH := float32(+p.Ase.FrameWidth), float32(+p.Ase.FrameHeight)

	// The entire hitbox of the player
	src := raylib.NewRectangle(float32(srcX), float32(srcY),
		playerW, playerH)

	// Drawing destination of the player
	dest := raylib.NewRectangle(float32(p.Position.X), float32(p.Position.Y),
		playerW*p.Scale, playerH*p.Scale)

	// Draw the player finally
	raylib.DrawTexturePro(p.Texture, src, dest, raylib.Vector2{}, 0, raylib.White)
}
