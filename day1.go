package main

import (
	"fmt"
	"log"
	"sort"
	"strconv"
)

func processDay1() string {
	input := getInput("1")

	sums := make([]int, 0)
	sum := 0
	for _, n := range input {
		if n == "" {
			sums = append(sums, sum)
			sum = 0
			continue
		}
		num, err := strconv.Atoi(n)
		if err != nil {
			log.Fatal(err)
		}
		sum = sum + num
	}
	sort.Ints(sums)
	return fmt.Sprintf("day 1:\npart 1: %d; part 2: %d", sums[len(sums)-1], sums[len(sums)-1]+sums[len(sums)-2]+sums[len(sums)-3])
}
