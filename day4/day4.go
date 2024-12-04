package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// Open the file
	// file, err := os.Open("day4ex.txt")
	file, err := os.Open("day4input.txt")
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
	count := 0
	for i, row := range data {
		for j := 0; j < len(row); j++ {
			if data[i][j] == 'X' {
				if j >= 3 && checkDirection(i, j, "up", data) {
					count += 1
				}
				if j+3 <= len(data)-1 && checkDirection(i, j, "down", data) {
					count += 1
				}
				if i >= 3 && checkDirection(i, j, "back", data) {
					count += 1
				}
				if i+3 <= len(row)-1 && checkDirection(i, j, "forward", data) {
					count += 1
				}
				if i >= 3 && j >= 3 && checkDirection(i, j, "upleft", data) {
					count += 1
				}
				if i >= 3 && j+3 <= len(data)-1 && checkDirection(i, j, "upright", data) {
					count += 1
				}
				if i+3 <= len(row)-1 && j >= 3 && checkDirection(i, j, "downleft", data) {
					count += 1
				}
				if i+3 <= len(row)-1 && j+3 <= len(data)-1 && checkDirection(i, j, "downright", data) {
					count += 1
				}
			}
		}
	}
	fmt.Println("Number of XMAS found:", count)
}

// the value at x,y should always be the X in XMAS
func checkDirection(x int, y int, direction string, arr []string) bool {
	if direction == "up" {
		if arr[x][y-1] == 'M' && arr[x][y-2] == 'A' && arr[x][y-3] == 'S' {
			return true
		}
	} else if direction == "down" {
		if arr[x][y+1] == 'M' && arr[x][y+2] == 'A' && arr[x][y+3] == 'S' {
			return true
		}
	} else if direction == "back" {
		if arr[x-1][y] == 'M' && arr[x-2][y] == 'A' && arr[x-3][y] == 'S' {
			return true
		}
	} else if direction == "forward" {
		if arr[x+1][y] == 'M' && arr[x+2][y] == 'A' && arr[x+3][y] == 'S' {
			return true
		}
	} else if direction == "upleft" {
		if arr[x-1][y-1] == 'M' && arr[x-2][y-2] == 'A' && arr[x-3][y-3] == 'S' {
			return true
		}
	} else if direction == "upright" {
		if arr[x-1][y+1] == 'M' && arr[x-2][y+2] == 'A' && arr[x-3][y+3] == 'S' {
			return true
		}
	} else if direction == "downleft" {
		if arr[x+1][y-1] == 'M' && arr[x+2][y-2] == 'A' && arr[x+3][y-3] == 'S' {
			return true
		}
	} else if direction == "downright" {
		if arr[x+1][y+1] == 'M' && arr[x+2][y+2] == 'A' && arr[x+3][y+3] == 'S' {
			return true
		}
	}
	return false
}

func toInt(str string) int {
	value, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println("Error converting string to int:", err)
		return 0
	}
	return value
}

func abs(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}

func removeAtIndex(arr []int, index int) []int {
	if index < 0 || index >= len(arr) {
		return arr // Invalid index, return the original array
	}
	return append(arr[:index], arr[index+1:]...)
}
