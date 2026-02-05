package day8

import (
	"AdventOfCode25/utils"
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

func Solution2() {
	fmt.Println("Day8 Exercise:")

	filaName := "day8/input.txt"
	//filaName := "day8/test.txt"

	lines, err := utils.ReadFile(filaName)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	var points []Point

	for i, line := range lines {
		coordinates := strings.Split(line, ",")
		coordinateX, _ := strconv.ParseFloat(coordinates[0], 64)
		coordinateY, _ := strconv.ParseFloat(coordinates[1], 64)
		coordinateZ, _ := strconv.ParseFloat(coordinates[2], 64)
		box := Point{i, coordinateX, coordinateY, coordinateZ}
		points = append(points, box)
	}

	var edges []Edge
	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			p1, p2 := points[i], points[j]
			dist := math.Sqrt(math.Pow(p1.x-p2.x, 2) + math.Pow(p1.y-p2.y, 2) + math.Pow(p1.z-p2.z, 2))
			edges = append(edges, Edge{p1.id, p2.id, dist})
		}
	}

	sort.Slice(edges, func(i, j int) bool {
		return edges[i].distance < edges[j].distance
	})

	dsu := NewDSU(len(points))
	numComponents := len(points)
	var lastX1, lastX2 float64

	for _, edge := range edges {
		rootP1 := dsu.Find(edge.p1)
		rootP2 := dsu.Find(edge.p2)

		if rootP1 != rootP2 {
			dsu.Union(edge.p1, edge.p2)
			numComponents--

			if numComponents == 1 {
				lastX1 = points[edge.p1].x
				lastX2 = points[edge.p2].x
				break
			}
		}
	}

	result := int64(lastX1) * int64(lastX2)
	fmt.Printf("La conexión final fue entre X=%.0f y X=%.0f\n", lastX1, lastX2)
	fmt.Printf("Resultado Parte 2: %d\n", result)
}
