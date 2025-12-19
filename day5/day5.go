package day5

import (
	"AdventOfCode25/utils"
	"fmt"
	"strconv"
	"strings"

	"github.com/hashicorp/go-set/v3"
)

type Pair struct {
	low, high int
}

func Solution() {

	filaName := "day5/input.txt"
	//filaName := "day5/input.txt"

	lines, err := utils.ReadFile(filaName)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	idsPos := 0

	availableNumbers := set.New[Pair](1000)
	for pos, line := range lines {
		if line == "" {
			idsPos = pos + 1
			break
		}
		pair := parseNumbers(line)
		availableNumbers.Insert(pair)
	}

	total := 0
	for i := idsPos; i < len(lines); i++ {
		total += isFresh(lines[i], availableNumbers)
	}

	println(total)
}

func parseNumbers(line string) Pair {
	numbers := strings.Split(line, "-")

	numberLow, err := strconv.Atoi(numbers[0])
	numberHigh, err2 := strconv.Atoi(numbers[1])
	if err != nil || err2 != nil {
		return Pair{-1, -1}
	}
	return Pair{numberLow, numberHigh}
}

func isFresh(id string, availableNumbers *set.Set[Pair]) int {

	number, err := strconv.Atoi(id)
	if err != nil {
		return 0
	}
	for rangeFresh := range availableNumbers.Items() {
		if number >= rangeFresh.low && number <= rangeFresh.high {
			return 1
		}
	}
	return 0
}
