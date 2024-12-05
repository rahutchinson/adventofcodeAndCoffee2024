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
	file, err := os.Open("day5input.txt")
	// file, err := os.Open("day5ex.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var pairs [][2]int
	var listOfLists [][]int
	scanner := bufio.NewScanner(file)
	isSecondSet := false

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			isSecondSet = true
			continue
		}

		if !isSecondSet {
			// Process the first set of numbers
			parts := strings.Split(line, "|")
			if len(parts) == 2 {
				num1, err1 := strconv.Atoi(parts[0])
				num2, err2 := strconv.Atoi(parts[1])
				if err1 == nil && err2 == nil {
					pairs = append(pairs, [2]int{num1, num2})
				}
			}
		} else {
			// Process the second set of numbers
			parts := strings.Split(line, ",")
			var list []int
			for _, part := range parts {
				num, err := strconv.Atoi(part)
				if err == nil {
					list = append(list, num)
				}
			}
			listOfLists = append(listOfLists, list)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	hash := make(map[int][]int)

	// Part 1
	for _, pair := range pairs {
		hash[pair[0]] = append(hash[pair[0]], pair[1])
	}

	var validUpdates [][]int
	for _, list := range listOfLists {
		if checkIfValid(list, hash) {
			validUpdates = append(validUpdates, list)
		}
	}

	sum := 0
	for _, list := range validUpdates {
		middleIndex := len(list) / 2
		// Get the middle element
		sum += list[middleIndex]
	}
	fmt.Println("Part 1:", sum)
	// Part 2
	for i, vals := range hash {
		fmt.Println("Hash:", i, vals)
	}

	var invalidUpdates [][]int
	for _, list := range listOfLists {
		if !checkIfValid(list, hash) {
			fixedList := list
			for !checkIfValid(fixedList, hash) {
				if !checkIfValid(fixedList, hash) {
					fmt.Println("BAD List:", list)
					fixedList = fixInvalidUpdate(fixedList, hash)
				}
			}
			invalidUpdates = append(invalidUpdates, fixedList)
		}
	}

	sumInValid := 0
	for _, list := range invalidUpdates {
		middleIndex := len(list) / 2
		// Get the middle element
		sumInValid += list[middleIndex]
	}
	fmt.Println("Part 2:", sumInValid)
	// 5786 too high
	// 5705 too high
	// 5479 
}

func fixInvalidUpdate(list []int, hash map[int][]int) []int {
	var fixedList []int
	for i := 0; i < len(list); i++ {
		if hash[list[i]] != nil {
			fixedList = append(fixedList, list[i])
			for _, toCheck := range hash[list[i]] {
				if checkIfBefore(toCheck, list[:i]) {
					fmt.Println("toCheck:", toCheck, fixedList)
					fixedList = removeValue(fixedList, toCheck)
					fmt.Println("Fixed list after remove:", fixedList)
					fixedList = append(fixedList, toCheck)
				}
			}
		} else {
			fixedList = append(fixedList, list[i])
		}
		fmt.Println("Fixed list:", fixedList)
	}
	fmt.Println("Fixed list:", fixedList)
	return fixedList
}

func checkIfValid(list []int, hash map[int][]int) bool {
	for i := 0; i < len(list); i++ {
		if hash[list[i]] != nil {
			if checkIfValInUpdateInvalid(list[:i], hash[list[i]]) {
				return false
			}
		}
	}
	return true
}

func checkIfValInUpdateInvalid(list []int, toCheck []int) bool {
	for _, check := range toCheck {
		if checkIfBefore(check, list) {
			return true
		}
	}
	return false
}

func checkIfBefore(toCheck int, list []int) bool {
	for _, num := range list {
		if toCheck == num {
			return true
		}
	}
	return false
}

// Function to remove the first occurrence of a value from a list of integers
func removeValue(slice []int, value int) []int {
	if len(slice) == 1 {
		return []int{}
	}
	for i, v := range slice {
		if v == value {
			return append(slice[:i], slice[i+1:]...)
		}
	}
	return slice
}
