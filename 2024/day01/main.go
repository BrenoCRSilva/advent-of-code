package main

import (
	_ "embed"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	keys, values := parseInput(input)
	p1 := part1(keys, values)
	p2 := part2(keys, values)
	fmt.Println(p1)
	fmt.Println(p2)
}

func parseInput(input string) ([]int, []int) {
	line := strings.Split(input, "\n")
	keys := make([]int, len(line))
	values := make([]int, len(line))
	for _, l := range line {
		num := strings.Split(l, "   ")
		var a, b int
		a, _ = strconv.Atoi(num[0])
		b, _ = strconv.Atoi(num[1])
		keys = append(keys, a)
		values = append(values, b)
	}
	return keys, values
}

func part1(keys, values []int) int {
	sort.Ints(keys)
	sort.Ints(values)
	var count int
	for i := 0; i < len(keys); i++ {
		if keys[i] > values[i] {
			count += keys[i] - values[i]
		} else {
			count += values[i] - keys[i]
		}

	}
	return count
}

func part2(keys, values []int) int {
	sort.Ints(keys)
	sort.Ints(values)
	numMap := make(map[int]int)
	for i := 0; i < len(keys); i++ {
		numMap[keys[i]] = 0
	}
	for i := 0; i < len(values); i++ {
		if _, ok := numMap[values[i]]; ok {
			numMap[values[i]] += values[i]
		}
	}
	var count int
	for _, v := range numMap {
		count += v
	}
	return count
}
