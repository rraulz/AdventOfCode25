package day4

import (
	"AdventOfCode25/utils"
	"fmt"
)

//Should be improved, very slow

func Solution() {
	filaName := "day4/input.txt"

	lines, err := utils.ReadFileInMatrix(filaName)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	total := 0
	for {
		rollsRemoved := 0
		for pos, line := range lines {
			var prev, next []string

			if pos > 0 {
				prev = lines[pos-1]
			}
			if pos < len(lines)-1 {
				next = lines[pos+1]
			}

			//rollsRemoved += papersAvailable(line, prev, next)
			rollsSubtractedInLine, cleanedLine := papersAvailableRemoving(line, prev, next)
			rollsRemoved += rollsSubtractedInLine
			lines[pos] = cleanedLine
		}

		if rollsRemoved == 0 {
			println(total)
			return
		}
		total += rollsRemoved
	}
}

func papersAvailable(line []string, linePrev []string, lineNext []string) int {
	total := 0

	for pos, object := range line {
		if object != "@" {
			continue
		}
		if countRollsLine(linePrev, pos)+(countRollsLine(line, pos)-1)+countRollsLine(lineNext, pos) < 4 {
			total++
		}
	}

	return total
}

func papersAvailableRemoving(line []string, linePrev []string, lineNext []string) (int, []string) {
	total := 0

	for pos, object := range line {
		if object != "@" {
			continue
		}
		if countRollsLine(linePrev, pos)+(countRollsLine(line, pos)-1)+countRollsLine(lineNext, pos) < 4 {
			line[pos] = "."
			total++
		}
	}

	return total, line
}

func countRollsLine(line []string, pos int) int {
	if line == nil {
		return 0
	}

	totalPoints := 0
	if pos != 0 {
		if line[pos-1] == "@" {
			totalPoints++
		}
	}
	if pos < len(line)-1 {
		if line[pos+1] == "@" {
			totalPoints++
		}
	}
	if line[pos] == "@" {
		totalPoints++
	}
	return totalPoints
}
