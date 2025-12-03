package day1

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

const (
	startLockPosition = 50
	lockModulo        = 100
)

func Solution() {
	fmt.Println("Day1 Exercise:")

	filaName := "day1/instructions.txt"
	file, err := os.Open(filaName)
	if err != nil {
		log.Fatalf("Error opening file(path: %s): %v", filaName, err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	code, err := processLockInstructions(scanner)
	if err != nil {
		fmt.Printf("Error processing instructions: %v\n", err)
		return
	}

	fmt.Printf("Solution is: %d\n", code)
}

func processLockInstructions(sc *bufio.Scanner) (int, error) {
	zeros := 0
	lockPosition := startLockPosition

	for sc.Scan() {
		line := sc.Text()
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
