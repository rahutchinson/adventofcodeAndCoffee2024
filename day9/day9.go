package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Set struct {
	num   int
	space int
	id    int
}

type File struct {
	id    int
	space bool
}

func main() {
	// Open the file
	file, err := os.Open("day9input.txt")
	// file, err := os.Open("day9ex.txt")
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
	// fmt.Println(line)

	vals := []Set{}
	identifier := 0
	for i := 0; i < len(line); i += 2 {
		if i+1 >= len(line) {
			vals = append(vals, Set{int(line[i] - '0'), 0, identifier})
		} else {
			vals = append(vals, Set{int(line[i] - '0'), int(line[i+1] - '0'), identifier})
			identifier++
		}
	}
	fmt.Println(vals)

	finalString := []File{}
	for _, val := range vals {
		finalString = append(finalString, repeat(val.id, val.num, false)...)
		finalString = append(finalString, repeat(val.id, val.space, true)...)
	}
	// fmt.Println(finalString)

	reduced := finalString
	// for i := len(finalString) - 1; i >= 0; i-- {
	// 	if !finalString[i].space {
	// 		nextSpace := firstSpace(reduced)
	// 		if nextSpace > i {
	// 			break
	// 		}
	// 		reduced = replaceCharAtIndex(reduced, nextSpace, File{finalString[i].id, false})
	// 		reduced = replaceCharAtIndex(reduced, i, File{000, true})
	// 	}
	// }
	fmt.Println(reduced)
	checksum := int64(0)
	// for i, mul := range reduced {
	// 	if !mul.space {
	// 		checksum += int64((i) * int(mul.id))
	// 	}
	// }
	fmt.Println(checksum)

	//part 2
	fmt.Println(part2(line))
}

func toCode(input string) []int {
	chars := strings.Split(input, "")
	code := make([]int, len(chars))
	for i, val := range chars {
		code[i], _ = strconv.Atoi(val)
	}
	return code
}

func part2(input string) int {
	code := toCode(input)

	blocksLen := 0
	for _, val := range code {
		blocksLen += val
	}

	blocks := make([]int, blocksLen)
	recs := make(map[int]int)
	i := 0

	for j := 0; j < len(code); j++ {
		recs[j] = i
		if j%2 == 0 {
			for k := code[j]; k > 0; k-- {
				blocks[i] = j / 2
				i++
			}
		} else {
			i += code[j]
		}
	}

	id := 1
	for j := len(code) - 1; j >= 0; j -= 2 {
		for k := id; k <= j; k += 2 {
			if code[k] >= code[j] {
				i := recs[k]
				for l := 0; l < code[j]; l++ {
					blocks[i] = j / 2
					i++
				}

				i = recs[j]
				for l := 0; l < code[j]; l++ {
					blocks[i] = 0
					i++
				}

				code[k] -= code[j]
				recs[k] += code[j]
				code[j] = 0

				if code[id] == 0 {
					id += 2
				}
				break
			}
		}
	}

	checksum := 0
	for i, val := range blocks {
		checksum += i * val
	}

	return checksum
}

func repeat(s int, count int, isSpace bool) []File {
	var result []File
	for i := 0; i < count; i++ {
		result = append(result, File{s, isSpace})
	}
	return result
}

// Function to find the first character in a string
func firstSpace(files []File) int {

	for i, r := range files {
		if r.space {
			return i
		}
	}
	return -1
}

// Function to replace a character in a string at a specific index
func replaceCharAtIndex(files []File, index int, newChar File) []File {
	out := []File{}
	for i, r := range files {
		if i == index {
			out = append(out, newChar)
		} else {
			out = append(out, r)
		}
	}
	return out
}
