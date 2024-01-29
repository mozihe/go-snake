package main

import (
	"Snake-Go/board"
	"Snake-Go/food"
	"Snake-Go/snake"
	"Snake-Go/tool"
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"
	"log"
)

type Game struct {
	board *board.Board
	food  *food.Food
	snake *snake.Snake
	Size  int
}

func newGame() *Game {
	myBoard := board.NewBoard()
	size := myBoard.GetSize()
	mySnake := snake.NewSnake(size)
	unable := mySnake.Body
	unable = append([]tool.Point{mySnake.Head}, unable...)
	myFood := food.NewFood(unable, size)
	return &Game{board: myBoard, Size: size, snake: mySnake, food: myFood}
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.Black)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 600, 600
}

func main() {

	ebiten.SetWindowSize(600, 600) // 设置窗口大小
	ebiten.SetWindowTitle("snake") // 设置窗口标题

	game := newGame()
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}

}
