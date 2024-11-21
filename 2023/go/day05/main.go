package main

import (
	_ "embed"
	"strings"
	"strconv"
)

//go:embed input.txt
var input string

func init() {
	input = strings.TrimRight(input, "\n")
	if len (input) == 0 {
		panic("empty input.txt")
	}
}

func ParseInput(input string) {
	document := strings.Split(input, "\n\n")
	seeds := strings.Split(document[0]," ")
	for i = 1; i < len(document); i++ {
		mapLines := strings.Split(document[i], "\n")[1:]

}

type Map struct {
	source, destination, count int
}
