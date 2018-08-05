package system

import (
	"fmt"
	"os"

	"github.com/gen2brain/raylib-go/raylib"
	"github.com/spf13/viper"
)

var (
	KeyBindings = make(map[string]int32)
	Config      Configuration
)

// EscapeCheck is used to close the window and gets called
// int the update portion of the game loop.

// HalfH returns half of the screen height
func HalfH() int32 { return raylib.GetScreenHeight() / 2 }

// HalfW returns half of the screen width
func HalfW() int32 { return raylib.GetScreenWidth() / 2 }

// Init initializes the controls
func Init() {
	_, Config = LoadViperConfig("/assets/config/")

	KeyBindings["left"] = Config.Controls.Left
	KeyBindings["right"] = Config.Controls.Right
	KeyBindings["up"] = Config.Controls.Up
	KeyBindings["down"] = Config.Controls.Down
	KeyBindings["action"] = raylib.MouseLeftButton
	KeyBindings["menu"] = Config.Controls.Menu
}

func LoadViperConfig(file string) (*viper.Viper, Configuration) {
	v := viper.New()
	wd, _ := os.Getwd()

	v.SetConfigName("config.development")
	v.AddConfigPath("$HOME/.go-xaro")
	v.AddConfigPath(wd + file)
	v.AddConfigPath(".")
	v.AddConfigPath(file)
	if err := v.ReadInConfig(); err != nil {
		panic(fmt.Errorf("fatal error config file: %s", err))
	}

	var c Configuration
	if err := v.Unmarshal(&c); err != nil {
		fmt.Printf("couldn't read config: %s", err)
	}

	//mode := v.GetString("window.mode")
	switch winmode := v.GetString("window.winmode"); winmode {
	case "fullscreen":
		c.Window.Mode = raylib.FlagFullscreenMode
	case "windowed":
		c.Window.Mode = raylib.FlagWindowResizable
	default:
		ChangeConfig(v, "window.winmode", "windowed")
		c.Window.Mode = raylib.FlagWindowResizable
	}

	return v, c
}

// ChangeConfig updates the current config file's value
func ChangeConfig(v *viper.Viper, key string, value interface{}) {
	v.Set(key, value)
	if err := v.WriteConfig(); err != nil {
		fmt.Printf("couldn't write config: %s", err)
	}
}
