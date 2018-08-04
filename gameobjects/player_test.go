package gameobjects

import (
	"testing"

	"github.com/damienfamed75/go-xaro/system"
	"github.com/gen2brain/raylib-go/raylib"
	. "github.com/smartystreets/goconvey/convey"
)

func TestNewPlayer(t *testing.T) {
	Convey("Given I want a new player", t, func() {
		testplayer := NewPlayer()
		Convey("Then the player should have values", func() {
			So(testplayer, ShouldNotBeNil)
			So(testplayer.Ase, ShouldNotBeNil)
			So(testplayer.Texture, ShouldNotBeNil)
		})
	})
}

func TestPlayerUpdate(t *testing.T) {
	raylib.InitWindow(1, 1, "TEST")
	defer raylib.CloseWindow()

	testplayer := NewPlayer()

	Convey("Given I call the update function for positions", t, func() {
		diff, pos := testplayer.Update(raylib.GetFrameTime())
		Convey("Then the returned values should not be nil", func() {
			So(diff, ShouldNotBeNil)
			So(pos, ShouldNotBeNil)
			So(diff, ShouldNotEqual, pos)
		})
	})

}

func TestPlayerUpdateAnimation(t *testing.T) {
	raylib.InitWindow(1, 1, "TEST")
	system.Init()
	defer raylib.CloseWindow()

	testplayer := NewPlayer()

	Convey("Given I call the update function for animations", t, func() {
		testplayer.Update(raylib.GetFrameTime())
		defer func() { raylibIsKeyDown = raylib.IsKeyDown }()

		Convey("Given I press left", func() {
			dir := "left"

			raylibIsKeyDown = func(key int32) bool { return key == system.KeyBindings[dir] }
			testplayer.Update(raylib.GetFrameTime())

			Convey("Then animation should switch", func() {
				So(testplayer.Velocity.X, ShouldEqual, -120)
				So(testplayer.Ase.IsPlaying(dir), ShouldBeTrue)
			})
		})

		Convey("Given I press right", func() {
			dir := "right"

			raylibIsKeyDown = func(key int32) bool { return key == system.KeyBindings[dir] }
			testplayer.Update(raylib.GetFrameTime())

			Convey("Then animation should switch", func() {
				So(testplayer.Velocity.X, ShouldEqual, 120)
				So(testplayer.Ase.IsPlaying(dir), ShouldBeTrue)
			})
		})

		Convey("Given I press up", func() {
			dir := "up"

			raylibIsKeyDown = func(key int32) bool { return key == system.KeyBindings[dir] }
			testplayer.Update(raylib.GetFrameTime())

			Convey("Then animation should switch", func() {
				So(testplayer.Velocity.Y, ShouldEqual, -120)
				So(testplayer.Ase.IsPlaying(dir), ShouldBeTrue)
			})
		})

		Convey("Given I press down", func() {
			dir := "down"

			raylibIsKeyDown = func(key int32) bool { return key == system.KeyBindings[dir] }
			testplayer.Update(raylib.GetFrameTime())

			Convey("Then animation should switch", func() {
				So(testplayer.Velocity.Y, ShouldEqual, 120)
				So(testplayer.Ase.IsPlaying(dir), ShouldBeTrue)
			})
		})
	})
}

func TestPlayerDraw(t *testing.T) {
	testplayer := NewPlayer()

	Convey("Given I draw the player", t, func() {
		Convey("Then player shouldn't panic", func() {
			So(testplayer.Draw, ShouldNotPanic)
		})
	})
}
