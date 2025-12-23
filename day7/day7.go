package day7

import (
	"AdventOfCode25/utils"
	"fmt"
)

func Solution() {
	filaName := "day7/input.txt"
	//filaName := "day7/test.txt"

	lines, err := utils.ReadFileInMatrix(filaName)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	total := calculateSplits(lines)
	fmt.Print(total)
}

func calculateSplits(lines [][]string) int {
	counter := 0
	for i := 1; i < len(lines); i++ {
		for j := 0; j < len(lines[i]); j++ {
			topChar := lines[i-1][j]
			actualChar := lines[i][j]

			if topChar == "S" || (topChar == "|" && actualChar != "^") {
				lines[i][j] = "|"
				continue
			}
			if actualChar == "^" && topChar == "|" {
				counter++
				lines[i][j+1] = "|"
				lines[i][j-1] = "|"
			}
		}
	}
	return counter
}
