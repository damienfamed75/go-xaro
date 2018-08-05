package system

import (
	"github.com/gen2brain/raylib-go/raylib"
)

// LoadedTextures are all the loaded sprites
var LoadedTextures = make(map[string]raylib.Texture2D)

// GetTexture returns the picture of a spritesheet
func GetTexture(textureName string) raylib.Texture2D {
	if _, ok := LoadedTextures[textureName]; !ok {
		LoadedTextures[textureName] = raylib.LoadTexture(textureName)
	}

	return LoadedTextures[textureName]
}
