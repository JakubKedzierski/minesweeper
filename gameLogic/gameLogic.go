package gamelogic
import "minesweeper/config"

type GameState struct {
	bombs                 [config.BOARD_HEIGHT][config.BOARD_WIDTH]bool
	surroundingBombsCount [config.BOARD_HEIGHT][config.BOARD_WIDTH]uint
}
