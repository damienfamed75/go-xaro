package gameobjects

import (
	"reflect"
	"testing"

	"github.com/gen2brain/raylib-go/raylib"
)

func TestNewWall(t *testing.T) {
	type args struct {
		x float32
		y float32
		w float32
		h float32
		c raylib.Color
	}
	tests := []struct {
		name string
		args args
		want Wall
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewWall(tt.args.x, tt.args.y, tt.args.w, tt.args.h, tt.args.c); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewWall() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWall_Draw(t *testing.T) {
	tests := []struct {
		name string
		w    *Wall
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.w.Draw()
		})
	}
}

func TestDrawPlatforms(t *testing.T) {
	type args struct {
		walls []Wall
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			DrawPlatforms(tt.args.walls...)
		})
	}
}
