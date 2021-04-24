package lib

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image"
)

type AnimationSystem struct {
	ticks int // this is number of ticks since the application start
}

type Animation struct {
	frames        []int
	sheet         *SpriteSheet
	ticksPerFrame int
	looping       bool
	finished      bool
	system        *AnimationSystem // read only!
}

// NewAnimationSystem creates an animation system object.
// AnimationSystem supposed to be a Singleton but you can do whatever you wish.
func NewAnimationSystem() *AnimationSystem {
	return &AnimationSystem{
		ticks: 0,
	}
}

func (system *AnimationSystem) Update() {
	system.ticks += 1
}

func (system *AnimationSystem) NewLooping(sheet *SpriteSheet, ticksPerFrame int, frames []int) *Animation {
	return &Animation{
		frames:        frames,
		sheet:         sheet,
		ticksPerFrame: ticksPerFrame,
		looping:       true,
		finished:      false,
		system:        system,
	}
}

func (system *AnimationSystem) NewNormal(sheet *SpriteSheet, ticksPerFrame int, frames []int) *Animation {
	return &Animation{
		frames:        frames,
		sheet:         sheet,
		ticksPerFrame: ticksPerFrame,
		looping:       false,
		finished:      false,
		system:        system,
	}
}

func (anim *Animation) GetFrame() *ebiten.Image {
	if !anim.looping && anim.finished {
		return anim.sheet.Last()
	}
	anim.finished = false

	frameIdx := anim.system.ticks / anim.ticksPerFrame
	if frameIdx >= len(anim.frames) {
		anim.finished = true
	}

	if !anim.looping && anim.finished {
		return anim.sheet.Last()
	}
	
	idx := anim.frames[frameIdx % len(anim.frames)]
	return anim.sheet.At(idx)
}

func (anim *Animation) GetSpriteSize() (int, int) {
	return anim.sheet.spriteWidth, anim.sheet.spriteHeight
}

type SpriteSheet struct {
	spriteWidth  int
	spriteHeight int
	sprites      []*ebiten.Image
}

func (sheet *SpriteSheet) At(idx int) *ebiten.Image {
	return sheet.sprites[idx%len(sheet.sprites)]
}

func (sheet *SpriteSheet) Last() *ebiten.Image {
	return sheet.sprites[len(sheet.sprites) - 1]
}

func NewSpriteSheet(
	texture *ebiten.Image,
	spriteWidth int,
	spriteHeight int,
) *SpriteSheet {
	w, h := texture.Size()
	cols := w / spriteWidth
	rows := h / spriteHeight

	sprites := make([]*ebiten.Image, cols*rows)

	for i := 0; i < cols; i++ {
		for j := 0; j < rows; j++ {
			rect := image.Rect(i*spriteWidth, j*spriteHeight, (i+1)*spriteWidth, (j+1)*spriteHeight)
			sprites[i+j*cols] = texture.SubImage(rect).(*ebiten.Image)
		}
	}

	return &SpriteSheet{
		spriteWidth:  spriteWidth,
		spriteHeight: spriteHeight,
		sprites:      sprites,
	}
}
