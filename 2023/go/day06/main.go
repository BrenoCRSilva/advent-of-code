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
	time, distance := parseInput(input, 1)
	fmt.Println("Part 1:", getWinningRaces(time, distance))

	time2, distance2 := parseInput(input, 2)
	fmt.Println(time2)
	fmt.Println(distance2)
	fmt.Println("Part 2:", getWinningRaces(time2, distance2))
}

func parseInput(input string, part int) (time []int, distance []int) {
	// part 2(skill issues) made this really ugly... jeez
	lines := strings.Split(input, "\n")
	var singleTime string
	var singleDistance string
	for i, line := range lines {
		a := strings.Split(line, ":")
		data := strings.Trim(a[1], " ")
		vars := strings.Split(data, " ")
		for _, v := range vars {
			v = strings.Trim(v, " ")
			if part == 1 {
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

			if i == 0 {
				singleTime += v
			} else if i == 1 {
				singleDistance += v
			}
		}
	}
	if part == 2 {
		tnumber, _ := strconv.Atoi(singleTime)
		dnumber, _ := strconv.Atoi(singleDistance)
		time = append(time, tnumber)
		distance = append(distance, dnumber)
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
