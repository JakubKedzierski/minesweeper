package uicontroller

import "minesweeper/config"

type UiState struct {
	visible      [config.BOARD_HEIGHT][config.BOARD_WIDTH]bool
	tickedByUser [config.BOARD_HEIGHT][config.BOARD_WIDTH]bool
}

func GetInput() (uint, uint) {
	return 0, 0
}

func RenderBoard() {

}
