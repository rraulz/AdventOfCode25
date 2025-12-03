package day1

import (
	"fmt"
	"strconv"
)

const (
	startLockPosition = 50
	lockModulo        = 100
)

func Solution() {
	fmt.Println("Day1 Exercise:")

	filaName := "day1/instructions.txt"

	fileLines, err := ReadFile(filaName)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	code, err := processLockInstructions(fileLines)
	if err != nil {
		fmt.Printf("Error processing instructions: %v\n", err)
		return
	}

	fmt.Printf("Solution is: %d\n", code)

	code_0x434C49434B, err := processLockInstructions_0x434C49434B(fileLines)
	if err != nil {
		fmt.Printf("Error processing instructions: %v\n", err)
		return
	}

	fmt.Printf("Solution 0x434C49434B is: %d\n", code_0x434C49434B)
}

func processLockInstructions(fileLines []string) (int, error) {
	zeros := 0
	lockPosition := startLockPosition

	for _, line := range fileLines {

		side, value, err := parseLine(line)
		if err != nil {
			return 0, fmt.Errorf("parse error in line '%s': %w", line, err)
		}

		switch side {
		case Left:
			lockPosition -= value
		case Right:
			lockPosition += value
		}

		lockPosition = ((lockPosition % lockModulo) + lockModulo) % lockModulo
		if lockPosition == 0 {
			zeros++
		}
	}
	return zeros, nil
}

func processLockInstructions_0x434C49434B(fileLines []string) (int, error) {
	zeros := 0
	lockPosition := startLockPosition
	for _, line := range fileLines {
		side, value, err := parseLine(line)
		if err != nil {
			return 0, fmt.Errorf("parse error in line '%s': %w", line, err)
		}

		remaining := value % 100
		zeros += value / 100

		if side == Left {
			if (lockPosition != 0 && lockPosition-remaining < 0) || lockPosition-remaining == 0 {
				zeros++
			}
			lockPosition = (lockPosition + (lockModulo - remaining)) % lockModulo
		} else {
			if lockPosition+remaining > lockModulo-1 {
				zeros++
			}
			lockPosition = (lockPosition + remaining) % lockModulo
		}
	}
	return zeros, nil
}

func parseLine(line string) (Side, int, error) {
	if len(line) < 2 {
		return Left, 0, fmt.Errorf("invalid instruction")
	}

	var side Side
	switch line[0] {
	case 'L', 'l':
		side = Left
	case 'R', 'r':
		side = Right
	default:
		return "", 0, fmt.Errorf("unknown side '%c'", line[0])
	}

	value, err := strconv.Atoi(line[1:])
	if err != nil {
		return "", 0, fmt.Errorf("invalid number: %w", err)
	}

	return side, value, nil
}
