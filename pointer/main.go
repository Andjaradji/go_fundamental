package main

import "fmt"

type Gamer struct {
	Name  string
	Games []string
}

func (gamer *Gamer) AddGame (game string) {
	gamer.Games = append(gamer.Games, game )
}

func main() {
	gamer := Gamer{Name: "Andjar"}

	gamer.AddGame("TETRIS")
	gamer.AddGame("MONOPOLY")
	gamer.AddGame("FIFA 2020")
	gamer.AddGame("COUNTER STRIKE")


	for _, game := range gamer.Games {
		fmt.Println(game)
	}
}
