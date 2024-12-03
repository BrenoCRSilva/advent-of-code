package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

func init() {
	input = strings.TrimRight(input, "\n")
	if len(input) == 0 {
		panic("input file is empty")
	}
}

func main() {
	head, last, instructions, nodeMap, startNodes, _ := parseInput(input)
	result := Traverse(head, last, instructions, nodeMap)
	fmt.Println(result)
	fmt.Println(startNodes)
	fmt.Println(LCM(16043, 20777, 13939, 18673, 11309, 17621))
}

type Node struct {
	value, left, right string
}

func parseInput(input string) (head, last, instructions string, nodeMap map[string]*Node, startNodes []string, zNumber int) {
	document := strings.Split(input, "\n\n")
	lines := strings.Split(document[1], "\n")
	nodeMap = make(map[string]*Node, len(lines))
	instructions = strings.TrimSuffix(document[0], " ")
	zNumber = 0
	for _, line := range lines {
		value := line[:3]
		left := line[7:10]
		right := line[12:15]
		nodeMap[value] = &Node{value: value, left: left, right: right}
		if strings.HasSuffix(value, "A") {
			startNodes = append(startNodes, value)
		}
		if strings.HasSuffix(value, "Z") {
			zNumber++
		}
	}
	head = "AAA"
	last = "ZZZ"
	return head, last, instructions, nodeMap, startNodes, zNumber
}

func Traverse(head, last, instructions string, nodeMap map[string]*Node) int {
	current := nodeMap[head]
	var step int
	i := 0
	for current != nodeMap[last] {
		step++
		if i == len(instructions) {
			i = 0
		}
		if instructions[i] == 'R' {
			current = nodeMap[current.right]
		} else {
			current = nodeMap[current.left]
		}
		i++
	}
	return step
}

func TraverseSimultaneously(start, instructions string, nodeMap map[string]*Node, zNumber int) []int {
	current := nodeMap[start]
	var zNodes []int
	i := 0
	step := 0
	for len(zNodes) < zNumber {
		step++
		if i == len(instructions) {
			i = 0
		}
		if strings.HasSuffix(current.value, "Z") {
			zNodes = append(zNodes, step)
		}
		if instructions[i] == 'R' {
			current = nodeMap[current.right]
		} else {
			current = nodeMap[current.left]
		}
		i++
	}
	return zNodes
}

func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}
