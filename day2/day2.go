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
		if line[0] > line[1] {
			sum += decrease(line, 0, 0)
		} else if line[0] < line[1] {
			sum += increase(line, 0, 0)
		} else {
			sum += 0
		}
	}
	fmt.Println("Part 1:", sum)
}

func decrease(arr []int, prob int, loc int) int {
	if len(arr) == loc+1 {
		return 1
	} else if arr[loc] > arr[loc+1] && (abs(arr[loc]-arr[loc+1]) <= 3) {
		loc += 1
		return decrease(arr, prob, loc)
	} else {
		prob += 1
		if prob > 1 {
			return 0
		}
		// TODO - remove the value at the loc and see if you get a good response then if not remove the loc +1 value
		arrAtLoc := removeAtIndex(arr, loc)
		if decrease(arrAtLoc, prob, loc) == 1 {
			return 1
		}
		arrAtLocLessOne := removeAtIndex(arr, loc-1)
		if decrease(arrAtLocLessOne, prob, loc) == 1 {
			return 1
		}
		return 0
	}
}

func increase(arr []int, prob int, loc int) int {
	if len(arr) == loc+1 {
		return 1
	} else if arr[loc] < arr[loc+1] && (abs(arr[loc]-arr[loc+1]) <= 3) {
		loc += 1
		return increase(arr, prob, loc)
	} else {
		prob += 1
		if prob > 1 {
			return 0
		}
		// TODO - remove the value at the loc and see if you get a good response then if not remove the loc +1 value
		arrAtLoc := removeAtIndex(arr, loc)
		if increase(arrAtLoc, prob, loc) == 1 {
			return 1
		}
		arrAtLocLessOne := removeAtIndex(arr, loc-1)
		if increase(arrAtLocLessOne, prob, loc) == 1 {
			return 1
		}
		return 0
	}
}

func abs(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}

func removeAtIndex(slice []int, index int) []int {
	if index < 0 || index >= len(slice) {
		return slice // Return the original slice if the index is out of range
	}
	return append(slice[:index], slice[index+1:]...)
}
