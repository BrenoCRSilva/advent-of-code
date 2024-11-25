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
	head, last, instructions, nodeMap := parseInput(input)
	result := Traverse(head, last, instructions, nodeMap)
	fmt.Println(result)
}

type Node struct {
	left, right string
}

func parseInput(input string) (head, last, instructions string, nodeMap map[string]*Node) {
	document := strings.Split(input, "\n\n")
	lines := strings.Split(document[1], "\n")
	nodeMap = make(map[string]*Node, len(lines))
	instructions = strings.TrimSuffix(document[0], " ")
	for _, line := range lines {
		value := line[:3]
		left := line[7:10]
		right := line[12:15]
		nodeMap[value] = &Node{left: left, right: right}
	}
	head = "AAA"
	last = "ZZZ"
	return head, last, instructions, nodeMap
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
