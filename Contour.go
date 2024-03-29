package gui


import (
	"image/color"

	"golang.org/x/image/font/basicfont"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2"
)

type Contour struct {
	Position Point
	Size Point
	Text string
}

func (c *Contour) Draw (screen *ebiten.Image) {
	/*
		0,0 -> 0,y
		0,y -> x,y
		x,0 -> x,y
		0,t -> x,0
	*/
	
	text.Draw(screen, c.Text, basicfont.Face7x13, int(c.Position.X + 5), int(c.Position.Y), color.RGBA{0xff, 0x00, 0x00, 0xff})
	ebitenutil.DrawLine(screen, c.Position.X, c.Position.Y - 13, c.Position.X, c.Position.Y + c.Size.Y, color.RGBA{0xff, 0x00, 0x00, 0xff})
	ebitenutil.DrawLine(screen, c.Position.X, c.Position.Y + c.Size.Y, c.Position.X + c.Size.X, c.Position.Y + c.Size.Y, color.RGBA{0xff, 0x00, 0x00, 0xff})
	ebitenutil.DrawLine(screen, c.Position.X + c.Size.X, c.Position.Y, c.Position.X + c.Size.X, c.Position.Y + c.Size.Y, color.RGBA{0xff, 0x00, 0x00, 0xff})
	ebitenutil.DrawLine(screen, c.Position.X + float64((len(c.Text) * 7) + 10), c.Position.Y, c.Position.X + c.Size.X, c.Position.Y, color.RGBA{0xff, 0x00, 0x00, 0xff})
}

func (c *Contour) Input (g *SceneManager) {
	return
}
