package day8

import (
	"AdventOfCode25/utils"
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

type Point struct {
	id      int
	x, y, z float64
}

type Edge struct {
	p1, p2   int
	distance float64
}

type DSU struct {
	parent []int
	size   []int
}

func NewDSU(n int) *DSU {
	parent := make([]int, n)
	size := make([]int, n)
	for i := range parent {
		parent[i] = i
		size[i] = 1
	}
	return &DSU{parent, size}
}

func (d *DSU) Find(i int) int {
	if d.parent[i] == i {
		return i
	}
	d.parent[i] = d.Find(d.parent[i])
	return d.parent[i]
}

func (d *DSU) Union(i, j int) {
	rootI := d.Find(i)
	rootJ := d.Find(j)
	if rootI != rootJ {
		if d.size[rootI] < d.size[rootJ] {
			rootI, rootJ = rootJ, rootI
		}
		d.parent[rootJ] = rootI
		d.size[rootI] += d.size[rootJ]
	}
}

func Solution() {
	fmt.Println("Day8 Exercise:")

	filaName := "day8/input.txt"
	//filaName := "day8/test.txt"
	numConnections := 1000

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

	for i := 0; i < numConnections; i++ {
		dsu.Union(edges[i].p1, edges[i].p2)
	}

	circuitSizes := []int{}
	visited := make(map[int]bool)
	for i := 0; i < len(points); i++ {
		root := dsu.Find(i)
		if !visited[root] {
			circuitSizes = append(circuitSizes, dsu.size[root])
			visited[root] = true
		}
	}

	sort.Sort(sort.Reverse(sort.IntSlice(circuitSizes)))

	result := 1
	for i := 0; i < 3 && i < len(circuitSizes); i++ {
		result *= circuitSizes[i]
	}

	fmt.Printf("Resultado: %d\n", result)
}
