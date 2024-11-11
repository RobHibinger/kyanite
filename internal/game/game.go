package game

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

func RunGame() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Kyanite")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)

	LoadResources()

	g, err := NewGame()
	if err != nil {
		fmt.Printf("Failed to load new game: %v", err)
		return
	}

	if err := ebiten.RunGame(&g); err != nil {
		fmt.Printf("Failed to run game: %v", err)
	}
}

func NewGame() (Game, error) {
	g := Game{
		Debug: false,
		Camera: Camera{
			Scale: Vec2{
				x: 2.0,
				y: 2.0,
			},
		},
		GameState: GameState{
			Player: CreateMoveableEntity(ResouceType_Knight, Vec2{}, Vec2{}),
			Entities: []Entity{
				CreateMoveableEntity(ResouceType_Knight, Vec2{x: 100, y: 100}, Vec2{}),
			},
		},
	}

	return g, nil
}

type GameState struct {
	Player   Entity
	Entities []Entity
}

type Camera struct {
	Position Vec2
	Scale    Vec2
}

type Game struct {
	Debug                     bool
	Count                     int
	Camera                    Camera
	GameState                 GameState
	InputState                InputState
	AnimSpriteResources       map[ResouceType]Resource
	ScreenWidth, ScreenHeight int
}

func (g *Game) Update() error {
	HandleInput(g)

	g.GameState.Player.Velocity = g.InputState.MoveDirection
	g.GameState.Player.Position.x += g.GameState.Player.Velocity.x * g.GameState.Player.AddedSpeedMultiplier
	g.GameState.Player.Position.y += g.GameState.Player.Velocity.y * g.GameState.Player.AddedSpeedMultiplier
	g.Camera.Position = g.GameState.Player.Position

	UpdateDirectionAnim(&g.GameState.Player)

	return nil
}

func (game *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{36, 54, 66, 0})

	if game.Debug {
		ebitenutil.DebugPrint(screen, fmt.Sprintf("FPS: %v\nSpeed: %v", ebiten.ActualFPS(), game.GameState.Player.AddedSpeedMultiplier))
		vector.DrawFilledCircle(screen, float32(game.ScreenWidth)/2, float32(game.ScreenHeight)/2, 5, color.Black, false)
	}

	RenderEntity(screen, &game.GameState.Player, &game.Camera)
	for _, e := range game.GameState.Entities {
		RenderEntity(screen, &e, &game.Camera)
	}

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	g.ScreenWidth, g.ScreenHeight = outsideWidth, outsideHeight
	return g.ScreenWidth, g.ScreenHeight
}
