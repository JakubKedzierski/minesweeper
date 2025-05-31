package main

import (
	"fmt"
	"minesweeper/uicontroller"
)

func main() {

	for {
		dx, dy := uicontroller.GetInput()
		fmt.Println(dx, dy)
		break
	}
}
