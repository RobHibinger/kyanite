package main

import (
	"fmt"
	"image"
	"image/color"
	"log"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type SpriteResouceType int

const (
	SpriteResouceType_KnightIdle SpriteResouceType = iota
	SpriteResouceType_KnightWalkNorth
	SpriteResouceType_KnightWalkEast
	SpriteResouceType_KnightWalkSouth
	SpriteResouceType_KnightWalkWest
)

type SpriteResource struct {
	Width, Height int
	Image         *ebiten.Image
	Frames        []Vec2
}

type SpriteState struct {
	Type  SpriteResouceType
	Index int
}

type SpriteAnimState struct {
	SpriteState
	StartTime time.Time
	Duration  time.Duration
}

func (sas *SpriteAnimState) Update() {

}

type Vec2 struct {
	x, y float64
}

type Entity struct {
	Position        Vec2
	Scale           Vec2
	Velocity        Vec2
	Sprite          SpriteState
	SpriteAnimState SpriteAnimState
}

func (e *Entity) Draw(screen *ebiten.Image, game *Game) {
	opts := ebiten.DrawImageOptions{}
	opts.GeoM.Scale(e.Scale.x, e.Scale.y)
	opts.GeoM.Translate(e.Position.x, e.Position.y)

	sprite := game.AnimSpriteResources[e.Sprite.Type]
	currentCell := sprite.Frames[e.Sprite.Index]
	screen.DrawImage(
		sprite.Image.SubImage(
			image.Rect(int(currentCell.x)*sprite.Width, int(currentCell.y)*sprite.Height, int(currentCell.x)*sprite.Width+sprite.Width, int(currentCell.y)*sprite.Height+sprite.Height),
		).(*ebiten.Image),
		&opts,
	)
}

type InputState struct {
	MoveDirection Vec2
}

type GameState struct {
	Player   Entity
	Entities []Entity
}

type Game struct {
	Debug               bool
	GameState           GameState
	InputState          InputState
	AnimSpriteResources map[SpriteResouceType]SpriteResource
}

func (g *Game) HandleInput() {
	if ebiten.IsKeyPressed(ebiten.KeyF1) && inpututil.IsKeyJustPressed(ebiten.KeyF1) {
		g.Debug = !g.Debug
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

func (g *Game) Update() error {
	g.HandleInput()

	g.GameState.Player.Velocity = g.InputState.MoveDirection
	g.GameState.Player.Position.x += g.GameState.Player.Velocity.x
	g.GameState.Player.Position.y += g.GameState.Player.Velocity.y

	if g.InputState.MoveDirection.y < 0 {
		// set player animation state to walk north
	} else if g.InputState.MoveDirection.y > 0 {
		// set player animation state to walk south
	} else if g.InputState.MoveDirection.x < 0 {
		// set player animation state to walk west
	} else if g.InputState.MoveDirection.x > 0 {
		// set player animation state to walk east
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{36, 54, 66, 0})

	if g.Debug {
		ebitenutil.DebugPrint(screen, fmt.Sprintf("FPS: %v", ebiten.ActualFPS()))
	}

	g.GameState.Player.Draw(screen, g)

	for _, e := range g.GameState.Entities {
		e.Draw(screen, g)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello, World!")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)

	knightIdle, _, err := ebitenutil.NewImageFromFile("assets/images/Knight/SeparateAnim/Idle.png")
	if err != nil {
		log.Fatal(err)
	}

	KnightWalk, _, err := ebitenutil.NewImageFromFile("assets/images/Knight/SeparateAnim/Walk.png")
	if err != nil {
		log.Fatal(err)
	}

	var g = Game{
		Debug: false,
		AnimSpriteResources: map[SpriteResouceType]SpriteResource{
			SpriteResouceType_KnightIdle: {
				Image:  knightIdle,
				Height: 16,
				Width:  16,
				Frames: []Vec2{
					{x: 0, y: 0},
				},
			},
			SpriteResouceType_KnightWalkNorth: {
				Image:  KnightWalk,
				Height: 16,
				Width:  16,
				Frames: []Vec2{
					{x: 1, y: 0},
					{x: 1, y: 1},
					{x: 1, y: 2},
					{x: 1, y: 3},
				},
			},
			SpriteResouceType_KnightWalkEast: {
				Image:  KnightWalk,
				Height: 16,
				Width:  16,
				Frames: []Vec2{
					{x: 3, y: 0},
					{x: 3, y: 1},
					{x: 3, y: 2},
					{x: 3, y: 3},
				},
			},
			SpriteResouceType_KnightWalkSouth: {
				Image:  KnightWalk,
				Height: 16,
				Width:  16,
				Frames: []Vec2{
					{x: 0, y: 0},
					{x: 0, y: 1},
					{x: 0, y: 2},
					{x: 0, y: 3},
				},
			},
			SpriteResouceType_KnightWalkWest: {
				Image:  KnightWalk,
				Height: 16,
				Width:  16,
				Frames: []Vec2{
					{x: 2, y: 0},
					{x: 2, y: 1},
					{x: 2, y: 2},
					{x: 2, y: 3},
				},
			},
		},
		GameState: GameState{
			Player: Entity{
				Sprite: SpriteState{
					Type:  SpriteResouceType_KnightIdle,
					Index: 0,
				},
				Position: Vec2{
					x: 100.0,
					y: 100.0,
				},
				Scale: Vec2{
					x: 2.0,
					y: 2.0,
				},
			},
		},
	}

	if err := ebiten.RunGame(&g); err != nil {
		log.Fatal(err)
	}
}
