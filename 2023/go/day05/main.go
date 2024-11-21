package main

import (
	_ "embed"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func init() {
	input = strings.TrimRight(input, "\n")
	if len(input) == 0 {
		panic("empty input file")
	}
}

func ParseInput(input string) ([]int, [][]Map) {
	var listOfSeeds []int
	var listOfMapRanges [][]Map
	document := strings.Split(input, "\n\n")
	seeds := strings.Split(document[0], " ")
	for i := 1; i < len(seeds); i++ {
		seed, _ := strconv.Atoi(seeds[i])
		listOfSeeds = append(listOfSeeds, seed)
	}
	for j := 1; j < len(document); j++ {
		mapLines := strings.Split(document[j], "\n")[1:]
		for _, line := range mapLines {
			var ranges []Map
			fields := strings.Split(line, " ")
			source, _ := strconv.Atoi(fields[1])
			destination, _ := strconv.Atoi(fields[0])
			count, _ := strconv.Atoi(fields[2])
			ranges = append(ranges, Map{
				source:      source,
				destination: destination,
				count:       count,
			})
			listOfMapRanges = append(listOfMapRanges, ranges)
		}
	}
	return listOfSeeds, listOfMapRanges
}

type Map struct {
	source, destination, count int
}
