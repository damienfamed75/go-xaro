package gameobjects

import (
	"fmt"
	"os"

	"github.com/gen2brain/raylib-go/raylib"

	"github.com/damienfamed75/go-xaro/system"
	tiled "github.com/lafriks/go-tiled"
)

// XaroTileMap is the Xaro version of a go-tiled tilemap
type XaroTileMap tiled.Map

// NewTileMap will return a new go-tiled
// tilemap that would be exported from the
// Tiled program.
func NewTileMap(path string) *XaroTileMap {
	m, err := tiled.LoadFromFile(path)
	if err != nil {
		fmt.Println("Couldn't parse map:")
		fmt.Println(path)
		os.Exit(2)
	}
	mp := XaroTileMap(*m)

	return &mp
}

// Draw renders the Tiled tilemap layer by layer. Top left to right
func (tilemap *XaroTileMap) Draw(tileset string) {

	// Load in entire tilemap
	tilemapImage := system.GetTexture(tileset)
	tileW := float32(tilemap.TileWidth)  // width of one tile
	tileH := float32(tilemap.TileHeight) // height of one tile

	/**********For each layer in the map**********/
	for _, layer := range tilemap.Layers {
		/****For each tile within the tile map****/
		for tileIndex, tile := range layer.Tiles {

			id := int32(tile.ID) // the id based on the position in tileset

			// the destination where the sprites will be rendered
			// tileDestX := int32(tileIndex%tilemap.Width) * int32(tileW)
			tileDestX := int32(tileIndex%tilemap.Width) * int32(tileW)
			tileDestY := int32(int32(tileIndex)/int32(tilemap.Width)) * int32(tileH)
			// tileDestY := int32(tileIndex / tilemap.Width * int(tileW))

			// tileSrc are the coordinates that the sprites appear in the tilesheet
			tileSrcX := id % (tilemapImage.Width / int32(tileW))
			tileSrcY := int32(id / (tilemapImage.Width / int32(tileW)))

			source := raylib.NewRectangle(float32(tileSrcX)*tileW, float32(tileSrcY)*tileH, tileW, tileH)

			var rotation float32

			if tile.HorizontalFlip {
				source.Width *= -1
			}

			if tile.VerticalFlip {
				source.Height *= -1
			}

			if tile.DiagonalFlip {
				source.Width *= -1
				rotation = 90
			}

			// Draw the tile using the specific tile.
			raylib.DrawTexturePro(tilemapImage,
				source,
				raylib.NewRectangle(float32(tileDestX), float32(tileDestY), tileW, tileH),
				raylib.NewVector2(tileW/2, tileH/2), rotation, raylib.White)
		}
	}
}
