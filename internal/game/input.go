package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type InputState struct {
	Debug          bool
	IncrementSpeed bool
	DecrementSpeed bool
	MoveDirection  Vec2
}

func HandleInput(inputState *InputState) {
	if inpututil.IsKeyJustPressed(ebiten.KeyF1) {
		inputState.Debug = !inputState.Debug
	}

	inputState.IncrementSpeed = inpututil.IsKeyJustPressed(ebiten.KeyKPAdd)
	inputState.DecrementSpeed = inpututil.IsKeyJustPressed(ebiten.KeyKPSubtract)

	inputState.MoveDirection = Vec2{}
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		inputState.MoveDirection.y = -1
	} else if ebiten.IsKeyPressed(ebiten.KeyS) {
		inputState.MoveDirection.y = 1
	}

	if ebiten.IsKeyPressed(ebiten.KeyA) {
		inputState.MoveDirection.x = -1
	} else if ebiten.IsKeyPressed(ebiten.KeyD) {
		inputState.MoveDirection.x = 1
	}
}
