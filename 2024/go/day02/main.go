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
	fmt.Println(part1(numMatrix))
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

func part1(numMatrix [][]int) int {
	count := 0
	for i := 0; i < len(numMatrix); i++ {
		countFlag := false
		for j := 0; j < len(numMatrix[i])-1; j++ {
			diff := numMatrix[i][j] - numMatrix[i][j+1]
			if numMatrix[i][0] > numMatrix[i][1] {
				if numMatrix[i][j] > numMatrix[i][j+1] && diff <= 3 && diff >= 1 {
					countFlag = true
				} else {
					countFlag = false
					break
				}
			} else if numMatrix[i][0] < numMatrix[i][1] {
				if numMatrix[i][j] < numMatrix[i][j+1] && diff >= -3 && diff <= -1 {
					countFlag = true
				} else {
					countFlag = false
					break
				}
			}
		}
		if countFlag {
			count++
		}
	}
	return count
}
