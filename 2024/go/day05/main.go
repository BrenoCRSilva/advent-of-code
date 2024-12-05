package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	pageMap, updatesMatrix := parseInput(input)
	validUpdates, invalidUpdates := getValidUpdates(pageMap, updatesMatrix)
	sortedUpdates := sortInvalidUpdates(pageMap, invalidUpdates)
	p1 := getMiddlePageCount(validUpdates)
	p2 := getMiddlePageCount(sortedUpdates)
	fmt.Println(p1)
	fmt.Println(p2)
}

func parseInput(input string) (map[int]map[int]struct{}, [][]int) {
	document := strings.Split(input, "\n\n")
	pages := document[0]
	orientations := strings.Split(pages, "\n")
	pageMap := make(map[int]map[int]struct{})
	for _, orientation := range orientations {
		mappings := strings.Split(orientation, "|")
		key, _ := strconv.Atoi(mappings[0])
		value, _ := strconv.Atoi(mappings[1])
		if _, ok := pageMap[key]; !ok {
			pageMap[key] = make(map[int]struct{})
		}
		pageMap[key][value] = struct{}{}
	}
	updates := document[1]
	lines := strings.Split(updates, "\n")
	updatesMatrix := make([][]int, len(lines))
	for i, line := range lines {
		update := strings.Split(line, ",")
		updatesMatrix[i] = make([]int, len(update))
		for j, u := range update {
			updatesMatrix[i][j], _ = strconv.Atoi(u)
		}

	}
	return pageMap, updatesMatrix
}

func getValidUpdates(pageMap map[int]map[int]struct{}, updatesMatrix [][]int) ([][]int, [][]int) {
	validUpdates := make([][]int, 0)
	invalidUpdates := make([][]int, 0)
	for _, update := range updatesMatrix {
		flag := true
		for i := 0; i < len(update)-1; i++ {
			page := update[i]
			for j := i + 1; j < len(update); j++ {
				nextPage := update[j]
				if _, ok := pageMap[page][nextPage]; !ok {
					flag = false
					break
				}
			}
		}
		if flag {
			validUpdates = append(validUpdates, update)
		} else {
			invalidUpdates = append(invalidUpdates, update)
		}
	}
	return validUpdates, invalidUpdates
}

func getMiddlePageCount(validUpdates [][]int) int {
	count := 0
	for _, update := range validUpdates {
		index := len(update) / 2
		count += update[index]
	}
	return count
}

func sortInvalidUpdates(pageMap map[int]map[int]struct{}, invalidUpdates [][]int) [][]int {
	sortedUpdates := make([][]int, 0)
	for _, update := range invalidUpdates {
		for i := 0; i < len(update)-1; i++ {
			for j := i + 1; j < len(update); j++ {
				if _, ok := pageMap[update[i]][update[j]]; !ok {
					update[i], update[j] = update[j], update[i]
				}
			}
		}
		sortedUpdates = append(sortedUpdates, update)
	}
	return sortedUpdates
}
