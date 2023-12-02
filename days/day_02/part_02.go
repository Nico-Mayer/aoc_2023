package day_02

import (
	"fmt"
	"strings"

	"github.com/nico-mayer/aoc_2023/utils"
)

func Part02() {
	data := strings.Split(utils.GetData("02", false), "\n")
	solution := 0
	for _, line := range data {
		game := parseLine(line)
		solution += game.findPower()
	}

	fmt.Println(solution)
}

func (game *Game) findPower() (power int) {
	r := 0
	g := 0
	b := 0

	for _, set := range game.Sets {
		for _, roll := range set.Rolls {
			switch roll.Color {
			case "red":
				if r < roll.Count {
					r = roll.Count
				}
			case "green":
				if g < roll.Count {
					g = roll.Count
				}
			case "blue":
				if b < roll.Count {
					b = roll.Count
				}
			}
		}
	}
	return r * g * b
}
