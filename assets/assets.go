package assets

import (
	"embed"
	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/font/opentype"
	"image"
	_ "image/png"
	"log"
	"path"
)

//go:embed textures
var textures embed.FS

//go:embed fonts
var fonts embed.FS

// LoadImage reads an image by name from the embedded textures folder.
// Panic on error
func LoadImage(name string) *ebiten.Image {
	imagePath := path.Join("textures", name)
	file, err := textures.Open(imagePath)
	if err != nil {
		log.Fatal(err)
	}
	img, _, err := image.Decode(file)
	if err != nil {
		log.Fatal(err)
	}

	return ebiten.NewImageFromImage(img)
}

func LoadFont(name string) *opentype.Font {
	fontPath := path.Join("fonts", name)
	file, err := fonts.ReadFile(fontPath)
	if err != nil {
		log.Fatal(err)
	}
	font, err := opentype.Parse(file)
	if err != nil {
		log.Fatal(err)
	}
	return font
}
