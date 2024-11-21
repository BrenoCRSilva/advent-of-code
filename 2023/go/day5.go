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

}

type Map struct {
	source, destination, count int
}
