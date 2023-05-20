package tilemap

import (
	"DP/assets"
	"DP/graphics"
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

var Tiles = map[Tile]*graphics.Frameset{
	Grass0: assets.Images["grass0"],
	Grass1: assets.Images["grass1"],
	Grass2: assets.Images["grass2"],
	Grass3: assets.Images["grass3"],
	Wall:   assets.Images["wall"],
}
