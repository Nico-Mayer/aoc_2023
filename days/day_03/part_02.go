package day_03

import (
	"strings"

	"github.com/nico-mayer/aoc_2023/utils"
)

type Gear struct {
	Sym   Symbol
	Valid bool
	Parts []PartNumber
}

func Part02() {
	data := strings.Split(utils.GetData("03", false), "\n")

	var runesData [][]rune
	var gears []Gear
	var solution int

	for i, line := range data {
		var runesRow []rune
		for j, char := range line {
			if char == '*' {
				gears = append(gears, Gear{
					Sym: Symbol{Char: char,
						Y: i,
						X: j},
				})
			}

			if char != '\r' {
				runesRow = append(runesRow, char)
			}
		}
		runesData = append(runesData, runesRow)
	}

	for _, gear := range gears {
		gear.Sym.getNeighbors(runesData)

		for _, nb := range gear.Sym.Neighbors {
			if isNumber(nb.Char) {
				partNum := getPartNum(nb, runesData)
				gear.Parts = append(gear.Parts, partNum)
			}
		}

		gear.filterParts()
		gear.checkIfValid()

		if gear.Valid {
			solution += gear.clacRatio()
		}

	}
	println(solution)

}

func (gear *Gear) filterParts() {
	uniqueKeys := make(map[string]struct{})
	var filteredNbPartNumbers []PartNumber
	for _, parts := range gear.Parts {
		if _, exists := uniqueKeys[parts.Key]; !exists {
			uniqueKeys[parts.Key] = struct{}{}
			filteredNbPartNumbers = append(filteredNbPartNumbers, parts)
		}
	}
	gear.Parts = filteredNbPartNumbers
}

func (gear *Gear) checkIfValid() {
	if len(gear.Parts) == 2 {
		gear.Valid = true
	} else {
		gear.Valid = false
	}
}

func (gear *Gear) clacRatio() int {
	one := gear.Parts[0].Value
	two := gear.Parts[1].Value

	return one * two
}
