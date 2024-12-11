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
	file, err := os.Open("day11input.txt")
	// file, err := os.Open("day11input copy.txt")
	// file, err := os.Open("day11ex.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var numbers []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line) // Split the line into fields (words)
		for _, part := range parts {
			num, err := strconv.Atoi(part) // Convert each part to an integer
			if err != nil {
				fmt.Println("Error converting to integer:", err)
				continue
			}
			numbers = append(numbers, num) // Append the integer to the slice
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	currentRocks := make(map[int]int)
	for _, rock := range numbers {
		currentRocks[rock]++
	}

	numOfTimes := 75

	// Iteration logic:
	for i := 0; i < numOfTimes; i++ {
		nextRocks := make(map[int]int)
		for rock, count := range currentRocks {
			processed := handleRock(rock)
			for _, newRock := range processed {
				nextRocks[newRock] += count
			}
		}
		currentRocks = nextRocks
		fmt.Println("Iteration", i+1, ":", sumRocks(currentRocks))
	}
}

func sumRocks(rockMap map[int]int) int {
	sum := 0
	for _, count := range rockMap {
		sum += count
	}
	return sum
}

// If the stone is engraved with the number 0, it is replaced by a stone engraved with the number 1.
// If the stone is engraved with a number that has an even number of digits, it is replaced by two stones. The left half of the digits are engraved on the new left stone, and the right half of the digits are engraved on the new right stone. (The new numbers don't keep extra leading zeroes: 1000 would become stones 10 and 0.)
// If none of the other rules apply, the stone is replaced by a new stone; the old stone's number multiplied by 2024 is engraved on the new stone.
func handleRock(rock int) []int {
	if rock == 0 {
		return []int{1}
	} else if len(strconv.Itoa(rock))%2 == 0 {
		return splitRock(rock)
	} else {
		return []int{rock * 2024}
	}
}

func splitRock(rock int) []int {
	rockStr := strconv.Itoa(rock)
	rock1, _ := strconv.Atoi(rockStr[:len(rockStr)/2])
	rock2, _ := strconv.Atoi(rockStr[len(rockStr)/2:])

	return []int{rock1, rock2}
}
