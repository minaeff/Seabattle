package main

// Statistics for battle

type BattleStats struct {
	BattleID    int
	Winner      Player
	Looser      Player
	WinnerStats StatsByUser
}

type StatsByUser struct {
	User        Player
	Count       int
	Points      int
	BestSeries  int
	TotalDamage int
}
