package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	numMatrix := parseInput(input)
	fmt.Println(getSafeReports(numMatrix))
	fmt.Println(getDampenedReports(numMatrix))
}

func parseInput(input string) [][]int {
	lines := strings.Split(input, "\n")
	numMatrix := make([][]int, len(lines))
	for i, line := range lines {
		nums := strings.Split(line, " ")
		numMatrix[i] = make([]int, len(nums))
		for j, num := range nums {
			n, _ := strconv.Atoi(num)
			numMatrix[i][j] = n
		}
	}
	return numMatrix
}

func isSafe(list []int) bool {
	t := 1
	if list[0] > list[1] {
		t = -1
	}
	for i := 0; i < len(list)-1; i++ {
		a := list[i]
		b := list[i+1]
		if a*t < b*t || (a-b)*t > 3 || (a-b)*t < 1 {
			return false
		}
	}
	return true
}

func getSafeReports(numMatrix [][]int) int {
	count := 0
	for _, list := range numMatrix {
		if isSafe(list) {
			count++
		}
	}
	return count
}

func getDampenedReports(numMatrix [][]int) int {
	count := 0
	for _, list := range numMatrix {
		if isSafe(list) {
			count++
		} else {
			for i := 0; i < len(list); i++ {
				newList := append(list[:i], list[i+1:]...)
				if isSafe(newList) {
					count++
					break
				}
			}
		}
	}
	return count
}
