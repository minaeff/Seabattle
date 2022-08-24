package main

// Описание игрового поля

type GameArea struct {
	Width  int
	Height int
	Cell   Cell
}

//cell coords
type Cell struct {
	X int
	Y int
}
