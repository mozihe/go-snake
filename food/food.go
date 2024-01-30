package food

import (
	"Snake-Go/tool"
	"math/rand"
)

type Ability int

const (
	Speedup Ability = iota
	Slowdown
	None
)

type Food struct {
	x, y    int
	ability Ability
}

func NewFood(unable []tool.Point, Size int) *Food {
	myTool := rand.Intn(10)
	var myAbility Ability
	switch myTool {
	case 0:

		myAbility = Speedup
	case 1:
		myAbility = Slowdown
	default:
		myAbility = None
	}
	var X, Y int
	for {
		X = rand.Intn(Size)
		Y = rand.Intn(Size)
		myPoint := tool.NewPoint(X, Y)
		if tool.Find(unable, *myPoint) == -1 {
			break
		}
	}
	return &Food{x: X, y: Y, ability: myAbility}

}

func (f *Food) GetX() int {
	return f.x
}

func (f *Food) GetY() int {
	return f.y
}

func (f *Food) GetAbility() Ability {
	return f.ability
}
