package main

import (
	"minesweeper/gamelogic"
	"minesweeper/uicontroller"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "Minesweeper",
		Bounds: pixel.R(0, 0, 1024, 768),
	}

	uiState := uicontroller.UiState{}
	gameState := gamelogic.GameState{}

	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	for !win.Closed() {
		//x, y := uicontroller.GetInput()
		uicontroller.RenderBoard(gameState, &uiState)
		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}
