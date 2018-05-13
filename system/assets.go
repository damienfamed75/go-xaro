package system

import (
	"fmt"

	"github.com/gen2brain/raylib-go/raylib"
)

// LoadedTextures are all the loaded sprites
var LoadedTextures = make(map[string]raylib.Texture2D)

// GetTexture returns the picture of a spritesheet
func GetTexture(textureName string) raylib.Texture2D {
	_, ok := LoadedTextures[textureName]
	if !ok {
		LoadedTextures[textureName] = raylib.LoadTexture(textureName)
	}

	fmt.Println("Loaded Textures: ")
	fmt.Println(LoadedTextures)

	return LoadedTextures[textureName]
}
