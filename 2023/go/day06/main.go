package main

import (
	_ "embed"
	"fmt"
	"log"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func init() {
	input = strings.TrimRight(input, "\n")
	if len(input) == 0 {
		log.Fatal("input.txt is empty")
	}
}

func main() {
	time, distance := parseInput(input)
	fmt.Println(time)
	fmt.Println(distance)
	fmt.Println(getWinningRaces(time, distance))
}

func parseInput(input string) (time []int, distance []int) {
	lines := strings.Split(input, "\n")
	for i, line := range lines {
		a := strings.Split(line, ":")
		data := strings.Trim(a[1], " ")
		vars := strings.Split(data, " ")
		for _, v := range vars {
			v = strings.Trim(v, " ")
			number, err := strconv.Atoi(v)
			if err != nil {
				continue
			}
			if i == 0 {
				time = append(time, number)
			} else if i == 1 {
				distance = append(distance, number)
			}
		}
	}
	return time, distance
}

func getWinningRaces(time []int, distance []int) int {
	wins := 1
	for i := 0; i < len(time); i++ {
		var speed int
		var count int
		for speed < time[i] {
			if (time[i]-speed)*speed > distance[i] {
				count++
			}
			speed++
		}
		wins *= count
	}
	return wins
}
