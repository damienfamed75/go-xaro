package main

import (
	"github.com/damienfamed75/go-xaro/gameobjects"
	"github.com/damienfamed75/go-xaro/system"
	r "github.com/gen2brain/raylib-go/raylib"
)

const (
	scrWidth  = 640
	scrHeight = 640
)

func main() {
	r.SetConfigFlags(r.FlagWindowResizable)
	r.InitWindow(scrWidth, scrHeight, "Xaro")
	r.SetTargetFPS(60)

	/*********************Configure*********************/
	system.SystemInit()
	player := gameobjects.NewPlayer()
	cam := gameobjects.NewCamera()

	walls := []gameobjects.Wall{
		gameobjects.NewWall(0, 600, 640, 40, r.Maroon),
		gameobjects.NewWall(0, 0, 40, 640, r.Pink),
		gameobjects.NewWall(0, 0, 640, 40, r.Orange),
		gameobjects.NewWall(600, 0, 40, 640, r.Gold),
		gameobjects.NewWall(320, 0, 1, 120, r.Yellow),
	}

	/*********************Game Loop*********************/
	for !r.WindowShouldClose() {
		/************Update************/
		system.EscapeCheck()
		cam.Update(player.Update(r.GetFrameTime(), walls))
		/***********Drawing***********/
		r.BeginDrawing()
		r.BeginMode2D(r.Camera2D(*cam))
		{
			r.ClearBackground(r.Black)

			player.Draw()
			gameobjects.DrawPlatforms(walls...)
		}
		r.EndMode2D()
		r.EndDrawing()
	}
	r.CloseWindow()
}
