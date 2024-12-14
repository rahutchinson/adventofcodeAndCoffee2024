// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// 	"strconv"
// 	"strings"
// )

// func readInput(filePath string) []string {
// 	file, err := os.Open(filePath)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer file.Close()

// 	scanner := bufio.NewScanner(file)
// 	var lines []string
// 	for scanner.Scan() {
// 		lines = append(lines, scanner.Text())
// 	}

// 	if err := scanner.Err(); err != nil {
// 		panic(err)
// 	}

// 	return lines
// }

// type Robot struct {
// 	x, y       int
// 	velX, velY int
// }

// func main() {
// 	lines := readInput("day14input.txt")
// 	// lines := readInput("day14ex.txt")
// 	robots := make([]Robot, 0)
// 	for _, line := range lines {
// 		loc := strings.Split(strings.Split(strings.Split(line, " ")[0], "=")[1], ",")
// 		x, _ := strconv.Atoi(loc[0])
// 		y, _ := strconv.Atoi(loc[1])
// 		vel := strings.Split(strings.Split(strings.Split(line, " ")[1], "=")[1], ",")
// 		velx, _ := strconv.Atoi(vel[0])
// 		vely, _ := strconv.Atoi(vel[1])
// 		robots = append(robots, Robot{x, y, velx, vely})
// 	}

// 	// width := 11
// 	// height := 7
// 	width := 101
// 	height := 103
// 	for i := 0; i < 100; i++ {
// 		for robot := range robots {
// 			robots[robot] = moveRobot(robots[robot], width, height)
// 		}
// 	}

// 	count1 := 0
// 	count2 := 0
// 	count3 := 0
// 	count4 := 0
// 	for robot := range robots {
// 		fmt.Println(robots[robot])
// 		if robots[robot].x < width/2 && robots[robot].y < height/2 {
// 			fmt.Println("1")
// 			count1++
// 		}
// 		if robots[robot].x > width/2 && robots[robot].y < height/2 {
// 			fmt.Println("2")
// 			count2++
// 		}
// 		if robots[robot].x < width/2 && robots[robot].y > height/2 {
// 			fmt.Println("3")
// 			count3++
// 		}
// 		if robots[robot].x > width/2 && robots[robot].y > height/2 {
// 			fmt.Println("4")
// 			count4++
// 		}
// 	}
// 	fmt.Println(count1, count2, count3, count4)
// 	fmt.Println(count1 * count2 * count3 * count4)
// 	// 499 too low

// }

// func moveRobot(robot Robot, width, height int) Robot {
// 	robot.x = (robot.x + robot.velX) % width
// 	robot.y = (robot.y + robot.velY) % height
// 	if robot.x < 0 {
// 		robot.x += width
// 	}
// 	if robot.y < 0 {
// 		robot.y += height
// 	}
// 	return robot
// }

// func abs(x int) int {
// 	if x < 0 {
// 		return -x
// 	}
// 	return x
// }

// func displayGrid(robots []Robot, h, w int) {
// 	// Create a 2D grid initialized to 0
// 	grid := make([][]int, h)
// 	for i := range grid {
// 		grid[i] = make([]int, w)
// 	}

// 	// Populate the grid with robot positions
// 	for _, r := range robots {
// 		if r.y >= 0 && r.y < h && r.x >= 0 && r.x < w {
// 			grid[r.y][r.x]++
// 		}
// 	}

// 	// Print the grid
// 	for i := 0; i < h; i++ {
// 		for j := 0; j < w; j++ {
// 			fmt.Printf("%d ", grid[i][j])
// 		}
// 		fmt.Println()
// 	}
// }
