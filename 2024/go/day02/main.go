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

func isSafeHigher(list []int) bool {
	for i := 0; i < len(list)-1; i++ {
		a := list[i]
		b := list[i+1]
		if a-b <= 3 && a-b >= 1 {
			continue
		} else {
			return false
		}
	}
	return true
}

func isSafeLower(list []int) bool {
	for i := 0; i < len(list)-1; i++ {
		a := list[i]
		b := list[i+1]
		if a-b >= -3 && a-b <= -1 {
			continue
		} else {
			return false
		}
	}
	return true
}

func getSafeReports(numMatrix [][]int) int {
	count := 0
	for _, list := range numMatrix {
		if isSafeHigher(list) || isSafeLower(list) {
			count++
		}
	}
	return count
}

func dampener(list []int, id int) []int {
	dampened := make([]int, 0)
	for i, l := range list {
		if i == id {
			continue
		}
		dampened = append(dampened, l)
	}
	return dampened
}

func getDampenedReports(numMatrix [][]int) int {
	count := 0
	for _, list := range numMatrix {
		if isSafeHigher(list) || isSafeLower(list) {
			count++
			continue
		} else {
			for j := 0; j < len(list); j++ {
				dampened := dampener(list, j)
				if isSafeHigher(dampened) || isSafeLower(dampened) {
					count++
					break
				}
			}
		}
	}
	return count
}
