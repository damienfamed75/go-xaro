package main

import (
	"github.com/damienfamed75/go-xaro/gameobjects"
	"github.com/damienfamed75/go-xaro/system"
	r "github.com/gen2brain/raylib-go/raylib"
)

const (
	mapPath = "D:\\gocode\\src\\github.com\\damienfamed75\\go-xaro\\assets\\maps\\testmap.tmx"
	tileset = "D:\\gocode\\src\\github.com\\damienfamed75\\go-xaro\\assets\\graphics\\BasicTileSet.png"
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
	tileMap := gameobjects.NewTileMap(mapPath)

	/*********************Game Loop*********************/
	for !r.WindowShouldClose() {
		/************Update************/
		cam.Update(player.Update(r.GetFrameTime()))

		/***********Drawing***********/
		r.BeginDrawing()
		r.BeginMode2D(r.Camera2D(*cam))
		{
			r.ClearBackground(r.Black)

			tileMap.Draw(tileset)
			player.Draw()
		}
		r.EndMode2D()
		r.EndDrawing()
	}
}
