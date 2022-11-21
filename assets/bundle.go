package assets

import (
	"bytes"
	"embed"
	"image"
	_ "image/png"
	"log"

	"github.com/tinne26/etxt"

	"github.com/hajimehoshi/ebiten/v2"
)

//go:embed data
var fs embed.FS
var Images = make(map[string]*ebiten.Image)

//var Fonts = make(map[string]*etxt.Font)

func Init() {
	Images["player"] = LoadImage(fs, "data/image/player.png")
	//Fonts["start"] = LoadFont(fs, "data/font/roboto-regular.ttf")
}

func LoadImage(fs embed.FS, path string) *ebiten.Image {
	data, err := fs.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	img, _, err := image.Decode(bytes.NewReader(data))
	if err != nil {
		log.Fatal(err)
	}
	return ebiten.NewImageFromImage(img)
}

func LoadFont(fs embed.FS, path string) *etxt.Font {
	f, _, err := etxt.ParseEmbedFontFrom(path, fs)
	if err != nil {
		log.Fatal(err)
	}

	return f
}