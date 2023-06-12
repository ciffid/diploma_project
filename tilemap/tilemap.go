package tilemap

import (
	"DP/assets"
	"DP/display"
	"DP/helper/random"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/solarlune/resolv"
)

type Tilemap struct {
	w, h      int
	tiles     []Tile
	boxes     []*resolv.Object
	offscreen *ebiten.Image
}

func NewTilemap(w, h int, space *resolv.Space) *Tilemap {
	t := &Tilemap{
		w:         w,
		h:         h,
		tiles:     make([]Tile, w*h),
		offscreen: ebiten.NewImage(w*assets.TileSize, h*assets.TileSize),
	}
	for y := 0; y < t.h; y++ {
		for x := 0; x < t.w; x++ {
			tile := Grass0
			switch {
			case random.Chance(0.2):
				tile = Grass1
			case random.Chance(0.2):
				tile = Grass2
			case random.Chance(0.2):
				tile = Grass3
			case random.Chance(0.1):
				tile = Wall
				obj := resolv.NewObject(float64(x), float64(y), 1, 1)
				space.Add(obj)
				t.boxes = append(t.boxes, obj)
			}
			t.set(x, y, tile)
		}
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
