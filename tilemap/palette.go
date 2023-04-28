package tilemap

import (
	"dota3/assets"
	"github.com/hajimehoshi/ebiten/v2"
)

type Tile int

const (
	Air Tile = iota
	Grass0
	Grass1
	Grass2
	Grass3
	Wall
)

var Tiles = map[Tile]*ebiten.Image{
	Grass0: assets.Images["grass0"],
	Grass1: assets.Images["grass1"],
	Grass2: assets.Images["grass2"],
	Grass3: assets.Images["grass3"],
	Wall:   assets.Images["wall"],
}
