package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func getInput(day string) []string {
	dat, err := os.ReadFile(fmt.Sprintf("/Users/nqu/code/advent-of-code/day%s.txt", day))
	if err != nil {
		log.Fatal(err)
	}

	return strings.Split(string(dat), "\n")
}

func main() {
	fmt.Println(processDay1())
	fmt.Println(processDay2())
	fmt.Println(processDay3())
	fmt.Println(processDay4())
}
