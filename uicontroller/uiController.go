package uicontroller

//import "minesweeper/config"
import (
	"image"
	_ "image/png"
	"minesweeper/config"
	"minesweeper/gamelogic"
	"os"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

type UiState struct {
	Win     *pixelgl.Window
	Sprites pixel.Picture
}

func InitWindow(uiState *UiState) {
	cfg := pixelgl.WindowConfig{
		Title:  "Minesweeper",
		Bounds: pixel.R(0, 0, config.WINDOW_WIDTH, config.WINDOW_HEIGHT),
		VSync:  true,
	}

	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	win.SetPos(pixel.Vec{450, 100})

	sprites, err := loadSpritesPicture("sprites/sprites_sheet.png")
	if err != nil {
		panic(err)
	}

	uiState.Sprites = sprites
	uiState.Win = win
}

func loadSpritesPicture(path string) (pixel.Picture, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}

	return pixel.PictureDataFromImage(img), nil
}

func GetInput(uiState *UiState) (gamelogic.UserInput, uint, uint) {
	leftPressed := uiState.Win.JustPressed(pixelgl.MouseButtonLeft)
	rightPressed := uiState.Win.JustPressed(pixelgl.MouseButtonRight)
	if leftPressed || rightPressed {
		mousePos := uiState.Win.MousePosition()
		if mousePos.X < config.BOX_LEN || mousePos.Y < config.BOX_LEN {
			return gamelogic.NoInput, 0, 0
		}

		mousePos.X -= (config.BOX_LEN)
		mousePos.Y -= (config.BOX_LEN)

		xBox := mousePos.X / config.BOX_LEN
		yBox := mousePos.Y / config.BOX_LEN

		if xBox < config.BOARD_WIDTH && yBox < config.BOARD_HEIGHT {
			if leftPressed {
				return gamelogic.TickBox, uint(xBox), uint(yBox)
			} else {
				return gamelogic.Flag, uint(xBox), uint(yBox)
			}
		}
	}
	return gamelogic.NoInput, 0, 0
}

func RenderBoard(gameState gamelogic.GameState, uiState *UiState) {
	uiState.Win.Clear(colornames.Lightgray)

	const SPRITES_Y_FLIP = 704
	const SPRITE_BOX_LEN = 16
	const SCALE = config.BOX_LEN / SPRITE_BOX_LEN

	const Y_LINE = SPRITES_Y_FLIP - 211
	boxSpriteRect := pixel.R(15, Y_LINE, 15+SPRITE_BOX_LEN, Y_LINE+SPRITE_BOX_LEN) // boxSprite location in pixels
	boxSprite := pixel.NewSprite(uiState.Sprites, boxSpriteRect)

	tickedBoxRect := pixel.R(31, Y_LINE, 31+SPRITE_BOX_LEN, Y_LINE+SPRITE_BOX_LEN)
	tickedBoxSprite := pixel.NewSprite(uiState.Sprites, tickedBoxRect)

	flagBoxRect := pixel.R(48, Y_LINE, 48+SPRITE_BOX_LEN, SPRITES_Y_FLIP-194)
	flagSprite := pixel.NewSprite(uiState.Sprites, flagBoxRect)

	const NUMBERS_COUNT = 8
	var numbersSprites [NUMBERS_COUNT]*pixel.Sprite
	for i := 0; i < NUMBERS_COUNT; i++ {
		leftVtx := 15 + float64(i*(SPRITE_BOX_LEN+1))
		numbersRect := pixel.R(leftVtx, SPRITES_Y_FLIP-227, leftVtx+SPRITE_BOX_LEN, SPRITES_Y_FLIP-211)
		numbersSprites[i] = pixel.NewSprite(uiState.Sprites, numbersRect)
	}

	for x := range config.BOARD_WIDTH {
		for y := range config.BOARD_HEIGHT {
			bottomLeftCornerX := float64(x*config.BOX_LEN + config.GUBARDBAND/2)
			bottomLeftCornerY := float64(y*config.BOX_LEN + config.GUBARDBAND/2)
			movLoc := pixel.Vec{X: bottomLeftCornerX, Y: bottomLeftCornerY}

			// choose sprite to draw based on game state:
			if !gameState.Visible[y][x] {
				boxSprite.Draw(uiState.Win, pixel.IM.Scaled(pixel.ZV, SCALE).Moved(movLoc))
			} else {
				if gameState.Flags[y][x] {
					flagSprite.Draw(uiState.Win, pixel.IM.Scaled(pixel.ZV, SCALE).Moved(movLoc))
				} else {
					numbersSprites[7].Draw(uiState.Win, pixel.IM.Scaled(pixel.ZV, SCALE).Moved(movLoc))
					tickedBoxSprite.Draw(uiState.Win, pixel.IM.Scaled(pixel.ZV, SCALE).Moved(movLoc))
				}
			}

		}
	}

	uiState.Win.Update()
}
