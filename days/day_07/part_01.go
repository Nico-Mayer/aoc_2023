package day_07

import (
	"fmt"
	"log"
	"slices"
	"strconv"
	"strings"

	"github.com/nico-mayer/aoc_2023/utils"
)

type Hand struct {
	Cards string
	Bid   int
	Type  int
}

func Part01() {
	data := strings.Split(utils.GetData("07", false), "\n")

	var cardValues = map[string]int{
		"2": 2,
		"3": 3,
		"4": 4,
		"5": 5,
		"6": 6,
		"7": 7,
		"8": 8,
		"9": 9,
		"T": 10,
		"J": 11,
		"Q": 12,
		"K": 13,
		"A": 14,
	}

	var hands []Hand

	for _, line := range data {
		line = strings.TrimSpace(line)
		hand := parseHand(line)
		hands = append(hands, hand)
	}

	slices.SortFunc(hands, func(hand1, hand2 Hand) int {
		if hand1.Type != hand2.Type {
			return hand1.Type - hand2.Type
		} else {
			return compareHands(hand1, hand2, cardValues)
		}
	})

	solution := 0

	for i, hand := range hands {
		solution += (i + 1) * hand.Bid
	}

	fmt.Printf("Solution: %d", solution)
}

func parseHand(line string) (hand Hand) {
	parts := strings.Split(line, " ")
	cards := strings.TrimSpace(parts[0])
	bidStr := strings.TrimSpace(parts[1])

	hand.Cards = cards
	hand.Type = getHandType(hand.Cards)

	bid, err := strconv.Atoi(bidStr)
	if err != nil {
		log.Fatal(err)
	}

	hand.Bid = bid

	return hand
}

func compareHands(hand1, hand2 Hand, cardValues map[string]int) int {

	for i := range hand1.Cards {
		h1Card := string(hand1.Cards[i])
		h2Card := string(hand2.Cards[i])

		if h1Card != h2Card {
			h1Value := cardValues[h1Card]
			h2Value := cardValues[h2Card]

			return h1Value - h2Value
		}
	}

	return 0
}

func getHandType(cards string) int {
	var handType = map[string]int{
		"highCard":     1,
		"onePair":      2,
		"twoPair":      3,
		"threeOfAKind": 4,
		"fullHouse":    5,
		"fourOfAKind":  6,
		"fiveOfAKind":  7,
	}

	uniqueCards := getUniqueCardsFromHand(cards)

	switch len(uniqueCards) {
	case 1:
		return handType["fiveOfAKind"]
	case 2:
		for _, uc := range uniqueCards {
			if strings.Count(cards, uc) == 4 {
				return handType["fourOfAKind"]
			} else if strings.Count(cards, uc) == 3 {
				return handType["fullHouse"]
			}
		}
	case 3:
		for _, uc := range uniqueCards {
			if strings.Count(cards, uc) == 3 {
				return handType["threeOfAKind"]
			} else if strings.Count(cards, uc) == 2 {
				return handType["twoPair"]
			}
		}
	case 4:
		return handType["onePair"]
	case 5:
		return handType["highCard"]

	}
	return -1
}

func getUniqueCardsFromHand(cards string) []string {
	var uniqueCards []string

	for _, card := range cards {
		if !slices.Contains(uniqueCards, string(card)) {
			uniqueCards = append(uniqueCards, string(card))
		}
	}

	return uniqueCards
}
