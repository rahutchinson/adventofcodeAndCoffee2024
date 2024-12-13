package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readInput(filePath string) []string {
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return lines
}

type Game struct {
	xA     int
	yA     int
	xB     int
	yB     int
	prizeX int
	prizeY int
}

func main() {
	lines := readInput("day13input.txt")
	// lines := readInput("day13ex.txt")
	games1 := make([]Game, 0)
	games2 := make([]Game, 0)
	for i := 0; i < len(lines); i += 4 {
		splitLineA := strings.Split(lines[i], " ")
		xA, _ := strconv.Atoi(splitLineA[2][2:4])
		yA, _ := strconv.Atoi(splitLineA[3][2:])
		splitLineB := strings.Split(lines[i+1], " ")
		xB, _ := strconv.Atoi(splitLineB[2][2:4])
		yB, _ := strconv.Atoi(splitLineB[3][2:])
		splitLineP := strings.Split(lines[i+2], " ")
		p := strings.Split(splitLineP[1], "=")[1]
		prizeX, _ := strconv.Atoi(p[:len(p)-1])
		prizeY, _ := strconv.Atoi(strings.Split(splitLineP[2], "=")[1])
		part2Add := 10000000000000
		games1 = append(games1, Game{xA, yA, xB, yB, prizeX, prizeY})
		games2 = append(games2, Game{xA, yA, xB, yB, prizeX + part2Add, prizeY + part2Add})
	}

	tokens := 0
	for _, g := range games1 {
		canWin, a, b := canGameBeWon(g)
		fmt.Println(canWin, a, b)
		if canWin {
			tokens += a * 3
			tokens += b
		}
	}
	fmt.Println(tokens)
	// 19908 too low
	// 31714 too high
	// 26810

	//part2
	tokens2 := 0
	for _, g := range games2 {
		tokens2 += calculateTokens(g)
	}
	fmt.Println(tokens2)
}

func canGameBeWon(g Game) (bool, int, int) {
	xPossible := findAllSolutions(g.xA, g.xB, g.prizeX)
	yPossible := findAllSolutions(g.yA, g.yB, g.prizeY)
	union := intersectionPairs(xPossible, yPossible)
	if len(union) > 0 {
		a, b := findMinXPair(union)
		return true, a, b
	}
	return false, 0, 0
}

// findAllSolutions finds all integer pairs (x, y) such that A*x + B*y = C
func findAllSolutions(A, B, C int) [][2]int {
	var solutions [][2]int

	// Define a reasonable range for x
	// This range can be adjusted based on the problem constraints
	for x := -1000; x <= 1000; x++ {
		if (C-A*x)%B == 0 {
			y := (C - A*x) / B
			solutions = append(solutions, [2]int{x, y})
		}
	}

	return solutions
}

// intersectionPairs returns the intersection of two slices of pairs (x, y)
func intersectionPairs(pairs1, pairs2 [][2]int) [][2]int {
	pairMap := make(map[[2]int]bool)
	var intersection [][2]int

	// Add pairs from the first slice to the map
	for _, pair := range pairs1 {
		pairMap[pair] = true
	}

	// Check pairs from the second slice for presence in the map
	for _, pair := range pairs2 {
		if pairMap[pair] {
			intersection = append(intersection, pair)
		}
	}

	return intersection
}

// findMinXPair finds the pair (x, y) with the minimum x value from a list of pairs
func findMinXPair(pairs [][2]int) (int, int) {
	minPair := pairs[0]
	minX := pairs[0][0]

	for _, pair := range pairs {
		if pair[0] < minX {
			minX = pair[0]
			minPair = pair
		}
	}

	return minPair[0], minPair[1]
}

func extendedGCD(a, b int) (int, int, int) {
	x0, x1, y0, y1 := 1, 0, 0, 1
	for b != 0 {
		q := a / b
		a, b = b, a%b
		x0, x1 = x1, x0-q*x1
		y0, y1 = y1, y0-q*y1
	}
	return a, x0, y0
}

func calculateTokens(machine Game) int {
	// use Cramer's rule to solve the system of equations
	a, b := solveEquation(machine)

	// if no valid solution, return 0
	if a <= 0 || b <= 0 {
		return 0
	}

	// calculate the total tokens (3 per A press, 1 per B press)
	return (3 * a) + b
}

func solveEquation(m Game) (int, int) {
	// using cramer's rule:
	// for system of eqs:
	// ax*A + bx*B = px (x eq)
	// ay*A + by*B = py (y eq)

	// determinant of coefficients matrix
	d := m.xA*m.yB - m.yA*m.xB

	// determinants for A and B
	d1 := m.prizeX*m.yB - m.prizeY*m.xB
	d2 := m.prizeY*m.xA - m.prizeX*m.yA

	// check if we have whole number sols
	if d1%d != 0 || d2%d != 0 {
		return 0, 0
	}

	return d1 / d, d2 / d
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
