package usecases

import (
	"fmt"
	"testing"
)

func TestPlayRound(t *testing.T) {
	for i := 0; i < 3; i++ {
		round := PlayRound(i)

		switch i {
		case ROCK:
			fmt.Println("El jugador elijio PIEDRA.")
		case PAPER:
			fmt.Println("El jugador elijio PAPEL.")
		case SCISSORS:
			fmt.Println("El jugador elijio TIJERAS.")
		}

		fmt.Printf("Message: %s\n", round.Message)
		fmt.Printf("ComputerChoice: %s\n", round.ComputerChoice)
		fmt.Printf("RoundResult: %s\n", round.RoundResult)
		fmt.Printf("ComputerChoiceInt: %d\n", round.ComputerChoiceInt)
		fmt.Printf("ComputerScore: %s\n", round.ComputerScore)
		fmt.Printf("PlayerScore: %s\n", round.PlayerScore)
		fmt.Println("\n========================================")
	}

}
