package main

import (
	"minesweeper/gamelogic"
	"minesweeper/uicontroller"

	"github.com/faiface/pixel/pixelgl"
)

func run() {

	// TODO:
	// - handle game over and game won

	uiState := uicontroller.UiState{}
	gameState := gamelogic.GameState{}
	uicontroller.InitWindow(&uiState)
	gamelogic.InitGameState(&gameState)

	for !uiState.Win.Closed() {
		userInput, x, y := uicontroller.GetInput(&uiState)
		if userInput != gamelogic.NoInput {
			gamelogic.UpdateLogic(&gameState, x, y, userInput)
		}
		uicontroller.RenderBoard(gameState, &uiState)
	}
}

func main() {
	pixelgl.Run(run)
}
