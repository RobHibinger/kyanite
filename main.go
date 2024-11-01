package main

import (
	"fmt"
	"image"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Vec3 struct {
	x, y, z float64
}

type Entity struct {
	position Vec3
	image    *ebiten.Image
}

func (e *Entity) Draw(screen *ebiten.Image) {
	opts := ebiten.DrawImageOptions{}
	opts.GeoM.Translate(e.position.x, e.position.y)

	screen.DrawImage(
		e.image.SubImage(
			image.Rect(0, 0, 16, 16),
		).(*ebiten.Image),
		&opts,
	)
}

type GameState struct {
	player   Entity
	entities []Entity
}

type Game struct {
	debug      bool
	game_state GameState
}

func (g *Game) HandleInput() {
	if ebiten.IsKeyPressed(ebiten.KeyF1) && inpututil.IsKeyJustPressed(ebiten.KeyF1) {
		g.debug = !g.debug
	}

	if ebiten.IsKeyPressed(ebiten.KeyW) {
		g.game_state.player.position.y += -1
	} else if ebiten.IsKeyPressed(ebiten.KeyS) {
		g.game_state.player.position.y += 1
	}

	if ebiten.IsKeyPressed(ebiten.KeyA) {
		g.game_state.player.position.x -= 1
	} else if ebiten.IsKeyPressed(ebiten.KeyD) {
		g.game_state.player.position.x += 1
	}
}

func (g *Game) Update() error {
	g.HandleInput()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{36, 54, 66, 0})

	if g.debug {
		ebitenutil.DebugPrint(screen, fmt.Sprintf("FPS: %v", ebiten.ActualFPS()))
	}

	g.game_state.player.Draw(screen)

	for _, e := range g.game_state.entities {
		e.Draw(screen)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello, World!")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)

	playerImg, _, err := ebitenutil.NewImageFromFile("assets/images/Knight/SpriteSheet.png")
	if err != nil {
		log.Fatal(err)
	}

	var g = Game{
		debug: false,
		game_state: GameState{
			player: Entity{
				image: playerImg,
				position: Vec3{
					x: 100.0,
					y: 100.0,
					z: 100.0,
				},
			},
		},
	}

	if err := ebiten.RunGame(&g); err != nil {
		log.Fatal(err)
	}
}
