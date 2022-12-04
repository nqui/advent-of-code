package main

import "fmt"

// A = ROCK
// B = PAPER
// C = SCISSOR

// X = ROCK = 1
// Y = PAPER = 2
// Z = SCISSOR = 3

// X = LOSE = 0
// Y = DRAW = 3
// Z = WIN = 6

const (
	Rock    = 1
	Paper   = 2
	Scissor = 3
	Win     = 6
	Draw    = 3
	Lose    = 0
)

func processDay2() string {
	input := getInput("2")

	resultMap1 := map[string]int{
		"A X": Draw + Rock,
		"A Y": Win + Paper,
		"A Z": Lose + Scissor,
		"B X": Lose + Rock,
		"B Y": Draw + Paper,
		"B Z": Win + Scissor,
		"C X": Win + Rock,
		"C Y": Lose + Paper,
		"C Z": Draw + Scissor,
	}

	resultMap2 := map[string]int{
		"A X": Lose + Scissor,
		"A Y": Draw + Rock,
		"A Z": Win + Paper,
		"B X": Lose + Rock,
		"B Y": Draw + Paper,
		"B Z": Win + Scissor,
		"C X": Lose + Paper,
		"C Y": Draw + Scissor,
		"C Z": Win + Rock,
	}

	score1 := 0
	score2 := 0
	for _, n := range input {
		score1 = score1 + resultMap1[n]
		score2 = score2 + resultMap2[n]
	}

	return fmt.Sprintf("day 2:\npart 1: %d; part 2: %d", score1, score2)
}
