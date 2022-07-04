package main

import (
	"fmt"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"luminoso.dev/rezgin/game"
)

func main() {
	ebiten.SetWindowSize(640, 480)
	if DEV_MODE {
		ebiten.SetWindowTitle(fmt.Sprintf("%v %v (DEV)", NAME, VERSION))
	} else {
		ebiten.SetWindowTitle(fmt.Sprintf("%v %v", NAME, VERSION))
	}
	game := game.Game{
		Debug: DEV_MODE,
		FS:    embedFS,
	}
	game.Init()
	if err := ebiten.RunGame(&game); err != nil {
		log.Fatal(err)
	}
}
