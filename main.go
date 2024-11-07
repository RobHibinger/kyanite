package main

import (
	"fmt"
	"image"
	"image/color"
	"log"
	"math"
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

type EntityAnimType int

const (
	EntityAnimType_MoveNorth EntityAnimType = iota
	EntityAnimType_MoveSouth
	EntityAnimType_MoveEast
	EntityAnimType_MoveWest
)

type SpriteAnimState struct {
	Resource  *SpriteResource
	Index     int
	StartTime time.Time
	Duration  time.Duration
}

type Vec2 struct {
	x, y float64
}

type Entity struct {
	Position               Vec2
	Scale                  Vec2
	Velocity               Vec2
	AddedSpeedMultiplier   float64
	SpriteAnimState        map[EntityAnimType]*SpriteAnimState
	CurrentSpriteAnimState EntityAnimType
}

func (e *Entity) Draw(screen *ebiten.Image, game *Game) {
	opts := ebiten.DrawImageOptions{}
	opts.GeoM.Scale(e.Scale.x, e.Scale.y)
	opts.GeoM.Translate(e.Position.x, e.Position.y)

	animState := e.SpriteAnimState[e.CurrentSpriteAnimState]
	sprite := e.SpriteAnimState[e.CurrentSpriteAnimState].Resource
	currentCell := sprite.Frames[animState.Index]
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
	Count               int
	GameState           GameState
	InputState          InputState
	AnimSpriteResources map[SpriteResouceType]SpriteResource
}

func (g *Game) HandleInput() {
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

func GetFaceDirection(e *Entity) EntityAnimType {
	if e.Velocity.y < 0 {
		return EntityAnimType_MoveNorth
	} else if e.Velocity.y > 0 {
		return EntityAnimType_MoveSouth
	} else if e.Velocity.x > 0 {
		return EntityAnimType_MoveEast
	} else if e.Velocity.x < 0 {
		return EntityAnimType_MoveWest
	}
	return e.CurrentSpriteAnimState
}

func UpdateDirectionAnim(e *Entity) {
	et := GetFaceDirection(e)
	state := e.SpriteAnimState[et]
	if e.CurrentSpriteAnimState != et {
		state.Index = 0
		state.StartTime = time.Now()
	}
	e.CurrentSpriteAnimState = et

	isMoving := math.Sqrt(e.Velocity.x*e.Velocity.x+e.Velocity.y*e.Velocity.y) > 0
	if isMoving {
		if time.Now().After(state.StartTime.Add(state.Duration)) {
			state.Index++
			state.StartTime = time.Now()
			if state.Index >= len(state.Resource.Frames) {
				state.Index = 0
			}
		}
	} else {
		state.Index = 0
	}
}

func (g *Game) Update() error {
	g.HandleInput()

	g.GameState.Player.Velocity = g.InputState.MoveDirection
	g.GameState.Player.Position.x += g.GameState.Player.Velocity.x * g.GameState.Player.AddedSpeedMultiplier
	g.GameState.Player.Position.y += g.GameState.Player.Velocity.y * g.GameState.Player.AddedSpeedMultiplier
	UpdateDirectionAnim(&g.GameState.Player)

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{36, 54, 66, 0})

	if g.Debug {
		ebitenutil.DebugPrint(screen, fmt.Sprintf("FPS: %v\nSpeed: %v", ebiten.ActualFPS(), g.GameState.Player.AddedSpeedMultiplier))
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

	var resources = map[SpriteResouceType]*SpriteResource{
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
	}

	var g = Game{
		Debug: false,
		GameState: GameState{
			Player: Entity{
				AddedSpeedMultiplier:   1.0,
				CurrentSpriteAnimState: EntityAnimType_MoveSouth,
				SpriteAnimState: map[EntityAnimType]*SpriteAnimState{
					EntityAnimType_MoveNorth: {
						Resource:  resources[SpriteResouceType_KnightWalkNorth],
						StartTime: time.Now(),
						Duration:  time.Millisecond * 200,
					},
					EntityAnimType_MoveSouth: {
						Resource:  resources[SpriteResouceType_KnightWalkSouth],
						StartTime: time.Now(),
						Duration:  time.Millisecond * 200,
					},
					EntityAnimType_MoveEast: {
						Resource:  resources[SpriteResouceType_KnightWalkEast],
						StartTime: time.Now(),
						Duration:  time.Millisecond * 200,
					},
					EntityAnimType_MoveWest: {
						Resource:  resources[SpriteResouceType_KnightWalkWest],
						StartTime: time.Now(),
						Duration:  time.Millisecond * 200,
					},
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
