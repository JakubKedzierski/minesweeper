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
		//x, y := uicontroller.GetInput()
		uicontroller.RenderBoard(gameState, &uiState)
		uiState.Win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}
