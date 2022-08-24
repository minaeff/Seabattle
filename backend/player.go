package main

/* Описание структуры типа Player */

// Игрок

type Player struct {
	Name     string
	Nickname string
	ID       int
	Rank     Rank
}

// Звание

type Rank struct {
	ID    int
	Rank  string
	Stars []byte // погоны
}
