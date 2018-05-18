package main

import (
	"fmt"
	"os"

	"github.com/damienfamed75/go-xaro/gameobjects"
	"github.com/damienfamed75/go-xaro/system"
	r "github.com/gen2brain/raylib-go/raylib"
	"github.com/lafriks/go-tiled"
)

const (
	scrWidth  = 1024
	scrHeight = 720

	mapPath = "assets\\maps\\testmap.tmx"
)

func main() {
	r.SetConfigFlags(r.FlagWindowResizable)
	r.InitWindow(scrWidth, scrHeight, "Xaro")
	r.SetTargetFPS(60)

	/*********************Configure*********************/
	system.SystemInit()
	player := gameobjects.NewPlayer()
	cam := gameobjects.NewCamera()

	tileMap, err := tiled.LoadFromFile(mapPath)
	if err != nil {
		fmt.Println("Couldn't parse map")
		os.Exit(2)
	}

	room := gameobjects.NewRoom([]gameobjects.Wall{
		gameobjects.NewWall(0, 600, 640, 40, r.Maroon),
		gameobjects.NewWall(0, 0, 40, 640, r.Pink),
		gameobjects.NewWall(0, 0, 640, 40, r.Orange),
		gameobjects.NewWall(600, 0, 40, 640, r.Gold),
	})

	/*********************Game Loop*********************/
	for !r.WindowShouldClose() {
		/************Update************/
		system.EscapeCheck()
		cam.Update(player.Update(r.GetFrameTime(), room.Walls))
		/***********Drawing***********/
		r.BeginDrawing()
		r.BeginMode2D(r.Camera2D(*cam))
		{
			r.ClearBackground(r.Black)

			DrawTileMap(tileMap)
			player.Draw()
		}
		r.EndMode2D()
		r.EndDrawing()
	}
	r.CloseWindow()
}

// DrawTileMap is just a debug function
// to display the tile map until separated
// into another file.
func DrawTileMap(tilemap *tiled.Map) {

	tilemapImage := system.GetTexture("assets\\graphics\\BasicTileSet.png")
	tileW := float32(tilemap.TileWidth)
	tileH := float32(tilemap.TileHeight)

	for _, layer := range tilemap.Layers {
		for tileIndex, tile := range layer.Tiles {

			id := int32(tile.ID)

			tileDestX := int32(tileIndex%tilemap.Width) * int32(tileW)
			tileDestY := int32(tileIndex) / int32(tilemap.Width) * int32(tileH)

			var tileSrcX int32 = id % (tilemapImage.Width / int32(tileW))
			var tileSrcY int32 = int32(id / tilemapImage.Width / int32(tileW))

			//fmt.Println(id)
			//fmt.Println(tileSrcX)

			// Something wrong with tiles being mixed up with each other.
			r.DrawTexturePro(tilemapImage,
				r.NewRectangle(float32(tileSrcX)*tileW, float32(tileSrcY)*tileH, tileW, tileH),
				r.NewRectangle(float32(tileDestX), float32(tileDestY), tileW, tileH),
				r.NewVector2(0, 0), 0, r.White)
		}
	}
}
