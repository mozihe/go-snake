package board

type Board struct {
	length, width int
	cellSize      int
	Size          int
}

func NewBoard() *Board {

	return &Board{length: 600, width: 600, cellSize: 20, Size: 30}
}

func (b *Board) GetSize() int {
	return b.Size
}
func (b *Board) GetCell() int {
	return b.cellSize
}
