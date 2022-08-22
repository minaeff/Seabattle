package main

// Ships description

type Ship struct {
	ID    int
	Name  string
	Class ClassShip
}

// Type of ship

type сlassShip struct {
	ID          int
	ClassName   string // линкор и т.д.
	Length      int
	HealthPoint int // для расчета статистики урона
}
