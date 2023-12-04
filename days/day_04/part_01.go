package day_04

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/nico-mayer/aoc_2023/utils"
)

type Card struct {
	Id      int
	Winning []int
	Pulled  []int
}

func Part01() {
	data := strings.Split(utils.GetData("04", false), "\n")

	var solution int
	var cards []Card

	for _, line := range data {
		card := parseCard(line)
		cards = append(cards, card)
	}

	for _, card := range cards {
		//  fmt.Printf("Card:%d v:%d\n", card.Id, card.calcValue())
		solution += card.calcValue()
	}
	fmt.Println(solution)
}

func parseCard(line string) Card {
	parts := strings.Split(line, ":")
	parts[0] = strings.ReplaceAll(parts[0], "Card", "")
	parts[0] = strings.ReplaceAll(parts[0], " ", "")

	cardId, err := strconv.Atoi(parts[0])
	if err != nil {
		fmt.Println(err)
	}

	values := strings.Split(parts[1], "|")
	values[0] = strings.TrimSpace(values[0])
	values[1] = strings.TrimSpace(values[1])

	winningStr, pulledStr := strings.Split(values[0], " "), strings.Split(values[1], " ")

	var winning []int
	var pulled []int

	for _, n := range winningStr {
		if strings.TrimSpace(n) == "" {
			continue
		}

		num, err := strconv.Atoi(n)
		if err != nil {
			log.Fatalf("strconv.Atoi(%v): %v", n, err)
		}

		winning = append(winning, num)
	}

	for _, n := range pulledStr {
		if strings.TrimSpace(n) == "" {
			continue
		}

		num, err := strconv.Atoi(n)
		if err != nil {
			log.Fatalf("strconv.Atoi(%v): %v", n, err)
		}

		pulled = append(pulled, num)
	}

	return Card{
		Id:      cardId,
		Winning: winning,
		Pulled:  pulled,
	}
}

func (card *Card) calcValue() int {
	var value int

	for _, pulled := range card.Pulled {
		for _, winning := range card.Winning {
			if pulled == winning {
				if value == 0 {
					value = 1
				} else {
					value *= 2
				}

			}
		}
	}

	return value
}
