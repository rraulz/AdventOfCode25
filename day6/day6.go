package day6

import (
	"AdventOfCode25/utils"
	"fmt"
	"strconv"
)

func Solution() {
	filaName := "day6/input.txt"

	lines, err := utils.ReadFileInMatrixBySpaces(filaName)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	total := calculate(lines)

	fmt.Print(total)
}

func calculate(allLines [][]string) int {
	total := 0
	lastLine := allLines[len(allLines)-1]
	for i := 0; i < len(lastLine); i++ {
		operator := lastLine[i]
		columnResult := 0
		for j := 0; j < len(allLines)-1; j++ {
			columnResult = operate(columnResult, operator, allLines[j][i])
		}
		total += columnResult
	}
	return total
}

func operate(total int, operator string, input string) int {
	inputNumber, err := strconv.Atoi(input)
	if err != nil {
		return total
	}
	if operator == "*" {
		if total == 0 {
			return inputNumber
		}
		return total * inputNumber
	}
	return total + inputNumber
}
