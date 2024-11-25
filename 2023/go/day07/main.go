package main

import (
	_ "embed"
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func init() {
	input = strings.TrimSuffix(input, "\n")
	if len(input) == 0 {
		log.Fatal("input file is empty")
	}
}

func main() {
	game := parseInput(input)
	sort.Sort(game)
	winnings := 0
	for i, hand := range game {
		winnings += hand.value * (i + 1)
		fmt.Println(hand.cards, hand.typeRank, hand.value, i+1, winnings)
	}
}

type Hand struct {
	cards           string
	value, typeRank int
}

type Game []Hand

func (g Game) Len() int {
	return len(g)
}

func (g Game) Swap(i, j int) {
	g[i], g[j] = g[j], g[i]
}

func (g Game) Less(i, j int) bool {
	hand1 := g[i].typeRank
	hand2 := g[j].typeRank
	if hand1 == hand2 {
		return tieBreaker(g[i], g[j])
	}
	return hand1 < hand2
}

func parseInput(input string) Game {
	lines := strings.Split(input, "\n")
	var game Game
	for _, line := range lines {
		a := strings.Split(line, " ")
		value, _ := strconv.Atoi(a[1])
		hand := Hand{cards: a[0], value: value}
		hand.typeRank = getTypeRank(hand.cards)
		game = append(game, hand)

	}
	return game
}

func getTypeRank(hand string) int {
	gameMap := make(map[rune]int, 5)
	var joker int
	for _, card := range hand {
		if card == 'J' {
			joker++
			continue
		}
		gameMap[card]++
	}
	var max, second int
	for _, count := range gameMap {
		if count > max {
			second = max
			max = count
		} else if count > second {
			second = count
		}
	}
	max += joker
	switch {
	case max == 5:
		return 7
	case max == 4:
		return 6
	case max == 3 && second == 2:
		return 5
	case max == 3:
		return 4
	case max == 2 && second == 2:
		return 3
	case max == 2:
		return 2
	default:
		return 1
	}
}

func tieBreaker(hand1, hand2 Hand) bool {
	cardRank := map[rune]int{
		'J': 0,
		'2': 1, '3': 2, '4': 3, '5': 4, '6': 5, '7': 6, '8': 7, '9': 8,
		'T': 9, 'Q': 10, 'K': 11, 'A': 12,
	}
	for i := 0; i < len(hand2.cards); i++ {
		if cardRank[rune(hand2.cards[i])] > cardRank[rune(hand1.cards[i])] {
			return true
		} else if cardRank[rune(hand2.cards[i])] < cardRank[rune(hand1.cards[i])] {
			return false
		}
	}
	panic("identical hands")
}
