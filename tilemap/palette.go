package tilemap

import (
	"dota3/assets"
	"github.com/hajimehoshi/ebiten/v2"
)

type Tile int

const (
	Air Tile = iota
	Grass
	Wall
)

var Tiles = map[Tile]*ebiten.Image{
	Grass: assets.Images["grass"],
	Wall:  assets.Images["wall"],
}
