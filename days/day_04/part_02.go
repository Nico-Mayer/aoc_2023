package day_04

import (
	"fmt"
	"strings"

	"github.com/nico-mayer/aoc_2023/utils"
)

func Part02() {
	data := strings.Split(utils.GetData("04", false), "\n")

	// var solution int
	var cards []Card

	for _, line := range data {
		card := parseCard(line)
		cards = append(cards, card)
	}

	for i := 0; i < len(cards); i++ {
		card := cards[i]
		// fmt.Printf("Card:%d v:%d\n", card.Id, card.winingCards())
		winning := card.winingCards()

		for j := 0; j < winning; j++ {
			newIndex := card.Id + j
			if newIndex < len(cards) {
				cards = append(cards, cards[newIndex])
				// fmt.Printf("Add Card:%d\n", cards[newIndex].Id)
			}
		}
		// fmt.Println()
	}

	fmt.Println(len(cards))
}

func (card *Card) winingCards() int {
	var value int

	for _, pulled := range card.Pulled {
		for _, winning := range card.Winning {
			if pulled == winning {
				value++
			}
		}
	}

	return value
}
