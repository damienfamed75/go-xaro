package gameobjects

import (
	"reflect"
	"testing"

	"github.com/gen2brain/raylib-go/raylib"
)

func TestNewPlayer(t *testing.T) {
	tests := []struct {
		name string
		want Player
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewPlayer(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPlayer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPlayer_Update(t *testing.T) {
	type args struct {
		dt    float32
		walls []Wall
	}
	tests := []struct {
		name  string
		p     *Player
		args  args
		want  raylib.Vector2
		want1 raylib.Vector2
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.p.Update(tt.args.dt, tt.args.walls)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Player.Update() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("Player.Update() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestPlayer_Draw(t *testing.T) {
	tests := []struct {
		name string
		p    *Player
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.p.Draw()
		})
	}
}
