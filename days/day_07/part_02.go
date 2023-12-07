package day_07

import (
	"fmt"
	"log"
	"slices"
	"strconv"
	"strings"

	"github.com/nico-mayer/aoc_2023/utils"
)

func Part02() {
	data := strings.Split(utils.GetData("07", false), "\n")

	var cardValues = map[string]int{
		"J": 1,
		"2": 2,
		"3": 3,
		"4": 4,
		"5": 5,
		"6": 6,
		"7": 7,
		"8": 8,
		"9": 9,
		"T": 10,
		"Q": 12,
		"K": 13,
		"A": 14,
	}

	var hands []Hand

	for _, line := range data {
		line = strings.TrimSpace(line)
		hand := parseHand2(line)
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

	fmt.Printf("\nSolution: %d", solution)
}

func parseHand2(line string) (hand Hand) {
	parts := strings.Split(line, " ")
	cards := strings.TrimSpace(parts[0])
	bidStr := strings.TrimSpace(parts[1])

	hand.Cards = cards
	hand.Type = getHandType(fillWithJokers(hand.Cards))

	bid, err := strconv.Atoi(bidStr)
	if err != nil {
		log.Fatal(err)
	}

	hand.Bid = bid

	return hand
}

func fillWithJokers(cards string) string {
	if !strings.Contains(cards, "J") {
		return cards
	}

	if cards == "JJJJJ" {
		return "AAAAA"
	}

	uniqueCards := getUniqueCardsFromHand(cards)

	maxCount := 0
	maxCard := ""

	for _, uc := range uniqueCards {
		count := strings.Count(cards, string(uc))

		if count > maxCount && uc != "J" {
			maxCard = uc
			maxCount = count
		}

	}

	return strings.ReplaceAll(cards, "J", maxCard)

}
