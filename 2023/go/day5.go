package main

import (
	_ "embed"
	"strings"
)

//go:embed day5.txt
var input string

func init() {
	input = strings.TrimRight(input, "\n")
}
func ParseInput(input string) {
	a := Map{source: 0, destination: 1, count: 1}
}

type Map struct {
	source, destination, count int
}
