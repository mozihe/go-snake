package main

import (
	"Snake-Go/board"
	"Snake-Go/food"
	"Snake-Go/snake"
	"Snake-Go/tool"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
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
	ebiten.SetMaxTPS(g.snake.Speed)
	g.snake.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.Black)
	ebitenutil.DrawRect(screen, float64(g.snake.Head.GetX()*g.board.GetCell()), float64(g.snake.Head.GetY()*g.board.GetCell()), float64(g.board.GetCell()), float64(g.board.GetCell()), color.RGBA{255, 0, 0, 255})
	for _, oneBody := range g.snake.Body {
		ebitenutil.DrawRect(screen, float64(oneBody.GetX()*g.board.GetCell()), float64(oneBody.GetY()*g.board.GetCell()), float64(g.board.GetCell()), float64(g.board.GetCell()), color.RGBA{255, 0, 0, 255})
	}
	ebitenutil.DrawRect(screen, float64(g.food.GetX()*g.board.GetCell()), float64(g.food.GetY()*g.board.GetCell()), float64(g.board.GetCell()), float64(g.board.GetCell()), color.RGBA{0, 0, 255, 255})

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
