package main

import (
	"bufio"
	"fmt"
	"os"
)

// Define a struct to represent a pair of integers
type Pair struct {
	X int
	Y int
}

type PairDir struct {
	X   int
	Y   int
	Dir string
}

func main() {
	// Open the file
	// file, err := os.Open("day6ex.txt")
	file, err := os.Open("day6input.txt")
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

	onboard := true
	locX, locY := checkIfOnboard(data)
	direction := "unknown"
	testBoard := make([]string, len(data))
	copy(testBoard, data)
	for onboard == true {
		locX, locY = checkIfOnboard(testBoard)
		if locX == -1 {
			onboard = false
		} else {
			direction = whatDirection(rune(testBoard[locX][locY]))
			obsX, obsY := findObstacle(testBoard, locX, locY, direction)
			if obsX == -1 || obsY == -1 {
				onboard = false
			} else {
				newX, newY := whereToMoveGaurd(obsX, obsY, direction)
				replaceAllBetween(testBoard, locX, locY, newX, newY, 'X')
				replace(testBoard, newX, newY, directionToSymbol(changeDirection(direction)))
			}
		}
	}
	goOffBoard(testBoard, locX, locY, direction)
	// need to add distance from pointing direction to edge of board
	writeToFile(testBoard, "output.txt")
	fmt.Println("Distinct: ", countDistinct(testBoard))

	// part 2
	// Brute force - replace every X with and # and see if the gaurd escapes?
	// but the halting problem.... ok but we can capture location and direction of gaurd
	// and if you see the same location and direction again, you know it will never escape

	// start by finding where every X is
	Xs := findXs(testBoard)
	fmt.Println("Xs: ", len(Xs))
	// og location
	locXog, locYog := checkIfOnboard(data)
	replace(testBoard, locXog, locYog, rune(data[locXog][locYog]))

	loopCount := 0
	checkCount := 0
	fmt.Println("ogloc ", locXog, locYog)
	for _, pair := range Xs {
		checkCount++
		fmt.Print("checking: ", pair)
		testObsBoard := make([]string, len(testBoard))
		copy(testObsBoard, testBoard)
		if pair.X == locXog && pair.Y == locYog {
			fmt.Println("skipping")
		} else {
			replace(testObsBoard, pair.X, pair.Y, 'O')
			if findIfLoop(testObsBoard) {
				loopCount++
			}
		}
		// fmt.Println("loopCount: ", loopCount)
	}

	fmt.Println("part2: ", loopCount)
	fmt.Println("checkCount: ", checkCount)
	// 1695 it too low -- not allowing anything on the same y or x as the starting spot
	// 1968 is still too low
}

func findIfLoop(loopBoard []string) bool {
	onboard := true
	locX, locY := checkIfOnboard(loopBoard)
	direction := "unknown"
	locDirMap := make(map[PairDir]int)
	for onboard == true {
		locX, locY = checkIfOnboard(loopBoard)
		if locDirMap[PairDir{locX, locY, whatDirection(rune(loopBoard[locX][locY]))}] > 1 {
			return true
		} else {
			locDirMap[PairDir{locX, locY, whatDirection(rune(loopBoard[locX][locY]))}] += 1
		}
		if locX == -1 {
			onboard = false
		} else {
			direction = whatDirection(rune(loopBoard[locX][locY]))
			obsX, obsY := findObstacle(loopBoard, locX, locY, direction)
			if obsX == -1 || obsY == -1 {
				onboard = false
			} else {
				newX, newY := whereToMoveGaurd(obsX, obsY, direction)
				replaceAllBetween(loopBoard, locX, locY, newX, newY, 'X')
				replace(loopBoard, newX, newY, directionToSymbol(changeDirection(direction)))
			}
		}
	}
	// fmt.Println(" no loop")

	return false
}

func findXs(data []string) []Pair {
	var pairs []Pair
	for i, row := range data {
		for j := range row {
			if data[i][j] == 'X' {
				pairs = append(pairs, Pair{i, j})
			}
		}
	}
	return pairs
}

