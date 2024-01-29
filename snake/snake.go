package snake

import "Snake-Go/tool"

type Snake struct {
	Head      tool.Point
	Body      []tool.Point
	Speed     int
	Direction tool.Vector
	OldBody   tool.Point
}

func (s *Snake) GetLength() int {
	return len(s.Body) + 1
}

func NewSnake(Size int) *Snake {
	head := tool.NewPoint(int(Size/2), int(Size/2))
	oneBody := tool.NewPoint(int(Size/2), int(Size/2)-1)
	var body []tool.Point
	body = append([]tool.Point{*oneBody}, body...)
	return &Snake{Head: *head, Body: body, Speed: 2, Direction: *tool.NewVector(0, -1)}
}

func (s *Snake) Update() {
	s.Body = append([]tool.Point{s.Head}, s.Body...)
	s.OldBody = s.Body[len(s.Body)-1]
	s.Body = s.Body[:len(s.Body)-1]
	s.Head = *tool.NewPoint(s.Head.GetX()+s.Direction.GetX(), s.Head.GetY()+s.Direction.GetY())
}
