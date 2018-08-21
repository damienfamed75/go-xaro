package main

import (
	"math"
	"strconv"

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
	tileMap := gameobjects.NewTileMap(system.MapPath)
	player := gameobjects.NewPlayer(r.NewVector2(float32(system.HalfW()), float32(system.HalfH())))
	cam := gameobjects.NewCamera()

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

			// offx, offy := cam.Offset.X-float32(player.Texture.Width/2), cam.Offset.Y-float32(player.Texture.Height/2)
			// x, y := int32((r.GetMousePosition().X)+298), int32((r.GetMousePosition().Y)+218)

			x, y := int32((r.GetMousePosition().X)-cam.Offset.X), int32((r.GetMousePosition().Y)-cam.Offset.Y)
			b, c := float32(x)-(player.Position.X+float32(player.Ase.FrameWidth/2)), float32(y)-(player.Position.Y+float32(player.Ase.FrameHeight/2))
			angle := math.Atan2(float64(b), float64(c))
			var col r.Color

			switch {
			case angle < 0:
				col = r.Red
			case angle > 0:
				col = r.Blue
			default:
				col = r.Black
			}

			r.DrawText(strconv.FormatFloat(angle, 'f', 6, 64), int32(player.Position.X)-10, int32(player.Position.Y)-10, 1, col)
			//x, y := int32(r.GetMousePosition().X-cam.Offset.X), int32(r.GetMousePosition().Y-cam.Offset.Y)
			r.DrawCircle(int32(float32(r.GetScreenWidth())/cam.Zoom), int32(float32(r.GetScreenHeight())/cam.Zoom), 5, r.White)
			r.DrawCircle(x, y, 5, col)
			r.DrawLine(int32(player.Position.X+float32(player.Ase.FrameWidth/2)), int32(player.Position.Y+float32(player.Ase.FrameHeight/2)), x, y, col)
			player.Draw()
		}
		r.EndMode2D()
		r.EndDrawing()
	}
}
