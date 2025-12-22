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

func (p Pair) Contains(other Pair) bool {
	if (other.high <= p.high) && (other.low >= p.low) {
		return true
	}
	return false
}

// P (25-35) - Other (20-30) = TRUE
func (p Pair) OverlapsHigh(other Pair) bool {
	if (other.high <= p.high) && (other.high >= p.low) && (other.low <= p.low) {
		return true
	}
	return false
}

// P (20-30) - Other (25-35) = TRUE
func (p Pair) OverlapsLow(other Pair) bool {
	if (other.low >= p.low) && (other.low <= p.high) && (other.high >= p.low) {
		return true
	}
	return false
}

func (p Pair) String() string {
	return fmt.Sprintf("%d-%d", p.low, p.high)
}

func Solution() {

	filaName := "day5/input.txt"

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
