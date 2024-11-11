package game

import (
	"fmt"
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

func RenderEntity(screen *ebiten.Image, entity *Entity, camera *Camera) {
	animState := entity.SpriteAnimState[entity.CurrentSpriteAnimState]
	resource := animState.Resource
	currentResourceFrame := resource.Frames[animState.ResourceFrameType][animState.Index]

	fmt.Printf("rft: %v\n", animState.ResourceFrameType)

	screenWidth, screenHeight := screen.Bounds().Dx(), screen.Bounds().Dy()
	screenCenterX, screenCenterY := float64(screenWidth)/2, float64(screenHeight)/2
	halfSizeX, halfSizeY := float64(resource.Width)*camera.Scale.x/2, float64(resource.Height)*camera.Scale.y/2

	opts := ebiten.DrawImageOptions{}
	opts.GeoM.Scale(camera.Scale.x, camera.Scale.y)
	opts.GeoM.Translate(entity.Position.x-camera.Position.x+screenCenterX-halfSizeX, entity.Position.y-camera.Position.y+screenCenterY-halfSizeY)

	screen.DrawImage(
		resource.Image.SubImage(
			image.Rect(
				int(currentResourceFrame.x)*resource.Width,
				int(currentResourceFrame.y)*resource.Height,
				int(currentResourceFrame.x)*resource.Width+resource.Width,
				int(currentResourceFrame.y)*resource.Height+resource.Height,
			),
		).(*ebiten.Image),
		&opts,
	)
}
