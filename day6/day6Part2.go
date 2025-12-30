package day6

import (
	"AdventOfCode25/utils"
	"fmt"
	"strings"
)

func Solution2() {
	//filaName := "day6/test.txt"
	filaName := "day6/input.txt"

	lines, err := utils.ReadFile(filaName)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	total := parseLines(lines)

	fmt.Print(total)
}

func parseLines(lines []string) int {
	total := 0
	operatorsLine := lines[len(lines)-1]
	lines = lines[:len(lines)-1]
	operators := strings.Fields(operatorsLine)

	columnProcessing := 0
	columnTotal := 0
	for i := 0; i < len(lines[0]); i++ {
		numberFormed := 0
		spacesFound := 0

		for _, line := range lines {
			if i > len(line)-1 {
				continue
			}
			number := int(line[i] - '0')
			if number == 240 { //Space
				spacesFound++
				continue
			}

			if numberFormed != 0 && number != 0 {
				numberFormed = shiftBase10(numberFormed)
			}
			numberFormed += number
		}

		if spacesFound == len(lines) {
			fmt.Printf("------> %d\n", columnTotal)
			total += columnTotal
			columnProcessing++
			columnTotal = 0
		} else {
			operator := operators[columnProcessing]
			if operator == "*" {
				if numberFormed == 0 {
					continue
				}
				if columnTotal == 0 {
					columnTotal = 1
				}
				fmt.Printf("%d * %d\n", columnTotal, numberFormed)
				columnTotal = columnTotal * numberFormed

			} else {
				fmt.Printf("%d + %d\n", columnTotal, numberFormed)
				columnTotal = columnTotal + numberFormed
			}
		}
	}

	fmt.Printf("------> %d\n", columnTotal)
	total += columnTotal

	return total
}

func shiftBase10(n int) int {
	return n * 10
}
