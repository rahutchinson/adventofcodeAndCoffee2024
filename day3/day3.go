package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	// Open the file
	file, err := os.Open("day3input.txt")
	// file, err := os.Open("day3ex.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var sb strings.Builder
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		sb.WriteString(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	line := sb.String()

	// regex to parse the line
	// Define the regex pattern
	pattern := `mul\((\d{1,3}),(\d{1,3})\)|do\(\)|don't\(\)`
	re := regexp.MustCompile(pattern)

	// Find all matches
	findAllStringSubmatch := re.FindAllStringSubmatch(line, -1)
	sum := 0
	current := "do"
	for _, match := range findAllStringSubmatch {
		if match[0] == "don't()" {
			current = "don't"
			continue
		} else if match[0] == "do()" {
			current = "do"
			continue
		}
		if current == "do" {
			fmt.Println("Match:", match)
			sum += (toInt(match[1]) * toInt(match[2]))
		}
	}
	fmt.Println("Sum:", sum)
	//33516115 is too low

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
