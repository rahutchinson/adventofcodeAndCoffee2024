package main

import (
	"bufio"
	"fmt"
	"os"
)

type complexNumber complex128

type Grid map[complexNumber]rune

func readInput(filePath string) []string {
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return lines
}

func createGrid(lines []string) Grid {
	grid := make(Grid)
	for y, line := range lines {
		for x, c := range line {
			grid[complexNumber(complex(float64(x), float64(y)))] = c
		}
	}
	return grid
}

func floodFill(grid Grid, start complexNumber) (set map[complexNumber]struct{}) {
	set = make(map[complexNumber]struct{})
	queue := []complexNumber{start}
	symbol := grid[start]
	set[start] = struct{}{}

	directions := []complexNumber{1, -1, 1i, -1i}
	for len(queue) > 0 {
		pos := queue[0]
		queue = queue[1:]
		for _, d := range directions {
			newPos := pos + d
			if _, exists := grid[newPos]; exists {
				if _, covered := set[newPos]; !covered && grid[newPos] == symbol {
					set[newPos] = struct{}{}
					queue = append(queue, newPos)
				}
			}
		}
	}
	return
}

func getArea(region map[complexNumber]struct{}) int {
	return len(region)
}

func getPerimeter(region map[complexNumber]struct{}) int {
	directions := []complexNumber{1, -1, 1i, -1i}
	perimeter := 0

	for pos := range region {
		for _, d := range directions {
			newPos := pos + d
			if _, exists := region[newPos]; !exists {
				perimeter++
			}
		}
	}
	return perimeter
}

func getSidesCount(region map[complexNumber]struct{}) int {
	perimeterObjects := make(map[complexNumber]complexNumber)
	directions := []complexNumber{1, -1, 1i, -1i}
	for pos := range region {
		for _, d := range directions {
			newPos := pos + d
			if _, exists := region[newPos]; !exists {
				perimeterObjects[newPos] = d
			}
		}
	}

	distinctSides := 0
	visited := make(map[complexNumber]struct{})
	for pos, d := range perimeterObjects {
		if _, alreadyVisited := visited[pos]; alreadyVisited {
			continue
		}
		distinctSides++
		current := pos
		for {
			visited[current] = struct{}{}
			next := current + d*1i
			if nextDir, ok := perimeterObjects[next]; ok && nextDir == d {
				current = next
			} else {
				break
			}
		}
		current = pos
		for {
			visited[current] = struct{}{}
			next := current + d*-1i
			if nextDir, ok := perimeterObjects[next]; ok && nextDir == d {
				current = next
			} else {
				break
			}
		}
	}
	return distinctSides
}

func main() {
	// lines := readInput("day12input.txt")
	lines := readInput("day12ex.txt")
	grid := createGrid(lines)

	regions := []struct {
		symbol rune
		region map[complexNumber]struct{}
	}{}

	uncovered := make(map[complexNumber]struct{})
	for pos := range grid {
		uncovered[pos] = struct{}{}
	}

	for len(uncovered) > 0 {
		for start := range uncovered {
			region := floodFill(grid, start)
			for pos := range region {
				delete(uncovered, pos)
			}
			regions = append(regions, struct {
				symbol rune
				region map[complexNumber]struct{}
			}{grid[start], region})
			break
		}
	}

	// Part 1
	price := 0
	for _, reg := range regions {
		area := getArea(reg.region)
		perimeter := getPerimeter(reg.region)
		price += area * perimeter
	}
	fmt.Println("Part 1:", price)

	// Part 2
	price = 0
	for _, reg := range regions {
		area := getArea(reg.region)
		sides := getSidesCount(reg.region)
		price += area * sides
	}
	fmt.Println("Part 2:", price)
}
