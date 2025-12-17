package day3

import (
	"AdventOfCode25/utils"
	"fmt"
	"math"
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
		total += findBiggestJoltagePart2(line)
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

func findBiggestJoltagePart2(line string) int {
	const joltageSize = 12
	joltage := make([]int, joltageSize)
	cleanArray(joltage, 0)

	lineSize := len(line)

	for joltageBatteryPos, char := range line {
		number, _ := strconv.Atoi(string(char))

		for existingBatteryPos, battery := range joltage {
			if (lineSize - joltageBatteryPos) < (joltageSize - existingBatteryPos) {
				continue
			}

			if number > battery {
				joltage[existingBatteryPos] = number
				cleanArray(joltage, existingBatteryPos+1)
				break
			}
		}
	}

	biggestJoltage := 0.0
	for pos, battery := range joltage {
		biggestJoltage += float64(battery) * math.Pow(float64(10), float64(joltageSize-1-pos))
	}
	fmt.Printf("Input: %s -> %f\n", line, biggestJoltage)
	return int(biggestJoltage)
}

func cleanArray(array []int, startingPos int) {
	for i := startingPos; i < len(array); i++ {
		array[i] = 0
	}
}
