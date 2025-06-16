package gamelogic

import (
	"math/rand/v2"
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
	for range config.BOMB_COUNT {
		x := rand.IntN(config.BOARD_WIDTH)
		y := rand.IntN(config.BOARD_HEIGHT)
		for gameState.Bombs[y][x] {
			x = rand.IntN(config.BOARD_WIDTH)
			y = rand.IntN(config.BOARD_HEIGHT)
		}
		gameState.Bombs[y][x] = true
	}

	initBombSurrounding(gameState)
}

func initBombSurrounding(gameState *GameState) {
	for y := range config.BOARD_HEIGHT {
		for x := range config.BOARD_WIDTH {
			if gameState.Bombs[y][x] {
				continue
			}

			var bombCount uint = 0
			for dx := -1; dx <= 1; dx++ {
				for dy := -1; dy <= 1; dy++ {
					if dx == 0 && dy == 0 {
						continue
					}

					neighbourX := x + dx
					neighbourY := y + dy
					if neighbourX < 0 || neighbourX >= config.BOARD_WIDTH {
						continue
					}
					if neighbourY < 0 || neighbourY >= config.BOARD_HEIGHT {
						continue
					}

					if gameState.Bombs[neighbourY][neighbourX] {
						bombCount++
					}
				}
			}

			gameState.SurroundingBombsCount[y][x] = bombCount
		}
	}
}

func UpdateLogic(gameState *GameState, x uint, y uint, userInput UserInput) {
	if gameState.Visible[y][x] && !gameState.Flags[y][x] {
		return
	}

	gameState.Visible[y][x] = true
	if userInput == Flag {
		gameState.Flags[y][x] = true
		return
	}

	gameState.Flags[y][x] = false

	if gameState.Bombs[y][x] {
		gameState.GameOver = true
	} else {
		gameState.GameWon = checkGameWin(gameState)

		if !gameState.GameWon {
			uncoverEmptyBoxes(gameState, x, y)
		}
	}
}

func checkGameWin(gameState *GameState) bool {
	for x := range config.BOARD_WIDTH {
		for y := range config.BOARD_HEIGHT {
			if gameState.Bombs[y][x] {
				continue
			} else if !gameState.Visible[y][x] {
				return false
			}
		}
	}

	return true
}

func uncoverEmptyBoxes(gameState *GameState, x, y uint) {
	type Dir struct {
		dx int
		dy int
	}

	directions := [...]Dir{
		{-1, 0},
		{1, 0},
		{0, -1},
		{0, 1},
	}

	for _, direction := range directions {
		neighbourX := int(x) + direction.dx
		neighbourY := int(y) + direction.dy
		if neighbourX < 0 || neighbourX >= config.BOARD_WIDTH {
			continue
		}

		if neighbourY < 0 || neighbourY >= config.BOARD_HEIGHT {
			continue
		}

		if gameState.Visible[neighbourY][neighbourX]{
			continue
		} 
		
		if !gameState.Bombs[neighbourY][neighbourX] &&
			gameState.SurroundingBombsCount[neighbourY][neighbourX] == 0 {
			gameState.Visible[neighbourY][neighbourX] = true
			uncoverEmptyBoxes(gameState, uint(neighbourX), uint(neighbourY))
		}
	}
}
