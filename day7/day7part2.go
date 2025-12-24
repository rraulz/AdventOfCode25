package day7

import (
	"AdventOfCode25/utils"
	"fmt"
)

func Solution2() {
	filaName := "day7/input.txt"
	//filaName := "day7/test.txt"

	lines, err := utils.ReadFileInMatrix(filaName)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	paths := executeWorlds(lines)
	fmt.Print(paths)
}

func executeWorlds(input [][]string) int {
	for col, char := range input[0] {
		if char == "S" {
			return StartTimeLine(input, 0, col)
		}
	}
	return 0
}

var cache = make(map[string]int)

func StartTimeLineCached(input [][]string, row, col int) int {
	key := fmt.Sprintf("%d,%d", row, col)
	//fmt.Printf("Cache? %d-%d=%s\n", row, col, input[row][col])
	if val, found := cache[key]; found {
		//fmt.Printf("Cache HIT %d-%d=%s\n", row, col, input[row][col])
		return val
	}

	res := StartTimeLine(input, row, col)
	//fmt.Printf("Cache MISS %d-%d=%s SAVING RES: %d\n", row, col, input[row][col], res)
	cache[key] = res
	return res
}

func StartTimeLine(input [][]string, row, col int) int {
	for r := row + 1; r < len(input); r++ {
		//fmt.Printf("Checking %d-%d=%s\n", r, col, input[r][col])
		if input[r][col] == "^" {
			return StartTimeLineCached(input, r, col-1) + StartTimeLineCached(input, r, col+1)
		}
	}
	return 1
}
