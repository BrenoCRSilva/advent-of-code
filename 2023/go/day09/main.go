package main

import (
	_ "embed"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	numbers := parseInput(input)
	p1, p2 := B(numbers)
	println("Part 1:", p1)
	println("Part 2:", p2)
}

func parseInput(input string) [][]int {
	lines := strings.Split(input, "\n")
	numbers := make([][]int, len(lines))
	for i, line := range lines {
		list := strings.Split(line, " ")
		numbers[i] = make([]int, len(list))
		for j, n := range list {
			num, _ := strconv.Atoi(n)
			numbers[i][j] = num
		}
	}
	return numbers
}

func A(nLine []int) (int, int) {
	p1 := nLine[len(nLine)-1]
	p2 := nLine[0]
	base := nLine
	for i := 1; i < len(nLine); i++ {
		lastIdx := len(nLine) - i
		lower := make([]int, lastIdx)
		for j := 0; j < len(lower); j++ {
			lower[j] = base[j+1] - base[j]
		}
		p1 += lower[lastIdx-1]
		p2 -= lower[0]
		base = lower
	}
	return p1, p2
}

func B(numbers [][]int) (int, int) {
	var result1 int
	var result2 int
	for i := 0; i < len(numbers); i++ {
		p1, p2 := A(numbers[i])
		result1 += p1
		result2 += p2
	}
	return result1, result2
}
