package game

import (
	"math"
	"time"
)

type SpriteAnimState struct {
	Resource          *Resource
	ResourceFrameType ResourceFrameType
	Index             int
	StartTime         time.Time
	Duration          time.Duration
}

type Vec2 struct {
	x, y float64
}

type EntityAnimType int

const (
	EntityAnimType_FaceNorth EntityAnimType = iota
	EntityAnimType_FaceSouth
	EntityAnimType_FaceEast
	EntityAnimType_FaceWest
)

type Entity struct {
	Position               Vec2
	Scale                  Vec2
	Velocity               Vec2
	AddedSpeedMultiplier   float64
	SpriteAnimState        map[EntityAnimType]*SpriteAnimState
	CurrentSpriteAnimState EntityAnimType
}

func UpdateDirectionAnim(entity *Entity) {
	faceDirection := GetFaceDirection(entity)
	spriteAnimState := entity.SpriteAnimState[faceDirection]
	if entity.CurrentSpriteAnimState != faceDirection {
		spriteAnimState.Index = 0
		spriteAnimState.StartTime = time.Now()
		entity.CurrentSpriteAnimState = faceDirection
	}

	isMoving := math.Sqrt(entity.Velocity.x*entity.Velocity.x+entity.Velocity.y*entity.Velocity.y) > 0
	if isMoving {
		if time.Now().After(spriteAnimState.StartTime.Add(spriteAnimState.Duration)) {
			spriteAnimState.Index++
			spriteAnimState.StartTime = time.Now()

			if spriteAnimState.Index >= len(spriteAnimState.Resource.Frames[spriteAnimState.ResourceFrameType]) {
				spriteAnimState.Index = 0
			}
		}
	} else {
		spriteAnimState.Index = 0
	}
}

func GetFaceDirection(entity *Entity) EntityAnimType {
	if entity.Velocity.y < 0 {
		return EntityAnimType_FaceNorth
	} else if entity.Velocity.y > 0 {
		return EntityAnimType_FaceSouth
	} else if entity.Velocity.x > 0 {
		return EntityAnimType_FaceEast
	} else if entity.Velocity.x < 0 {
		return EntityAnimType_FaceWest
	}
	return entity.CurrentSpriteAnimState
}

func CreateMoveableEntity(resourceType ResouceType, position Vec2, scale Vec2) Entity {
	return Entity{
		AddedSpeedMultiplier:   1.0,
		CurrentSpriteAnimState: EntityAnimType_FaceSouth,
		SpriteAnimState: map[EntityAnimType]*SpriteAnimState{
			EntityAnimType_FaceNorth: {
				Resource:          GameResources[resourceType],
				ResourceFrameType: ResourceFrameType_North,
				StartTime:         time.Now(),
				Duration:          time.Millisecond * 200,
			},
			EntityAnimType_FaceSouth: {
				Resource:          GameResources[resourceType],
				ResourceFrameType: ResourceFrameType_South,
				StartTime:         time.Now(),
				Duration:          time.Millisecond * 200,
			},
			EntityAnimType_FaceEast: {
				Resource:          GameResources[resourceType],
				ResourceFrameType: ResourceFrameType_East,
				StartTime:         time.Now(),
				Duration:          time.Millisecond * 200,
			},
			EntityAnimType_FaceWest: {
				Resource:          GameResources[resourceType],
				ResourceFrameType: ResourceFrameType_West,
				StartTime:         time.Now(),
				Duration:          time.Millisecond * 200,
			},
		},
		Position: position,
		Scale: Vec2{
			x: 1.0,
			y: 1.0,
		},
	}
}
