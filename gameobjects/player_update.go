package gameobjects

import (
	"strings"
	"time"

	"github.com/damienfamed75/go-xaro/system"
	"github.com/gen2brain/raylib-go/raylib"
)

var (
	directions = []string{"left", "right", "up", "down"}
	timeStamp  = time.Now().UnixNano() / int64(time.Millisecond)
	timer      time.Timer
)

func (p *Player) updateMovement() {
	if raylibIsKeyDown(system.KeyBindings["left"]) {
		p.Velocity.X = -p.MoveSpeed
	}
	if raylibIsKeyDown(system.KeyBindings["right"]) {
		p.Velocity.X = +p.MoveSpeed
	}
	if raylibIsKeyDown(system.KeyBindings["up"]) {
		p.Velocity.Y = -p.MoveSpeed
	}
	if raylibIsKeyDown(system.KeyBindings["down"]) {
		p.Velocity.Y = +p.MoveSpeed
	}
}

func (p *Player) inAction() bool {
	if raylib.IsMouseButtonDown(system.KeyBindings["action"]) {
		p.Ase.PlaySpeed = p.ShootSpeed
		animName := p.Ase.CurrentAnimation.Name

		if strings.HasSuffix(animName, "idle") {
			p.Ase.Play(strings.TrimSuffix(animName, "idle") + "action")
		} else if !strings.HasSuffix(animName, "action") {
			p.Ase.Play(animName + "action")
		}

		if timeStamp += int64(raylib.GetFrameTime()); timeStamp <= time.Now().UnixNano()/int64(time.Millisecond) {
			timeStamp = time.Now().UnixNano()/int64(time.Millisecond) + int64((float32(p.Ase.CurrentAnimation.End-(p.Ase.CurrentAnimation.Start-1))*100)/p.ShootSpeed)
			p.Action()
		}
		return true
	}
	return false
}

func (p *Player) updateAnimation() {
	if !p.inAction() {
		p.Ase.PlaySpeed = 1.0
		for _, dir := range directions {
			if raylibIsKeyDown(system.KeyBindings[dir]) {
				p.Ase.Play(dir)
			}
		}
	}
}

func (p *Player) updateIdleAnimation() {
	if p.Velocity.X == 0 && p.Velocity.Y == 0 && !strings.HasSuffix(p.Ase.CurrentAnimation.Name, "idle") && !raylib.IsMouseButtonDown(system.KeyBindings["action"]) {
		if strings.HasSuffix(p.Ase.CurrentAnimation.Name, "action") {
			fix := strings.TrimSuffix(p.Ase.CurrentAnimation.Name, "action")
			p.Ase.Play(fix + "idle")
		} else {
			p.Ase.Play(p.Ase.CurrentAnimation.Name + "idle")
		}
	}
}
