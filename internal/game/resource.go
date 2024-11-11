package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var GameResources map[ResouceType]*Resource

type ResouceType int

const (
	ResouceType_Knight ResouceType = iota
)

type ResourceFrameType int

const (
	ResourceFrameType_North ResourceFrameType = iota
	ResourceFrameType_South
	ResourceFrameType_East
	ResourceFrameType_West
)

type Resource struct {
	Width, Height int
	Image         *ebiten.Image
	Frames        map[ResourceFrameType][]Vec2
}

func LoadResources() error {
	knightSpritesheet, _, err := ebitenutil.NewImageFromFile("assets/images/Knight/SpriteSheet.png")
	if err != nil {
		return err
	}

	GameResources = map[ResouceType]*Resource{
		ResouceType_Knight: {
			Image:  knightSpritesheet,
			Height: 16,
			Width:  16,
			Frames: map[ResourceFrameType][]Vec2{
				ResourceFrameType_North: {
					{x: 1, y: 0},
					{x: 1, y: 1},
				},
				ResourceFrameType_South: {
					{x: 0, y: 0},
					{x: 0, y: 1},
				},
				ResourceFrameType_East: {
					{x: 3, y: 0},
					{x: 3, y: 1},
				},
				ResourceFrameType_West: {
					{x: 2, y: 0},
					{x: 2, y: 1},
				},
			},
		},
	}

	return nil
}
