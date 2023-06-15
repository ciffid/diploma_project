package assets

import (
	"DP/graphics"
	"DP/helper/load"
	"DP/helper/split"
	"embed"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"time"

	"github.com/tinne26/etxt"
)

//go:embed data
var fs embed.FS
var Images = make(map[string]*graphics.Frameset)
var Fonts = make(map[string]*etxt.Font)

const (
	TileSize    = 32
	PlayerW     = 16
	PlayerH     = 32
	TileMapSize = 36
	Animation   = 100 * time.Millisecond
)

func init() {
	tiles := load.Image(fs, "data/image/tiles.png")
	player := load.Image(fs, "data/image/player-animation.png")
	enemy := load.Image(fs, "data/image/enemy-animation.png")

	Images["enemy_look"] = split.Single(enemy, PlayerW*4, 0, PlayerW, PlayerH, false)
	for f := 0; f < 2; f++ {
		flipped := f == 1
		postfix := "left"
		if flipped {
			postfix = "right"
		}
		Images["enemy_run_"+postfix] = split.Multi(enemy, PlayerW*0, 0, PlayerW, PlayerH, 3, flipped, Animation)
	}

	Images["enemy_run_down"] = split.Multi(enemy, PlayerW*3, 0, PlayerW, PlayerH, 3, false, Animation)
	Images["enemy_run_up"] = split.Multi(enemy, PlayerW*6, 0, PlayerW, PlayerH, 3, false, Animation)

	Images["player_look"] = split.Single(player, PlayerW*4, 0, PlayerW, PlayerH, false)

	for f := 0; f < 2; f++ {
		flipped := f == 1
		postfix := "left"
		if flipped {
			postfix = "right"
		}
		Images["player_run_"+postfix] = split.Multi(player, PlayerW*0, 0, PlayerW, PlayerH, 3, flipped, Animation)
	}

	Images["player_run_down"] = split.Multi(player, PlayerW*3, 0, PlayerW, PlayerH, 3, false, Animation)
	Images["player_run_up"] = split.Multi(player, PlayerW*6, 0, PlayerW, PlayerH, 3, false, Animation)

	Images["grass0"] = split.Single(tiles, TileSize*0, 0, TileSize, TileSize, false)
	Images["grass1"] = split.Single(tiles, TileSize*1, 0, TileSize, TileSize, false)
	Images["grass2"] = split.Single(tiles, TileSize*2, 0, TileSize, TileSize, false)
	Images["grass3"] = split.Single(tiles, TileSize*3, 0, TileSize, TileSize, false)
	Images["wall"] = split.Single(tiles, TileSize*7, 0, TileSize, TileSize, false)

	Images["background"] = graphics.NewFramesetSingle(load.Image(fs, "data/image/background.png"))
	Images["settings"] = graphics.NewFramesetSingle(load.Image(fs, "data/image/settings.png"))
	Images["start-location"] = graphics.NewFramesetSingle(load.Image(fs, "data/image/start-location.jpg"))

	Images["button-enabled"] = graphics.NewFramesetSingle(load.Image(fs, "data/image/button-enabled.png"))
	Images["button-disabled"] = graphics.NewFramesetSingle(load.Image(fs, "data/image/button-disabled.png"))

	Fonts["start"] = LoadFont(fs, "data/font/Motel King Medium(RUS by Slavchansky).ttf")
}

func LoadFont(fs embed.FS, path string) *etxt.Font {
	f, _, err := etxt.ParseEmbedFontFrom(path, fs)
	if err != nil {
		log.Fatal(err)
	}

	return f
}
