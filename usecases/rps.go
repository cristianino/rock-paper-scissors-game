package usecases

import (
	"math/rand"
	"rpsweb/models"
	"strconv"
)

const (
	ROCK     = 0 //Piedra. Vence tijeras. (tijeras + 1) % 3 = 0
	PAPER    = 1 //Papel. Vence piedra. (piedra + 1) % 3 = 1
	SCISSORS = 2 //Tijeras. Vence papel. (papel + 1) % 3 = 2
)

var winMessages = []string{
	"¡Bien hecho!",
	"¡Buen trabajo!",
	"Deberías comprar un boleto de lotería",
}

var loseMessages = []string{
	"¡Que lástima!",
	"¡Inténtalo de nuevo!",
	"Hoy simplemente no es tu día.",
}

var drawMessages = []string{
	"Las grandes mentes piensan igual.",
	"Oh Oh. Inténtalo de nuevo.",
	"Nadie gana, pero puedes intentarlo de nuevo.",
}

var ComputerScore, PlayerScore int

func PlayRound(playerValue int) models.Round {
	computerValue := rand.Intn(3)
	var computerChoice, roundResult string
	var computerChoiceInt int

	switch computerValue {
	case ROCK:
		computerChoiceInt = ROCK
		computerChoice = "La computadora eligió PIEDRA."

	case PAPER:
		computerChoiceInt = PAPER
		computerChoice = "La computadora eligió PAPEL."

	case SCISSORS:
		computerChoiceInt = SCISSORS
		computerChoice = "La computadora eligió TIJERA."
	}

	messageInt := rand.Intn(len(winMessages))

	var message string

	if playerValue == computerValue {
		roundResult = "Es un empate"
		message = drawMessages[messageInt]
	} else if playerValue == (computerValue+1)%3 {
		PlayerScore++
		roundResult = "El jugador gana"
		message = winMessages[messageInt]
	} else {
		ComputerScore++
		roundResult = "La computadora gana"
		message = loseMessages[messageInt]
	}

	return models.Round{
		Message:           message,
		ComputerChoice:    computerChoice,
		RoundResult:       roundResult,
		ComputerChoiceInt: computerChoiceInt,
		ComputerScore:     strconv.Itoa(ComputerScore),
		PlayerScore:       strconv.Itoa(PlayerScore),
	}
}
