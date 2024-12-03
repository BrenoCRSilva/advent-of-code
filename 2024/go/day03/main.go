package main

import (
	_ "embed"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	mulList := parseMul(input)
	fmt.Println(getResult(mulList))
	instructionsList := parseMulInstructions(input)
	matrix := getMulMatrix(instructionsList)
	mulList2 := parseValidMul(matrix)
	fmt.Println(getResult(mulList2))
}

func parseMul(input string) []string {
	re := regexp.MustCompile(`mul\(\d+,\d+\)`)
	mulList := re.FindAllString(input, -1)
	return mulList

}

func getResult(mulList []string) int {
	count := 0
	for _, mul := range mulList {
		if mul == "do()" {
			continue
		}
		mul = strings.TrimPrefix(mul, "mul(")
		mul = strings.TrimSuffix(mul, ")")
		nums := strings.Split(mul, ",")
		a, _ := strconv.Atoi(nums[0])
		b, _ := strconv.Atoi(nums[1])
		count += a * b
	}
	return count
}

func parseMulInstructions(input string) []string {
	re := regexp.MustCompile(`mul\(\d+,\d+\)|do\(\)|don't\(\)`)
	instructionsList := re.FindAllString(input, -1)
	return instructionsList

}

func getMulMatrix(instructionsList []string) [][]string {
	doNotList := make([]int, 0)
	matrix := make([][]string, 0)
	for i, mul := range instructionsList {
		if mul == "don't()" {
			doNotList = append(doNotList, i)
		}
	}
	matrix = append(matrix, instructionsList[0:doNotList[0]])
	for i := 0; i < len(doNotList)-1; i++ {
		before := doNotList[i] + 1
		after := doNotList[i+1]
		matrix = append(matrix, instructionsList[before:after])
	}
	matrix = append(matrix, instructionsList[doNotList[len(doNotList)-1]:][1:])
	return matrix
}

func parseValidMul(matrix [][]string) []string {
	mulList := append([]string{}, matrix[0]...)
	for i, line := range matrix {
		for j, mul := range line {
			if mul == "do()" {
				mulList = append(mulList, matrix[i][j+1:]...)
				break
			}
		}
	}
	return mulList
}
