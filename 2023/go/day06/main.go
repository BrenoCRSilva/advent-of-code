package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

type Race struct {
	time, distance int
}

func main() {
	line := parseInput(input)
	fmt.Println(line)
}

func parseInput(input string) [][]string {
	lines := strings.Split(input, "\n")
	var listOfLines [][]string
	for _, line := range lines {
		a := strings.Split(line, " ")
		listOfLines = append(listOfLines, a[1:])
	}
	return listOfLines
}

func getWinningRaces(races []Race) int {
	var count int
	var speed int
	for _, race := range races {
		for speed < race.time {
			if (race.time-speed)*speed > race.distance {
				count++
			}
			speed++
		}
	}
	return count
}
