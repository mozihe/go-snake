package tool

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"image/color"
	"math"
)

type Point struct {
	x, y int
}

type Vector struct {
	x, y int
}

func NewVector(X, Y int) *Vector {
	return &Vector{x: X, y: Y}
}

func NewPoint(X, Y int) *Point {
	return &Point{x: X, y: Y}
}

func GetDistance(p1, p2 Point) float64 {
	return math.Pow(math.Pow(float64(p1.x-p2.x), 2)+math.Pow(float64(p1.y-p2.y), 2), 0.5)
}

func Find[T comparable](arr []T, target T) int {
	for i, v := range arr {
		if v == target {
			return i
		}
	}
	return -1
}

func (p *Point) GetX() int {
	return p.x
}

func (p *Point) GetY() int {
	return p.y
}

func (p *Vector) GetX() int {
	return p.x
}

func (p *Vector) GetY() int {
	return p.y
}

func DrawBlock(screen *ebiten.Image, cell int, p Point, c color.RGBA) {
	vector.DrawFilledRect(screen, float32(p.GetX()*cell), float32(p.GetY()*cell), float32(cell), float32(cell), c, true)
}

func CompareVector(v1, v2 Vector) bool {
	return v1.x == v2.x && v1.y == v2.y
}

func ComparePoint(p1, p2 Point) bool {
	return p1.x == p2.x && p1.y == p2.y
}
