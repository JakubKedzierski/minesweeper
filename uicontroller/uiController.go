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

func GetInput() (uint, uint) {
	return 0, 0
}

func RenderBoard(gamestate gamelogic.GameState, uiState *UiState) {
	uiState.Win.Clear(colornames.Lightgray)

	const SPRITES_Y_FLIP = 704
	const SCALE_BOX_SPRITE = 1.2
	boxSpriteRect := pixel.R(15, SPRITES_Y_FLIP-211, 32, SPRITES_Y_FLIP-194) // boxSprite location in pixels
	boxSprite := pixel.NewSprite(uiState.Sprites, boxSpriteRect)

	for x := range config.BOARD_WIDTH {
		for y := range config.BOARD_HEIGHT {
			bottomLeftCornerX := float64(x + x*int(boxSpriteRect.W()*SCALE_BOX_SPRITE) + config.GUBARDBAND/2)
			bottomLeftCornerY := float64(y + y*int(boxSpriteRect.H()*SCALE_BOX_SPRITE) + config.GUBARDBAND/2)

			movLoc := pixel.Vec{X: bottomLeftCornerX, Y: bottomLeftCornerY}
			boxSprite.Draw(uiState.Win, pixel.IM.Scaled(pixel.ZV, SCALE_BOX_SPRITE).Moved(movLoc))
		}
	}
}
