package main

import (
	"fmt"
	"sort"
	"strings"
	"time"

	"AdventOfCode/utils"
)

type HandType int

const (
	Undefined HandType = iota
	HighCard
	OnePair
	TwoPair
	ThreeOfAKind
	FullHouse
	FourOfAKind
	FiveOfAKind
)

type Hand struct {
	cards    string
	handType HandType
	bid      int
}

func parseInput(input []string) []Hand {
	hands := []Hand{}
	for _, line := range input {
		hand := parseLine(line)
		hands = append(hands, hand)
	}
	return hands
}

func parseLine(line string) Hand {
	parts := strings.Split(line, " ")

	cards := parts[0]
	handType := Undefined
	bidStr := parts[1]
	bid := utils.ToInt(bidStr)

	hand := Hand{
		cards:    cards,
		handType: handType,
		bid:      bid,
	}
	return hand
}

func getHandType(cards string) HandType {
	counts := map[rune]int{}
	for _, card := range cards {
		counts[card]++
	}

	maxCount := 0
	for _, count := range counts {
		if count > maxCount {
			maxCount = count
		}
	}

	switch maxCount {
	case 5:
		return FiveOfAKind
	case 4:
		return FourOfAKind
	case 3:
		if len(counts) == 2 {
			return FullHouse
		}
		return ThreeOfAKind
	case 2:
		if len(counts) == 3 {
			return TwoPair
		}
		return OnePair
	default:
		return HighCard
	}
}

func lessHand(hand1, hand2 Hand) bool {
	if hand1.handType != hand2.handType {
		return hand1.handType < hand2.handType
	}

	cardValues := map[byte]int{
		'2': 2,
		'3': 3,
		'4': 4,
		'5': 5,
		'6': 6,
		'7': 7,
		'8': 8,
		'9': 9,
		'T': 10,
		'J': 11,
		'Q': 12,
		'K': 13,
		'A': 14,
	}

	for i := 0; i < 5; i++ {
		value1 := cardValues[hand1.cards[i]]
		value2 := cardValues[hand2.cards[i]]
		if value1 != value2 {
			return value1 < value2
		}
	}

	return false
}

func getHandTypeWithJoker(cards string) HandType {
	counts := map[rune]int{}
	jokerCount := 0
	for _, card := range cards {
		if card == 'J' {
			jokerCount++
		} else {
			counts[card]++
		}
	}

	maxCount := 0
	for _, count := range counts {
		if count > maxCount {
			maxCount = count
		}
	}
	maxCount += jokerCount

	switch maxCount {
	case 5:
		return FiveOfAKind
	case 4:
		return FourOfAKind
	case 3:
		if len(counts) == 2 {
			return FullHouse
		}
		return ThreeOfAKind
	case 2:
		if len(counts) == 3 {
			return TwoPair
		}
		return OnePair
	default:
		return HighCard
	}
}

func lessHandWithJoker(hand1, hand2 Hand) bool {
	if hand1.handType != hand2.handType {
		return hand1.handType < hand2.handType
	}

	cardValues := map[byte]int{
		'2': 2,
		'3': 3,
		'4': 4,
		'5': 5,
		'6': 6,
		'7': 7,
		'8': 8,
		'9': 9,
		'T': 10,
		'J': 0,
		'Q': 12,
		'K': 13,
		'A': 14,
	}

	for i := 0; i < 5; i++ {
		value1 := cardValues[hand1.cards[i]]
		value2 := cardValues[hand2.cards[i]]
		if value1 != value2 {
			return value1 < value2
		}
	}

	return false
}

func Part1(input []string) int {
	hands := parseInput(input)
	for i := 0; i < len(hands); i++ {
		hands[i].handType = getHandType(hands[i].cards)
	}

	sort.Slice(hands, func(i, j int) bool {
		return lessHand(hands[i], hands[j])
	})

	res := 0
	for i, hand := range hands {
		res += hand.bid * (i + 1)
	}
	return res
}

func Part2(input []string) int {
	hands := parseInput(input)
	for i := 0; i < len(hands); i++ {
		hands[i].handType = getHandTypeWithJoker(hands[i].cards)
	}

	sort.Slice(hands, func(i, j int) bool {
		return lessHandWithJoker(hands[i], hands[j])
	})

	res := 0
	for i, hand := range hands {
		res += hand.bid * (i + 1)
	}
	return res
}

func main() {
	fileName := "input.txt"
	input := utils.ReadFile(fileName)

	start1 := time.Now()
	fmt.Println("Answer 1 : ", Part1(input))
	fmt.Println(time.Since(start1))

	start2 := time.Now()
	fmt.Println("Answer 2 : ", Part2(input))
	fmt.Println(time.Since(start2))
}
