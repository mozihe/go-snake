package main

import (
	"Snake-Go/board"
	"Snake-Go/food"
	"Snake-Go/input"
	"Snake-Go/snake"
	"Snake-Go/tool"
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"golang.org/x/image/font/basicfont"
	"image/color"
	"log"
	"math/rand"
	"time"
)

var ifOver bool = false

type Game struct {
	board *board.Board
	food  *food.Food
	snake *snake.Snake
	Size  int
	Score int
}

func newGame() *Game {
	myBoard := board.NewBoard()
	size := myBoard.GetSize()
	mySnake := snake.NewSnake(size)
	unable := mySnake.Body
	unable = append([]tool.Point{mySnake.Head}, unable...)
	myFood := food.NewFood(unable, size)
	return &Game{board: myBoard, Size: size, snake: mySnake, food: myFood, Score: 0}
}

func (g *Game) Update() error {
	if ifOver {

		return nil
	}

	g.snake.Direction = *input.Update(&g.snake.Direction)
	ebiten.SetTPS(g.snake.Speed)
	g.snake.Update()
	if g.ifOver() {
		ifOver = true
	}
	if g.ifEat() {
		g.eat()
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	if ifOver {
		screen.Fill(color.Black)
		text.Draw(screen, "Game Over", basicfont.Face7x13, 280, 300, color.White)
		scoreStr := fmt.Sprintf("Score: %d", g.Score)
		text.Draw(screen, scoreStr, basicfont.Face7x13, 280, 400, color.White)
		return
	}

	screen.Fill(color.Black)
	tool.DrawBlock(screen, g.board.GetCell(), g.snake.Head, color.RGBA{255, 0, 0, 255})
	for _, oneBody := range g.snake.Body {
		tool.DrawBlock(screen, g.board.GetCell(), oneBody, color.RGBA{0, 255, 0, 255})
	}
	switch g.food.GetAbility() {
	case food.Speedup:
		tool.DrawBlock(screen, g.board.GetCell(), *tool.NewPoint(g.food.GetX(), g.food.GetY()), color.RGBA{255, 255, 0, 255})
	case food.Slowdown:
		tool.DrawBlock(screen, g.board.GetCell(), *tool.NewPoint(g.food.GetX(), g.food.GetY()), color.RGBA{0, 0, 255, 255})
	default:
		tool.DrawBlock(screen, g.board.GetCell(), *tool.NewPoint(g.food.GetX(), g.food.GetY()), color.RGBA{255, 0, 255, 255})
	}
	vector.DrawFilledRect(screen, 0, 600, 600, 1, color.RGBA{255, 255, 255, 255}, true)

	scoreStr := fmt.Sprintf("Score: %d Speed: %d", g.Score, g.snake.Speed)
	text.Draw(screen, scoreStr, basicfont.Face7x13, 10, 615, color.White)

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 600, 620
}

func main() {

	rand.Seed(time.Now().UnixNano())
	ebiten.SetWindowSize(600, 620) // 设置窗口大小
	ebiten.SetWindowTitle("snake") // 设置窗口标题

	game := newGame()
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}

}

func (g *Game) ifOver() bool {
	if g.snake.Head.GetX() < 0 || g.snake.Head.GetX() >= g.Size || g.snake.Head.GetY() < 0 || g.snake.Head.GetY() >= g.Size {
		return true
	}
	for _, oneBody := range g.snake.Body {
		if tool.ComparePoint(g.snake.Head, oneBody) {
			return true
		}
	}
	return false
}

func (g *Game) ifEat() bool {
	if tool.ComparePoint(g.snake.Head, *tool.NewPoint(g.food.GetX(), g.food.GetY())) {
		return true
	}
	return false
}

func (g *Game) eat() {
	g.snake.Body = append(g.snake.Body, g.snake.OldBody)
	if g.food.GetAbility() == food.Speedup && g.snake.Speed < 5 {
		g.snake.Speed++
	}
	if g.food.GetAbility() == food.Slowdown && g.snake.Speed > 1 {
		g.snake.Speed--
	}
	g.Score++
	g.food = food.NewFood(g.snake.Body, g.Size)
}
