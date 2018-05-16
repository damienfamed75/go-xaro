package gameobjects

import (
	"reflect"
	"testing"

	"github.com/gen2brain/raylib-go/raylib"
)

func TestNewCamera(t *testing.T) {
	tests := []struct {
		name string
		want XaroCamera
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewCamera(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCamera() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestXaroCamera_Update(t *testing.T) {
	type args struct {
		diff raylib.Vector2
		curr raylib.Vector2
	}
	tests := []struct {
		name string
		c    *XaroCamera
		args args

		want raylib.Vector2
	}{
		{name: "testCamera", c: NewCamera(), args: args{diff: raylib.NewVector2(1, 1), curr: raylib.NewVector2(2, 2)}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.c.Update(tt.args.diff, tt.args.curr)
		})

		tt.want = raylib.Vector2{X: -tt.args.diff.X, Y: -tt.args.diff.Y}
		if got := tt.c.Offset; !reflect.DeepEqual(got, tt.want) {
			t.Errorf("Offset = %v, want %v", got, tt.want)
		}
	}
}
