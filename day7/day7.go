package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

type Data struct {
	Result int
	Values []int64
}

func main() {
	// Open the file
	file, err := os.Open("day7input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var dataList []Data
	scanner := bufio.NewScanner(file)
	lineNumber := 0
	for scanner.Scan() {
		lineNumber++
		line := scanner.Text()
		if strings.TrimSpace(line) == "" {
			fmt.Printf("Skipping empty line at %d\n", lineNumber)
			continue
		}

		parts := strings.Split(line, ":")
		if len(parts) != 2 {
			fmt.Printf("Invalid line format at line %d: %s\n", lineNumber, line)
			continue
		}

		// Parse the result
		result, err := strconv.Atoi(strings.TrimSpace(parts[0]))
		if err != nil {
			fmt.Printf("Invalid result at line %d: %s\n", lineNumber, parts[0])
			continue
		}

		// Parse the values
		valueParts := strings.Split(parts[1], " ")
		var values []int64
		for _, part := range valueParts {
			if strings.TrimSpace(part) == "" {
				continue
			}
			value, err := strconv.ParseInt(strings.TrimSpace(part), 10, 64)
			if err != nil {
				fmt.Printf("Invalid value at line %d: %s\n", lineNumber, part)
				continue
			}
			values = append(values, value)
		}

		// Add the struct to the list
		dataList = append(dataList, Data{Result: result, Values: values})
	}
	// Print the map
	fmt.Println("Data map:")
	for key, values := range dataList {
		fmt.Println(key, values)
	}
	fmt.Println(len(dataList))
	// I think we are going to random brute force this
	// nope - try a random and then if its too high change one * for + if too low then change one + for *
	sum := int64(0)
	for _, values := range dataList {
		length := len(values.Values) - 1
		possibleCombos := computeCombos(length)
		success := false
		for i := len(possibleCombos) - 1; i >= 0; i-- {
			// fmt.Println("PossibleCombo:", possibleCombos[i])
			result := compute(values.Values, possibleCombos[i])
			if result == int64(values.Result) {
				sum += int64(values.Result)
				success = true
				break
			}
		}
		if !success {
			// fmt.Println("possibleCombos:", possibleCombos)
			// fmt.Println("Length:", length, "Key:", key, "Values:", values)
		}

	}
	fmt.Println("Sum of the keys:", sum)
	// 4364915411238 too low
	// 4364915411238
	// 4364915411238
	// 4364915411363 wtf
	// part 2
	// 4364915434894 too low
	// 38322057216320
}

func computeCombos(length int) []string {
	var result []string
	generateCombos("", length, &result)
	return result
}

// Helper function to generate combinations recursively
func generateCombos(current string, length int, result *[]string) {
	if length == 0 {
		*result = append(*result, current)
		return
	}
	generateCombos(current+"+", length-1, result)
	generateCombos(current+"*", length-1, result)
	generateCombos(current+"|", length-1, result)
}

// Function to compute the result based on values and combo
func compute(values []int64, combo string) int64 {
	result := int64(values[0])
	for i, char := range combo {
		if char == '+' {
			result += int64(values[i+1])
		} else if char == '*' {
			result *= int64(values[i+1])
		} else if char == '|' {
			results := strconv.Itoa(int(result)) + strconv.Itoa(int(values[i+1]))
			resultt, _ := strconv.Atoi(results)
			result = int64(resultt)

		}
	}
	return result
}

func getRandom() string {
	// Generate a random 1 or 0
	randomValue := rand.Intn(2) // rand.Intn(2) generates a random number in [0, 2), i.e., 0 or 1
	if randomValue == 0 {
		return "+"
	} else {
		return "+"
	}
}

func contains(targets []rune, char rune) bool {
	for _, target := range targets {
		if target == char {
			return true
		}
	}
	return false
}
