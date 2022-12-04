package main

import (
	"fmt"
	"golang.org/x/exp/slices"
	"log"
	"strconv"
	"strings"
)

func processDay4() string {
	input := getInput("4")

	nums := make([]int, 100)
	for i, _ := range nums {
		nums[i] = i
	}

	pairs := 0
	overlaps := 0
	for _, line := range input {
		firstSet := strings.Split(strings.Split(line, ",")[0], "-")
		firstIndex, err := strconv.Atoi(firstSet[0])
		if err != nil {
			log.Fatal(err)
		}
		secondIndex, err := strconv.Atoi(firstSet[1])
		if err != nil {
			log.Fatal(err)
		}
		firstSetNums := nums[firstIndex : secondIndex+1]

		secondSet := strings.Split(strings.Split(line, ",")[1], "-")
		firstIndex, err = strconv.Atoi(secondSet[0])
		if err != nil {
			log.Fatal(err)
		}
		secondIndex, err = strconv.Atoi(secondSet[1])
		if err != nil {
			log.Fatal(err)
		}
		secondSetNums := nums[firstIndex : secondIndex+1]

		overlap := false
		contained := true
		for _, n := range firstSetNums {
			if !slices.Contains(secondSetNums, n) {
				contained = false
				break
			} else {
				if !overlap {
					overlaps++
					overlap = true
				}
			}
		}

		if contained {
			pairs++
		} else {
			contained = true
			for _, n := range secondSetNums {
				if !slices.Contains(firstSetNums, n) {
					contained = false
					break
				} else {
					if !overlap {
						overlaps++
						overlap = true
					}
				}
			}
			if contained {
				pairs++
			}
		}
	}
	return fmt.Sprintf("day 4:\npart 1: %d; part 2: %d", pairs, overlaps)
}
