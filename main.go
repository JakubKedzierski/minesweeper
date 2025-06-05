package main

import (
	"minesweeper/gamelogic"
	"minesweeper/uicontroller"
	"github.com/faiface/pixel/pixelgl"
)

func run() {

	uiState := uicontroller.UiState{}
	gameState := gamelogic.GameState{}
	uicontroller.InitWindow(&uiState)

	for !uiState.Win.Closed() {
		boxTicked, x, y := uicontroller.GetInput(&uiState)
		if boxTicked {
			gamelogic.UpdateLogic(&gameState, x, y)
		}
		uicontroller.RenderBoard(gameState, &uiState)
	}
}

func main() {
	pixelgl.Run(run)
}
