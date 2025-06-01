package gamelogic

import "minesweeper/config"

type GameState struct {
	bombs                 [config.BOARD_HEIGHT][config.BOARD_WIDTH]bool
	surroundingBombsCount [config.BOARD_HEIGHT][config.BOARD_WIDTH]uint
	visible               [config.BOARD_HEIGHT][config.BOARD_WIDTH]bool
	tickedByUser          [config.BOARD_HEIGHT][config.BOARD_WIDTH]bool
}

func TickBox(gameState *GameState, x uint, y uint){
	
}
