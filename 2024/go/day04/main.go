package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	matrix := parseInput(input)
	p1 := getXmas(matrix)
	p2 := getMasOfX(matrix)
	fmt.Println(p1)
	fmt.Println(p2)
}

func parseInput(input string) [][]string {
	lines := strings.Split(input, "\n")
	matrix := make([][]string, len(lines))
	for i, line := range lines {
		matrix[i] = make([]string, len(line))
		letters := strings.Split(line, "")
		for j, letter := range letters {
			matrix[i][j] = letter
		}
	}
	return matrix
}

type X struct {
	topRight    string
	topLeft     string
	bottomRight string
	bottomLeft  string
	center      string
}

func (x *X) countXOfMas() int {
	diagonalLeft := x.topLeft + x.center + x.bottomRight
	diagonalRight := x.topRight + x.center + x.bottomLeft
	if (diagonalLeft == "MAS" || diagonalLeft == "SAM") && (diagonalRight == "MAS" || diagonalRight == "SAM") {
		return 1
	}
	return 0
}

func getMasOfX(matrix [][]string) int {
	count := 0
	for i := 1; i < len(matrix)-1; i++ {
		for j := 1; j < len(matrix[i])-1; j++ {
			if matrix[i][j] == "A" {
				x := X{
					topRight:    matrix[i-1][j+1],
					topLeft:     matrix[i-1][j-1],
					bottomRight: matrix[i+1][j+1],
					bottomLeft:  matrix[i+1][j-1],
					center:      matrix[i][j],
				}
				count += x.countXOfMas()
			}
		}
	}
	return count
}

func getXmas(matrix [][]string) int {
	count := 0
	for i := 0; i < len(matrix); i++ {
		fmt.Println(matrix[i])
		for j := 0; j < len(matrix[i]); j++ {
			if j < len(matrix[i])-3 {
				count += getHorizontal(matrix[i], j)
			}
			if i < len(matrix)-3 {
				count += getVertical(matrix, i, j)
				count += getDiagonal(matrix, i, j)
			}
		}
	}
	return count
}

func getHorizontal(line []string, index int) int {
	word := ""
	for i := 0; i < 4; i++ {
		word += line[index+i]
	}
	if word == "SAMX" || word == "XMAS" {
		return 1
	}
	return 0
}

func getVertical(matrix [][]string, length int, width int) int {
	word := ""
	for i := 0; i < 4; i++ {
		word += matrix[length+i][width]
	}
	if word == "SAMX" || word == "XMAS" {
		return 1
	}
	return 0
}

func getDiagonal(matrix [][]string, length int, width int) int {
	count := 0
	word := ""
	word2 := ""
	for i := 0; i < 4; i++ {
		if width < 3 {
			word += matrix[length+i][width+i]
		} else if width >= len(matrix[length])-3 {
			word2 += matrix[length+i][width-i]
		} else {
			word += matrix[length+i][width-i]
			word2 += matrix[length+i][width+i]
		}

	}
	if word == "SAMX" || word == "XMAS" {
		count++
	}
	if word2 == "SAMX" || word2 == "XMAS" {
		count++
	}
	return count
}
