package main

import (
	"github.com/damienfamed75/go-xaro/gameobjects"
	"github.com/damienfamed75/go-xaro/system"
	r "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	system.Init()

	r.SetConfigFlags(system.Config.Window.Mode)
	r.InitWindow(system.Config.Window.Width, system.Config.Window.Height, "Xaro")
	r.SetTargetFPS(system.Config.Window.TargetFPS)
	r.SetExitKey(r.KeyEscape)

	defer r.CloseWindow()
	/*********************Variables*********************/
	player := gameobjects.NewPlayer()
	cam := gameobjects.NewCamera()
	tileMap := gameobjects.NewTileMap(system.MapPath)

	/*********************Game Loop*********************/
	for !r.WindowShouldClose() {
		/************Update************/
		cam.Update(player.Update(r.GetFrameTime()))

		/***********Drawing***********/
		r.BeginDrawing()
		r.BeginMode2D(r.Camera2D(*cam))
		{
			r.ClearBackground(r.Black)

			tileMap.Draw(system.TileSet)
			player.Draw()
		}
		r.EndMode2D()
		r.EndDrawing()
	}
}
