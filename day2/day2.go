package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Open the file
	// file, err := os.Open("day2ex.txt")
	file, err := os.Open("day2input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var lists [][]int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		var list []int
		for _, part := range strings.Fields(line) {
			if number, err := strconv.Atoi(part); err == nil {
				list = append(list, number)
			}
		}
		lists = append(lists, list)
	}
	sum := 0
	// part 2 is more challenging - if you run into an issue then you need to remove the value that lead to that issue
	// but its more complicated because of the initial check for increase or decrease
	// so for 1 3 2 4 5
	// going to try to just keep moving forward and counting failures when they are more then 1 then return
	// can't do that - need to remove the value that caused the issue not just count
	// but you also need to remove the right value not just the first value that has an issue
	// so if you have an issue then you need to create the list with
	for _, line := range lists {
		sum += isLineSafe(line)
	}
	fmt.Println("Part 1:", sum)
}

func isLineSafe(line []int) int {
	if checkOrder(line, true) || checkOrder(line, false) {
		return 1
	} else {
		for indx := range line {
			newLine := make([]int, len(line))
			copy(newLine, line)
			newLine = removeAtIndex(newLine, indx)
			if checkOrder(newLine, true) || checkOrder(newLine, false) {
				return 1
			}
		}
		return 0
	}
}

// checkOrder checks if the array is in the specified order (increasing or decreasing)
// and returns a boolean indicating if the order is correct
func checkOrder(arr []int, increasing bool) bool {
	if len(arr) <= 1 {
		return true
	}

	if (increasing && arr[0] < arr[1] && abs(arr[0]-arr[1]) <= 3) ||
		(!increasing && arr[0] > arr[1] && abs(arr[0]-arr[1]) <= 3) {
		return checkOrder(arr[1:], increasing)
	} else {
		return false
	}
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
