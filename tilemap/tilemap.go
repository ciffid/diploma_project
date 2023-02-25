package tilemap

import (
	"dota3/assets"
	"dota3/display"
	"dota3/random"
	"github.com/hajimehoshi/ebiten/v2"
)

type Tilemap struct {
	w, h  int
	tiles []Tile
}

func NewTilemap(w, h int) *Tilemap {
	t := &Tilemap{
		w:     w,
		h:     h,
		tiles: make([]Tile, w*h),
	}
	for i := 0; i < len(t.tiles); i++ {
		tile := Grass
		if random.Chance(0.2) {
			tile = Wall
		}
		t.tiles[i] = tile
	}
	return t
}

func (t *Tilemap) Draw(screen *ebiten.Image, camera *display.Camera) {
	for y := 0; y < t.h; y++ {
		for x := 0; x < t.w; x++ {
			tile := t.get(x, y)
			if tile == Air {
				continue
			}
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(x*assets.TileSize), float64(y*assets.TileSize))
			op.GeoM.Translate(-camera.X, -camera.Y)
			screen.DrawImage(Tiles[tile], op)
		}
	}
}

func (t *Tilemap) outside(x, y int) bool {
	xOut := x < 0 || x >= t.w
	yOut := y < 0 || y >= t.h
	return xOut || yOut
}

func (t *Tilemap) index(x, y int) int {
	return y*t.w + x
}

func (t *Tilemap) get(x, y int) Tile {
	if t.outside(x, y) {
		return Air
	}
	return t.tiles[t.index(x, y)]
}

func (t *Tilemap) set(x, y int, tile Tile) {
	if t.outside(x, y) {
		return
	}
	t.tiles[t.index(x, y)] = tile
}
