package tilemap

import (
	"DP/assets"
	"DP/display"
	"DP/helper/random"
	"github.com/hajimehoshi/ebiten/v2"
)

type Tilemap struct {
	w, h      int
	tiles     []Tile
	offscreen *ebiten.Image
}

func NewTilemap(w, h int) *Tilemap {
	t := &Tilemap{
		w:         w,
		h:         h,
		tiles:     make([]Tile, w*h),
		offscreen: ebiten.NewImage(w*assets.TileSize, h*assets.TileSize),
	}
	for i := 0; i < len(t.tiles); i++ {
		tile := Grass0
		if random.Chance(0.7) {
			tile = Grass1
			if random.Chance(0.5) {
				tile = Grass2
				if random.Chance(0.3) {
					tile = Grass3
				}
			}
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
			t.offscreen.DrawImage(Tiles[tile].Image(), op)
		}
	}
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(camera.Scale, camera.Scale)
	op.GeoM.Translate(-camera.X, -camera.Y)
	screen.DrawImage(t.offscreen, op)
	t.offscreen.Clear()
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
