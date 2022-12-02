package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var charsToSigns = map[rune]int{
	'X': ROCK,
	'Y': PAPER,
	'Z': SCISSORS,
	'A': ROCK,
	'B': PAPER,
	'C': SCISSORS,
}

const (
	LOSE     = 0
	DRAW     = 3
	WIN      = 6
	ROCK     = 1
	PAPER    = 2
	SCISSORS = 3
)

func rockPaperScissorsPart1() {
	f, err := os.Open("inputs/input2.txt")
	if err != nil {
		log.Fatalln("Unable to open file", err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	totalscore := 0
	for scanner.Scan() {
		player, oponent := ' ', ' '
		fmt.Sscanf(scanner.Text(), "%c %c\n", &oponent, &player)
		totalscore += checkWin(charsToSigns[oponent], charsToSigns[player])
	}
	fmt.Printf("Total score in RPS: %d\n", totalscore)
}

func checkWin(oponent, player int) int {
	switch {
	case oponent == player:
		//fmt.Printf("%d %d DRAW\n", oponent, player)
		return player + DRAW
	case (player == ROCK && oponent == SCISSORS) || (player == PAPER && oponent == ROCK) || (player == SCISSORS && oponent == PAPER):
		//fmt.Printf("%d %d WIN\n", oponent, player)
		return player + WIN
	default:
		//fmt.Printf("%d %d LOSE\n", oponent, player)
		return player + LOSE
	}
}

var rune2outcome = map[rune]int{
	'X': LOSE,
	'Y': DRAW,
	'Z': WIN,
}

func rockPaperScissorsPart2() {
	f, err := os.Open("inputs/input2.txt")
	if err != nil {
		log.Fatalln("Unable to open file", err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	totalscore := 0
	for scanner.Scan() {
		oponent, outcome := ' ', ' '
		fmt.Sscanf(scanner.Text(), "%c %c\n", &oponent, &outcome)
		totalscore += findwinbyoutcome(rune2outcome[outcome], charsToSigns[oponent])
	}
	fmt.Printf("Total score in RPS by known outcome: %d\n", totalscore)
}

func findwinbyoutcome(outcome, oponent int) int {
	sign := 0
	switch outcome {
	case LOSE:
		switch oponent {
		case ROCK:
			sign = SCISSORS
		case PAPER:
			sign = ROCK
		case SCISSORS:
			sign = PAPER
		}
	case DRAW:
		sign = oponent
	case WIN:
		switch oponent {
		case ROCK:
			sign = PAPER
		case PAPER:
			sign = SCISSORS
		case SCISSORS:
			sign = ROCK
		}
	}
	return sign + outcome
}

func main() {
	rockPaperScissorsPart1()
	rockPaperScissorsPart2()
}
