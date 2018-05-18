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

	mapPath = "assets/maps/testmap.tmx"
	tileset = "assets/graphics/BasicTileSet.png"
)

func main() {
	r.SetConfigFlags(r.FlagWindowResizable)
	r.InitWindow(scrWidth, scrHeight, "Xaro")
	r.SetTargetFPS(60)
	r.SetExitKey(raylib.KeyEscape)
	system.SystemInit()

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
	r.CloseWindow()
}
