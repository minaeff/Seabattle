package main

// Ships description

type Ship struct{
	ID int
	Name string
	Class ClassShip
}

// Type of ship

type ClassShip struct{
	ID int
	ClassName string
	Length int
	HealthPoint int
}