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
	// file, err := os.Open("day8ex.txt")
	file, err := os.Open("day8input.txt")
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

	aMap := make(map[rune][]Pair)

	for i, line := range data {
		for j, loc := range line {
			if loc != '.' {
				aMap[loc] = append(aMap[loc], Pair{i, j})
			}
		}
	}
	fmt.Println(aMap)
	antiMap := make(map[Pair]int)
	for _, value := range aMap {
		// fmt.Println("Checking", key)
		if len(value) > 1 {
			for i := 0; i < len(value); i++ {
				for _, ant := range value[i+1:] {
					xDiff, yDiff := computeXYDiff(value[i], ant)
					toTest := getOptions(xDiff, yDiff, value[i], ant, len(data), len(data[0]), 1, []Pair{})
					// fmt.Println("Testing", value[i], ant, toTest)
					for _, test := range toTest {
						// fmt.Print("Testing", test, " ")
						if areCollinear(value[i], ant, test) {
							// fmt.Println("Collinear")
							antiMap[test]++
						} else {
							// fmt.Println("Not collinear")
						}
					}

				}
			}
			// every combination of the pairs
			// with x and y diff values
			// if the x

			// should I do some kind of geometry? like on a line?
		}
	}
	// fmt.Println(antiMap)
	count := []Pair{}
	for key := range antiMap {
		if key.X >= 0 && key.X < len(data) && key.Y >= 0 && key.Y < len(data[0]) {
			count = append(count, key)
		}
	}

	for _, key := range count {
		replace(data, key.X, key.Y, '#')
	}
	printBoard(data)

	fmt.Println(len(count))

	//part 2 752 too low

}

func printBoard(data []string) {
	fmt.Println("____________________________________")
	for _, row := range data {
		fmt.Println(row)
	}
	fmt.Println("____________________________________")
}

func replace(data []string, row, col int, val rune) {
	rowS := []rune(data[row])
	rowS[col] = val
	data[row] = string(rowS)
}

func getOptions(xDiffinit, yDiffinit int, aPair, bPair Pair, maxX, maxY int, iter int, list []Pair) []Pair {
	xDiff := xDiffinit * iter
	yDiff := yDiffinit * iter
	if iter > maxX || iter > maxY {
		return list
	} else {
		if xDiff == 0 {
			if aPair.Y < bPair.Y {
				list = append(list, Pair{aPair.X, aPair.Y - yDiff}, Pair{bPair.X, bPair.Y + yDiff})
				return append(list, getOptions(xDiffinit, yDiffinit, aPair, bPair, maxX, maxY, iter+1, list)...)
			} else {
				list = append(list, Pair{aPair.X, aPair.Y + yDiff}, Pair{bPair.X, bPair.Y - yDiff})
				return append(list, getOptions(xDiffinit, yDiffinit, aPair, bPair, maxX, maxY, iter+1, list)...)
			}
		}
		if yDiff == 0 {
			if aPair.X < bPair.X {
				list = append(list, Pair{aPair.X - xDiff, aPair.Y}, Pair{bPair.X + xDiff, bPair.Y})
				return append(list, getOptions(xDiffinit, yDiffinit, aPair, bPair, maxX, maxY, iter+1, list)...)

			} else {
				list = append(list, Pair{aPair.X + xDiff, aPair.Y}, Pair{bPair.X - xDiff, bPair.Y})
				return append(list, getOptions(xDiffinit, yDiffinit, aPair, bPair, maxX, maxY, iter+1, list)...)
			}
		}
		if yDiff == 0 && xDiff == 0 {
			return []Pair{}
		}

		list = append(list, Pair{aPair.X - xDiff, aPair.Y + yDiff})
		list = append(list, Pair{aPair.X - xDiff, aPair.Y - yDiff})
		list = append(list, Pair{aPair.X + xDiff, aPair.Y + yDiff})
		list = append(list, Pair{aPair.X + xDiff, aPair.Y - yDiff})

		list = append(list, Pair{bPair.X + xDiff, bPair.Y + yDiff})
		list = append(list, Pair{bPair.X - xDiff, bPair.Y - yDiff})
		list = append(list, Pair{bPair.X + xDiff, bPair.Y - yDiff})
		list = append(list, Pair{bPair.X - xDiff, bPair.Y + yDiff})
		// for _, val := range list {
		// 	if val.X == aPair.X && val.Y == aPair.Y ||
		// 		val.X == bPair.X && val.Y == bPair.Y {
		// 		list = removeValue(list, val)
		// 	}
		// }
		return append(list, getOptions(xDiffinit, yDiffinit, aPair, bPair, maxX, maxY, iter+1, list)...)

	}
}

// Function to remove the first occurrence of a value from a list of integers
func removeValue(slice []Pair, value Pair) []Pair {
	if len(slice) == 1 && slice[0] == value {
		return []Pair{}
	}
	for i, v := range slice {
		if v == value {
			return append(slice[:i], slice[i+1:]...)
		}
	}
	return slice
}

func computeXYDiff(p1 Pair, p2 Pair) (int, int) {
	return abs(p1.X - p2.X), abs(p1.Y - p2.Y)
}

// Function to check if three points are collinear
func areCollinear(a, b, check Pair) bool {
	// Calculate the area of the triangle formed by the points
	// If the area is zero, the points are collinear
	area := a.X*(b.Y-check.Y) + b.X*(check.Y-a.Y) + check.X*(a.Y-b.Y)
	return area == 0
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
