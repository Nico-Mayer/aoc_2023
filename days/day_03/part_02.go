package day_03

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/nico-mayer/aoc_2023/utils"
)

type Gear struct {
	Sym   Symbol
	Valid bool
	Parts map[string]int
}

func Part02() {
	data := strings.Split(utils.GetData("03", true), "\n")

	var gears []Gear
	var solution int

	for i, line := range data {
		for j, char := range line {
			if char == '*' {
				gears = append(gears, Gear{
					Sym: Symbol{Char: char,
						Y: i,
						X: j},
					Parts: make(map[string]int),
				})
			}
		}

	}

	for _, gear := range gears {
		gear.Sym.getNeighbors(data)

		for _, nb := range gear.Sym.Neighbors {
			if unicode.IsDigit(nb.Char) {
				key, value := getPartNum(nb, data)
				gear.Parts[key] = value
			}
		}

		gear.Parts = utils.RemoveDuplicates(gear.Parts)
		gear.checkIfValid()

		if gear.Valid {
			solution += gear.clacRatio()
		}

	}

	fmt.Printf("Solution P2: %d\n", solution)
}

func (gear *Gear) checkIfValid() {
	if len(gear.Parts) == 2 {
		gear.Valid = true
	} else {
		gear.Valid = false
	}
}

func (gear *Gear) clacRatio() int {
	result := 1

	for _, value := range gear.Parts {
		result *= value
	}

	return result
}
