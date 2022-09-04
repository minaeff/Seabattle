package main

// Описание игрового поля

type GameArea struct {
	Width  int
	Height int
}

// массив ячеек (игровое поле)
type Matrix [][]Cell

// ячейка
type Cell struct {
	IsEmpty bool
}

// mark cell if player
func (c Cell) MarkAsEmpty() {
	c.IsEmpty = true
}
