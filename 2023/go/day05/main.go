package main

import (
	_ "embed"
	"log"
	"math"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func init() {
	input = strings.TrimRight(input, "\n")
	input = strings.TrimSpace(input)
	if len(input) == 0 {
		log.Fatal("input.txt is empty")
	}
}

func main() {
	listOfSeeds, listOfMapBlocks := parseInput(input)
	closest := getClosestLocation(listOfSeeds, listOfMapBlocks)
	println(closest)
}

func getClosestLocation(listOfSeeds []int, listOfMapBlocks [][]Map) int {
	closest := math.MaxInt
	for _, seed := range listOfSeeds {
		location := getMap(seed, listOfMapBlocks)
		if location < closest {
			closest = location
		}
	}
	return closest
}

func parseInput(input string) ([]int, [][]Map) {
	var listOfSeeds []int
	var listOfMapBlocks [][]Map
	document := strings.Split(input, "\n\n")
	seeds := strings.Split(document[0], " ")
	for i := 1; i < len(seeds); i++ {
		seed, _ := strconv.Atoi(seeds[i])
		listOfSeeds = append(listOfSeeds, seed)
	}
	for j := 1; j < len(document); j++ {
		mapBlocks := strings.Split(document[j], "\n")[1:]
		var ranges []Map
		for _, block := range mapBlocks {
			lines := strings.Split(block, " ")
			source, _ := strconv.Atoi(lines[1])
			destination, _ := strconv.Atoi(lines[0])
			count, _ := strconv.Atoi(lines[2])
			ranges = append(ranges, Map{
				source:      source,
				destination: destination,
				count:       count,
			})
		}
		listOfMapBlocks = append(listOfMapBlocks, ranges)
	}
	return listOfSeeds, listOfMapBlocks
}

func getMap(seed int, listOfMapBlocks [][]Map) int {
	locations := seed
	for _, block := range listOfMapBlocks {
		for _, b := range block {
			if seed >= b.source && seed <= b.source+b.count {
				locations = b.destination + (locations - b.source)
				break
			}
		}
	}
	return locations
}

type Map struct {
	source, destination, count int
}
