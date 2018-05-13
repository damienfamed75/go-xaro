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
	player := newPlayer()

	plats := []gameobjects.Platform{
		gameobjects.NewPlatform(300, 500, 200, 10, r.Maroon),
		gameobjects.NewPlatform(100, 450, 150, 10, r.Pink),
	}
	/*********************Game Loop*********************/
	for !r.WindowShouldClose() {
		/************Update************/
		system.EscapeCheck()

		player.update(r.GetFrameTime(), plats)
		/***********Drawing***********/
		r.BeginDrawing()
		{
			r.ClearBackground(r.Black)

			player.draw()
			gameobjects.DrawPlatforms(plats...)
		}
		r.EndDrawing()
	}
	r.CloseWindow()
}
