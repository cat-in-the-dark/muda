package sszb

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
)

type Message struct {
	text          string
	font          font.Face
	key           ebiten.Key
	framesShow    float64
	framesElapsed float64
}

func (m *Message) IsExpired() bool {
	if m.framesElapsed >= m.framesShow {
		return true
	}

	if inpututil.IsKeyJustPressed(m.key) {
		return true
	}

	return false
}

func NewMessage(text string, framesShow float64) *Message {
	return &Message{
		text:          text,
		font:          DefaultFont,
		key:           ebiten.KeyF,
		framesShow:    framesShow,
		framesElapsed: 0,
	}
}

type Hud struct {
	pos       *Vector2
	size      *Vector2
	messages  []*Message
	lw        float64 // line width
	txtOffset *Vector2
}

func NewHud() *Hud {
	return &Hud{
		pos:       NewVector2((ScreenWidth-HudWidth)/2, ScreenHeight-HudHeight-HudOffsetY),
		size:      NewVector2(HudWidth, HudHeight),
		messages:  make([]*Message, 0),
		lw:        2,
		txtOffset: NewVector2(8, 8+FontSize),
	}
}

func (h *Hud) Show(message *Message) {
	if len(h.messages) != 0 {
		// Ignore same messages
		if h.messages[len(h.messages)-1].text == message.text {
			return
		}
	}
	h.messages = append(h.messages, message)
}

func (h *Hud) ShowReset(message *Message) {
	h.messages = make([]*Message, 1)
	h.messages[0] = message
}

func (h *Hud) Draw(screen *ebiten.Image) {
	if len(h.messages) == 0 {
		return
	}

	msg := h.messages[0]

	ebitenutil.DrawRect(screen, h.pos.x+3, h.pos.y+3, h.size.x, h.size.y, ColorHudLine)
	ebitenutil.DrawRect(screen, h.pos.x, h.pos.y, h.size.x, h.size.y, ColorHudLine)
	ebitenutil.DrawRect(screen, h.pos.x+h.lw, h.pos.y+h.lw, h.size.x-h.lw*2, h.size.y-h.lw*2, ColorHudBody)

	text.Draw(screen, msg.text, msg.font, int(h.pos.x+h.txtOffset.x), int(h.pos.y+h.txtOffset.y), ColorText)

	h.messages[0].framesElapsed++
	if h.messages[0].IsExpired() {
		h.messages = h.messages[1:]
	}
}

func (h *Hud) Update() {}
