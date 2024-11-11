package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type InputState struct {
	MoveDirection Vec2
}

func HandleInput(g *Game) {
	if inpututil.IsKeyJustPressed(ebiten.KeyF1) {
		g.Debug = !g.Debug
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyKPAdd) {
		g.GameState.Player.AddedSpeedMultiplier += .1
	} else if inpututil.IsKeyJustPressed(ebiten.KeyKPSubtract) {
		g.GameState.Player.AddedSpeedMultiplier -= .1
	}

	g.InputState.MoveDirection = Vec2{}
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		g.InputState.MoveDirection.y = -1
	} else if ebiten.IsKeyPressed(ebiten.KeyS) {
		g.InputState.MoveDirection.y = 1
	}

	if ebiten.IsKeyPressed(ebiten.KeyA) {
		g.InputState.MoveDirection.x = -1
	} else if ebiten.IsKeyPressed(ebiten.KeyD) {
		g.InputState.MoveDirection.x = 1
	}
}
