package day2

import (
	"AdventOfCode25/utils"
	"fmt"
	"strconv"
	"strings"
)

func Solution() {
	fmt.Println("Day2 Exercise:")

	//filaName := "day2/test.txt"
	filaName := "day2/instructions.txt"

	fileLines, err := utils.ReadFile(filaName)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	inputs := strings.Split(fileLines[0], ",")

	result := 0
	for _, input := range inputs {
		invalidIdsSum, err := sumInvalidIds(input)
		if err != nil {
			fmt.Println("Error detecting invalid ids: %v \n", err)
			return
		}
		result += invalidIdsSum
	}

	fmt.Println(result)

}

func sumInvalidIds(input string) (int, error) {
	edges := strings.Split(input, "-")

	start, err := strconv.Atoi(edges[0])
	if err != nil {
		return 0, err
	}

	end, err := strconv.Atoi(edges[1])
	if err != nil {
		return 0, err
	}

	sumInvalidIds := 0
	for i := start; i <= end; i++ {
		if isIdInvalidSecondPart(i) {
			sumInvalidIds += i
		}
	}

	return sumInvalidIds, nil
}

func isIdInvalid(i int) bool {
	str := fmt.Sprintf("%d", i)
	l := len(str)
	if l%2 == 0 && str[0:l/2] == str[l/2:] {
		return true
	}
	return false
}

func isIdInvalidSecondPart(i int) bool {
	str := fmt.Sprintf("%d", i)
	l := len(str)
	for chunkSize := 1; chunkSize <= l/2; chunkSize++ {
		chunkToCompare := str[0:chunkSize]
		if checkChunks(str, l, chunkSize, chunkToCompare) {
			return true
		}
	}
	return false
}

func checkChunks(word string, wordLength int, chunkSize int, chunkToCompare string) bool {
	if wordLength%chunkSize != 0 {
		return false
	}

	for pointer := chunkSize; pointer+chunkSize <= wordLength; pointer = pointer + chunkSize {
		if word[pointer:pointer+chunkSize] != chunkToCompare {
			return false
		}
	}

	println(word)
	return true
}
