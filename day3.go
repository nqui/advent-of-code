package main

import (
	"fmt"
	"golang.org/x/exp/slices"
)

func processDay3() string {
	input := getInput("3")

	priorityMap := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
		"d": 4,
		"e": 5,
		"f": 6,
		"g": 7,
		"h": 8,
		"i": 9,
		"j": 10,
		"k": 11,
		"l": 12,
		"m": 13,
		"n": 14,
		"o": 15,
		"p": 16,
		"q": 17,
		"r": 18,
		"s": 19,
		"t": 20,
		"u": 21,
		"v": 22,
		"w": 23,
		"x": 24,
		"y": 25,
		"z": 26,
		"A": 27,
		"B": 28,
		"C": 29,
		"D": 30,
		"E": 31,
		"F": 32,
		"G": 33,
		"H": 34,
		"I": 35,
		"J": 36,
		"K": 37,
		"L": 38,
		"M": 39,
		"N": 40,
		"O": 41,
		"P": 42,
		"Q": 43,
		"R": 44,
		"S": 45,
		"T": 46,
		"U": 47,
		"V": 48,
		"W": 49,
		"X": 50,
		"Y": 51,
		"Z": 52,
	}
	sum1 := 0
	for _, line := range input {
		fHalf := []byte(line[:len(line)/2])
		sHalf := []byte(line[len(line)/2:])
		for _, n := range fHalf {
			if slices.Contains(sHalf, n) {
				sum1 = sum1 + priorityMap[string(n)]
				break
			}
		}
	}

	sum2 := 0
	i := 0
	for i < len(input) {
		group := input[i : i+3]
		line1 := []byte(group[0])
		line2 := []byte(group[1])
		line3 := []byte(group[2])
		for _, n := range line1 {
			if slices.Contains(line2, n) && slices.Contains(line3, n) {
				sum2 = sum2 + priorityMap[string(n)]
				break
			}
		}
		i = i + 3
	}

	return fmt.Sprintf("day 3:\npart 1: %d; part 2: %d", sum1, sum2)
}
