package day9

import (
	"AdventOfCode25/utils"
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

type Edge struct {
	x, y int
}

func Solution() {
	fmt.Println("Day9 Exercise:")

	filaName := "day9/input.txt"
	//filaName := "day9/test.txt"

	lines, err := utils.ReadFile(filaName)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	var points []Edge
	for _, line := range lines {
		coordinates := strings.Split(line, ",")
		coordinateX, _ := strconv.Atoi(coordinates[0])
		coordinateY, _ := strconv.Atoi(coordinates[1])
		red := Edge{coordinateX, coordinateY}
		points = append(points, red)
	}

	sort.Slice(points, func(i, j int) bool {
		return points[i].y < points[j].y
	})

	largestArea := float64(0)
	for i := 0; i < len(points)-1; i++ {
		for j := i + 1; j < len(points); j++ {
			a := points[i]
			b := points[j]

			ladoA := math.Abs(float64(a.x-b.x)) + 1
			ladoB := math.Abs(float64(a.y-b.y)) + 1
			area := ladoA * ladoB
			if area > largestArea {
				largestArea = area
			}

		}
	}

	fmt.Printf("%.2f", largestArea)
}
