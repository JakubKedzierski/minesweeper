package gamelogic

import "minesweeper/config"

type GameState struct {
	bombs                 [config.BOARD_HEIGHT][config.BOARD_WIDTH]bool
	surroundingBombsCount [config.BOARD_HEIGHT][config.BOARD_WIDTH]uint
	visible               [config.BOARD_HEIGHT][config.BOARD_WIDTH]bool
	tickedByUser          [config.BOARD_HEIGHT][config.BOARD_WIDTH]bool
}

func UpdateLogic(gameState *GameState, x uint, y uint) {

}
