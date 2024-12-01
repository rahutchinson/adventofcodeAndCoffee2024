package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	data, err := os.ReadFile("day1input.txt")
	// data, err := os.ReadFile("day1ex.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	list1 := []int{}
	list2 := []int{}
	for i, part := range strings.Fields(string(data)) {
		if number, err := strconv.Atoi(part); err == nil {
			if i%2 == 0 {
				list1 = append(list1, number)
			} else {
				list2 = append(list2, number)
			}
		}
	}

	list1 = quickSort(list1)
	list2 = quickSort(list2)
	sum := 0
	for i, n1 := range list1 {
		sum += abs(n1 - list2[i])
	}
	fmt.Println("Part 1: Sum of differences:", sum)
	part2Sum := 0
	for _, n1 := range list1 {
		part2Sum += n1 * countOccurrencesSorted(list2, n1)
	}
	fmt.Println("Part 2: Sum of products:", part2Sum)
}

func quickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	left, right := 0, len(arr)-1

	pivot := len(arr) / 2

	arr[pivot], arr[right] = arr[right], arr[pivot]

	for i := range arr {
		if arr[i] < arr[right] {
			arr[i], arr[left] = arr[left], arr[i]
			left++
		}
	}

	arr[left], arr[right] = arr[right], arr[left]

	quickSort(arr[:left])
	quickSort(arr[left+1:])

	return arr
}

func abs(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}

func countOccurrencesSorted(list []int, target int) int {
	first := sort.Search(len(list), func(i int) bool { return list[i] >= target })
	if first == len(list) || list[first] != target {
		return 0
	}
	last := sort.Search(len(list), func(i int) bool { return list[i] > target })
	return last - first
}
