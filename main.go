package main

import (
	"github.com/damienfamed75/go-xaro/gameobjects"
	"github.com/damienfamed75/go-xaro/system"
	"github.com/gen2brain/raylib-go/raylib"
	r "github.com/gen2brain/raylib-go/raylib"
)

const (
	scrWidth  = 1024
	scrHeight = 720

	mapPath = "D:\\gocode\\src\\github.com\\damienfamed75\\go-xaro\\assets\\maps\\testmap.tmx"
	tileset = "D:\\gocode\\src\\github.com\\damienfamed75\\go-xaro\\assets\\graphics\\BasicTileSet.png"
)

func main() {
	r.SetConfigFlags(r.FlagWindowResizable)
	r.InitWindow(scrWidth, scrHeight, "Xaro")
	r.SetExitKey(raylib.KeyEscape)
	r.SetTargetFPS(144)
	system.Init() // configure the game settings and controls

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
