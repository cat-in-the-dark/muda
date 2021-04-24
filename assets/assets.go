package assets

import (
	"embed"
	"github.com/hajimehoshi/ebiten/v2"
	"image"
	_ "image/png"
	"log"
	"path"
)

//go:embed textures
var textures embed.FS

//go:embed sounds
var sounds embed.FS


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
