package board

import (
	"Snake-Go/config"
)

type Board struct {
	length, width int
	cellSize      int
	Size          int
}

func NewBoard() *Board {
	LengthValue, err := config.GetConfig("BoardWidth")
	if err != nil {
		LengthValue = 600.0
		println("BoardWidth is not set in config.json")
	}

	WidthValue, err := config.GetConfig("BoardLength")
	if err != nil {
		WidthValue = 600.0
		println("BoardLength is not set in config.json")
	}

	CellValue, err := config.GetConfig("CellSize")
	if err != nil {
		CellValue = 20
		println("CellSize is not set in config.json")
	}

	Length, ok := LengthValue.(float64)
	if !ok {
		println("BoardWidth in config.json is not a float64")
		Length = 600.0
	}

	Width, ok := WidthValue.(float64)
	if !ok {
		println("BoardLength in config.json is not a float64")
		Width = 600.0
	}

	Cell, ok := CellValue.(float64)
	if !ok {
		println("CellSize in config.json is not a float64")
		Cell = 20
	}

	return &Board{length: int(Length), width: int(Width), cellSize: int(Cell), Size: int(Length / Cell)}
}

func (b *Board) GetSize() int {
	return b.Size
}
func (b *Board) GetCell() int {
	return b.cellSize
}