func countDistinct(data []string) int {
	sum := 0
	for _, row := range data {
		for _, char := range row {
			if char == 'X' {
				sum++
			}
		}
	}
	return sum
}

func goOffBoard(data []string, locX int, locY int, direction string) {
	switch direction {
	case "left":
		replaceAllBetween(data, locX, locY, locX, 0, 'X')
	case "right":
		replaceAllBetween(data, locX, locY, locX, len(data[locX])-1, 'X')
	case "up":
		replaceAllBetween(data, locX, locY, 0, locY, 'X')
	case "down":
		replaceAllBetween(data, locX, locY, len(data)-1, locY, 'X')
	}

}

func printBoard(data []string) {
	fmt.Println("____________________________________")
	for _, row := range data {
		fmt.Println(row)
	}
	fmt.Println("____________________________________")
}

func directionToSymbol(direction string) rune {
	switch direction {
	case "left":
		return '<'
	case "right":
		return '>'
	case "up":
		return '^'
	case "down":
		return 'v'
	}
	return ' '
}

func replaceAllBetween(data []string, row1, col1, row2, col2 int, val rune) {
	if row1 == row2 {
		if col1 < col2 {
			for i := col1; i <= col2; i++ {
				replace(data, row1, i, val)
			}
		} else {
			for i := col1; i >= col2; i-- {
				replace(data, row1, i, val)
			}
		}
	} else if col1 == col2 {
		if row1 < row2 {
			for i := row1; i <= row2; i++ {
				replace(data, i, col1, val)
			}
		} else {
			for i := row1; i >= row2; i-- {
				replace(data, i, col1, val)
			}
		}
	}
}

func replace(data []string, row, col int, val rune) {
	rowS := []rune(data[row])
	rowS[col] = val
	data[row] = string(rowS)
}

func changeDirection(direction string) string {
	switch direction {
	case "left":
		return "up"
	case "right":
		return "down"
	case "up":
		return "right"
	case "down":
		return "left"
	}
	return "unknown"
}

func whereToMoveGaurd(obsX int, obsY int, direction string) (int, int) {
	switch direction {
	case "left":
		return obsX, obsY + 1
	case "right":
		return obsX, obsY - 1
	case "up":
		return obsX + 1, obsY
	case "down":
		return obsX - 1, obsY
	}
	return -1, -1
}

func findObstacle(data []string, locX int, locY int, direction string) (int, int) {
	switch direction {
	case "left":
		for i := locY; i >= 0; i-- {
			if data[locX][i] == '#' || data[locX][i] == 'O' {
				return locX, i
			}
		}
	case "right":
		for i := locY; i < len(data[locX]); i++ {
			if data[locX][i] == '#' || data[locX][i] == 'O' {
				return locX, i
			}
		}
	case "up":
		for i := locX; i >= 0; i-- {
			if data[i][locY] == '#' || data[i][locY] == 'O' {
				return i, locY
			}
		}
	case "down":
		for i := locX; i < len(data); i++ {
			if data[i][locY] == '#' || data[i][locY] == 'O' {
				return i, locY
			}
		}
	}
	return -1, -1
}

func whatDirection(val rune) string {
	switch val {
	case '<':
		return "left"
	case '>':
		return "right"
	case 'v':
		return "down"
	case '^':
		return "up"
	default:
		return "unknown"
	}
}

func checkIfOnboard(data []string) (int, int) {
	target := []rune{'<', '>', 'v', '^'}
	for i, row := range data {
		for j, char := range row {
			if contains(target, char) {
				return i, j
			}
		}
	}
	return -1, -1 // Return -1, -1 if the value is not found
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

// contains checks if a character is in the list of target characters
func writeToFile(data []string, filename string) {
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, line := range data {
		_, err := writer.WriteString(line + "\n")
		if err != nil {
			fmt.Println("Error writing to file:", err)
			return
		}
	}
	writer.Flush()
}

func contains(targets []rune, char rune) bool {
	for _, target := range targets {
		if target == char {
			return true
		}
	}
	return false
}
