package gamelogic

import (
	"fmt"
	"minesweeper/config"
)

type GameState struct {
	bombs                 [config.BOARD_HEIGHT][config.BOARD_WIDTH]bool
	surroundingBombsCount [config.BOARD_HEIGHT][config.BOARD_WIDTH]uint
	visible               [config.BOARD_HEIGHT][config.BOARD_WIDTH]bool
	tickedByUser          [config.BOARD_HEIGHT][config.BOARD_WIDTH]bool
	gameOver              bool
	gameWon               bool
}

func UpdateLogic(gameState *GameState, x uint, y uint) {
	fmt.Printf("%v %v \n", x, y)
}
