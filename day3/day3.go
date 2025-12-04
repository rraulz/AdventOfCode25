package day3

import (
	"AdventOfCode25/utils"
	"fmt"
	"strconv"
)

func Solution() {

	//filaName := "day3/test.txt"
	filaName := "day3/input.txt"

	lines, err := utils.ReadFile(filaName)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	total := 0
	for _, line := range lines {
		total += findBiggestJoltage(line)
	}

	println(total)
}

func findBiggestJoltage(line string) int {
	biggestJoltage := 0

	leaderNumber := 0

	for _, char := range line {
		number, _ := strconv.Atoi(string(char))

		joltage := (leaderNumber * 10) + number
		if joltage > biggestJoltage {
			biggestJoltage = joltage
		}

		if number > leaderNumber {
			leaderNumber = number
		}

	}

	return biggestJoltage
}
