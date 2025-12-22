package day5

import (
	"AdventOfCode25/utils"
	"fmt"
	"strconv"
	"strings"

	"github.com/hashicorp/go-set/v3"
)

func Solution2() {

	//filaName := "day5/test.txt"
	filaName := "day5/input.txt"

	lines, err := utils.ReadFile(filaName)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	availableNumbers := set.New[Pair](1000)
	for _, line := range lines {
		if line == "" {
			break
		}

		updatedExistingNumbers := numbersFresh(line, availableNumbers)
		availableNumbers = updatedExistingNumbers
	}

	total := 0
	for pairFresh := range availableNumbers.Items() {
		total += pairFresh.high - pairFresh.low + 1
	}
	println(total)
}

func numbersFresh(line string, numberRangeParsed *set.Set[Pair]) *set.Set[Pair] {
	numbers := strings.Split(line, "-")
	numberLow, err := strconv.Atoi(numbers[0])
	numberHigh, err2 := strconv.Atoi(numbers[1])
	if err != nil || err2 != nil {
		return numberRangeParsed
	}

	newFreshNumber := Pair{numberLow, numberHigh}
	availableNumbers := set.New[Pair](1000)

	for rangeParsed := range numberRangeParsed.Items() {
		if newFreshNumber.Contains(rangeParsed) {
			//fmt.Printf("%s ->  deleted by %s\n", rangeParsed, newFreshNumber)
			continue //Existing gets eliminated
		}
		if rangeParsed.Contains(newFreshNumber) {
			//fmt.Printf("%s ->  contained\n", newFreshNumber)
			return numberRangeParsed //Its already contained
		}
		if newFreshNumber.OverlapsHigh(rangeParsed) {
			availableNumbers.Insert(Pair{rangeParsed.low, newFreshNumber.low - 1})
			//fmt.Printf("%s ->  Overlapped High by %s new = %s\n", rangeParsed, newFreshNumber, Pair{rangeParsed.low, newFreshNumber.low - 1})
			continue //Inserted missing part
		}
		if newFreshNumber.OverlapsLow(rangeParsed) {
			availableNumbers.Insert(Pair{newFreshNumber.high + 1, rangeParsed.high})
			//fmt.Printf("%s ->  Overlapped low by %s new = %s\n", rangeParsed, newFreshNumber, Pair{newFreshNumber.high + 1, rangeParsed.high})
			continue //Inserted missing part
		}
		availableNumbers.Insert(rangeParsed)
	}

	availableNumbers.Insert(newFreshNumber)
	return availableNumbers
}
