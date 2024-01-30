package input

import (
	"Snake-Go/tool"
	"github.com/hajimehoshi/ebiten/v2"
)

func Update(CurrentVector *tool.Vector) *tool.Vector {
	var input []rune
	input = ebiten.AppendInputChars(input)
	if len(input) > 0 {
		switch input[len(input)-1] {
		case 'w':
			if !tool.CompareVector(*CurrentVector, *tool.NewVector(0, 1)) {
				return tool.NewVector(0, -1)
			}
		case 's':
			if !tool.CompareVector(*CurrentVector, *tool.NewVector(0, -1)) {
				return tool.NewVector(0, 1)
			}
		case 'a':
			if !tool.CompareVector(*CurrentVector, *tool.NewVector(1, 0)) {
				return tool.NewVector(-1, 0)
			}
		case 'd':
			if !tool.CompareVector(*CurrentVector, *tool.NewVector(-1, 0)) {
				return tool.NewVector(1, 0)
			}
		}

	}
	return CurrentVector
}
