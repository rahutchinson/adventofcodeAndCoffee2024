package main

import (
	"bufio"
	"fmt"
	"os"
)

type Pair struct {
	X int
	Y int
}

func main() {
	// Open the file
	file, err := os.Open("day10input.txt")
	// file, err := os.Open("day10ex.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var data []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		data = append(data, line) // Append each line to the slice
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	fmt.Println(data)

	paths := make([][]int, 0)
	for i := 0; i < len(data); i++ {
		line := []int{}
		for j := 0; j < len(data[i]); j++ {
			line = append(line, int(data[i][j]-'0'))
		}
		paths = append(paths, line)
	}

	totals := []int{}
	for x, val := range paths {
		for y, val2 := range val {
			allPaths := map[Pair]int{}
			if val2 == 0 {
				findPath(paths, x, y, &allPaths)
			}
			if len(allPaths) > 0 {
				totals = append(totals, len(allPaths))
			}
		}
	}

	fmt.Println(sumInts(totals))
	//Part 2
	total2 := []int{}
	for x, val := range paths {
		for y, val2 := range val {
			allPaths := map[Pair]int{}
			if val2 == 0 {
				findPath(paths, x, y, &allPaths)
			}
			if len(allPaths) > 0 {
				total := 0
				for _, val := range allPaths {
					total += val
				}
				total2 = append(total2, total)
			}
		}
	}
	fmt.Println(total2)
	fmt.Println(sumInts(total2))
}

// Function to calculate the sum of a list of integers
func sumInts(ints []int) int {
	sum := 0
	for _, v := range ints {
		sum += v
	}
	return sum
}

func findPath(woods [][]int, x, y int, maps *map[Pair]int) {
	if woods[x][y] == 9 {
		(*maps)[Pair{x, y}] += 1
		return
	}
	directions := []Pair{Pair{0, 1}, Pair{1, 0}, Pair{0, -1}, Pair{-1, 0}}

	for _, direction := range directions {
		newX := x + direction.X
		newY := y + direction.Y
		if newX >= 0 && newX < len(woods) && newY >= 0 && newY < len(woods[0]) {
			if woods[newX][newY] == woods[x][y]+1 {
				findPath(woods, newX, newY, maps)
			}
		}
	}
}

func findPath2(woods [][]int, x, y int, maps *map[Pair]int) {
	if woods[x][y] == 9 {
		(*maps)[Pair{x, y}] += 1
		return
	}
	directions := []Pair{Pair{0, 1}, Pair{1, 0}, Pair{0, -1}, Pair{-1, 0}}

	for _, direction := range directions {
		newX := x + direction.X
		newY := y + direction.Y
		if newX >= 0 && newX < len(woods) && newY >= 0 && newY < len(woods[0]) {
			if woods[newX][newY] == woods[x][y]+1 {
				findPath(woods, newX, newY, maps)
			}
		}
	}
}
