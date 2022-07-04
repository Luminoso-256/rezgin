package game

import (
	"embed"
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

// The Main Ebiten-based Game struct.
type Game struct {
	//needed to pass data from bootstrap
	Debug bool
	FS    embed.FS
	//all internal variables are lowercase.
	assets AssetRegistry
	px, py int
}

func (g *Game) Init() error {
	g.px = 0
	g.py = 0
	g.assets = LoadAssets(g.Debug, g.FS)
	return nil
}

func (g *Game) Update() error {
	if inpututil.IsKeyJustPressed(ebiten.KeyRight) {
		g.px += 1
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyLeft) {
		if g.px > 0 {
			g.px -= 1
		}
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyDown) {
		g.py += 1
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyUp) {
		if g.py > 0 {
			g.py -= 1
		}
	}
	return nil
}

var testmap = [][]uint8{
	{4, 4, 0, 0, 0, 0, 0, 0, 3, 3},
	{4, 4, 0, 0, 0, 0, 0, 0, 3, 3},
	{4, 4, 1, 0, 0, 2, 0, 0, 3, 3},
	{4, 4, 0, 0, 0, 2, 0, 0, 2, 0},
	{4, 4, 0, 1, 0, 2, 0, 2, 0, 0},
	{4, 4, 1, 0, 0, 2, 2, 0, 0, 0},
	{4, 4, 0, 0, 0, 2, 0, 0, 0, 0},
	{4, 4, 0, 0, 0, 0, 0, 0, 0, 0},
	{4, 4, 4, 4, 4, 4, 4, 4, 4, 4},
	{4, 4, 4, 4, 4, 4, 4, 4, 4, 4},
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.White)
	for i := 0; i < 8; i++ {
		for j := 8; j >= 0; j-- {
			op := &ebiten.DrawImageOptions{}
			//quasi-isometric. Pretty sure this is technically not quite correct
			//but it looks good so i'll take it.
			x := (float64(j) * 8) + (float64(i) * 8)
			y := (float64(i) * 4) - (float64(j) * 4)
			y += 42
			x -= 32
			op.GeoM.Translate(x, y)
			if len(testmap) > i+g.px && len(testmap[i+g.px]) > j+g.py {
				screen.DrawImage(g.assets.img[fmt.Sprintf("tile/p%v", testmap[i+g.px][j+g.py])], op)
			}
		}
	}

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(30, 32) // player doesn't move, world does.
	screen.DrawImage(g.assets.img["ent/player_temp"], op)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 64, 64 //Lowrez
}
