package main

import (
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	Tiles []MapTile
}

// NewGame creates a new Game object and initializes it
func NewGame() *Game {
	g := &Game{}
	g.Tiles = CreateTiles()
	return g
}

// Update is called each tic.
func (g *Game) Update() error {
	return nil
}

// Draw is called each draw cycle  and is where we will blit
func (g *Game) Draw(screen *ebiten.Image) {
	gd := NewGameData()
	for x := 0; x < gd.ScreenWidth; x++ {
		for y := 0; y < gd.ScreenHeight; y++ {
			tile := g.Tiles[GetIndexFromXY(x, y)]
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(tile.PixelX), float64(tile.PixelY))
			screen.DrawImage(tile.Image, op)
		}
	}
}

// Layout will return the screen dimensions.
func (g *Game) Layout(w, h int) (int, int) {
	return 1280, 800
}

func main() {
	g := NewGame()
	ebiten.SetWindowResizable(true)
	ebiten.SetWindowTitle("Tower")
	err := ebiten.RunGame(g)
	if err != nil {
		log.Fatal(err)
	}
}
