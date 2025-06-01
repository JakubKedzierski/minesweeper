package uicontroller

//import "minesweeper/config"
import (
	"image"
	_ "image/png"
	"minesweeper/config"
	"minesweeper/gamelogic"
	"os"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
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

	spriteRect := pixel.R(32, 64, 64, 96) // sprite location in pixels
	sprite := pixel.NewSprite(uiState.Sprites, spriteRect)
	sprite.Draw(uiState.Win, pixel.IM.Moved(uiState.Win.Bounds().Center()))

	imd := imdraw.New(nil)
	imd.Color = colornames.Black

	for x := range config.BOARD_WIDTH {
		for y := range config.BOARD_HEIGHT {
			bottomLeftCornerX := float64(x + x*config.RECT_WIDTH + config.GUBARDBAND/2)
			bottomLeftCornerY := float64(y + y*config.RECT_HEIGHT + config.GUBARDBAND/2)
			topRightCornerX := bottomLeftCornerX + config.RECT_WIDTH
			topRightCornerY := bottomLeftCornerY + config.RECT_HEIGHT

			imd.Push(pixel.V(bottomLeftCornerX, bottomLeftCornerY), pixel.V(topRightCornerX, topRightCornerY))
			imd.Rectangle(config.RECT_LINE_WIDTH)
		}
	}

	imd.Polygon(0)
	imd.Draw(uiState.Win)
}
