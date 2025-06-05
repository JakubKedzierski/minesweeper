package gamelogic

import (
	"minesweeper/config"
)

type GameState struct {
	Bombs                 [config.BOARD_HEIGHT][config.BOARD_WIDTH]bool
	SurroundingBombsCount [config.BOARD_HEIGHT][config.BOARD_WIDTH]uint
	Visible               [config.BOARD_HEIGHT][config.BOARD_WIDTH]bool
	Flags                 [config.BOARD_HEIGHT][config.BOARD_WIDTH]bool
	GameOver              bool
	GameWon               bool
}

type UserInput uint

const (
	NoInput UserInput = iota
	TickBox
	Flag
)

func InitGameState(gameState *GameState) {

}

func UpdateLogic(gameState *GameState, x uint, y uint, userInput UserInput) {
	if gameState.Visible[y][x] {
		return
	}

	gameState.Visible[y][x] = true
	if userInput == Flag {
		gameState.Flags[y][x] = true
		return
	}

	if gameState.Bombs[y][x] {
		gameState.GameOver = true
	} else {
		// check if game is over and user won
		for x := range config.BOARD_WIDTH {
			for y := range config.BOARD_HEIGHT {
				if gameState.Bombs[y][x] {
					continue
				} else if !gameState.Visible[y][x] {
					return
				}
			}
		}

		gameState.GameWon = true
	}
}
