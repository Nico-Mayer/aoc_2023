package day_03

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/nico-mayer/aoc_2023/utils"
)

type PartNumber struct {
	Value int
	Key   string
}

type Neighbor struct {
	Char rune
	X    int
	Y    int
}

type Symbol struct {
	Char      rune
	X         int
	Y         int
	Neighbors []Neighbor
}

func Part01() {
	data := strings.Split(utils.GetData("03", false), "\n")

	var partNumbers []PartNumber
	var runesData [][]rune
	var symbols []Symbol

	for i, line := range data {
		var runesRow []rune
		for j, char := range line {
			if !isNumber(char) && char != '.' && char != '\r' {
				symbols = append(symbols, Symbol{
					Char: char,
					Y:    i,
					X:    j,
				})
			}

			if char != '\r' {
				runesRow = append(runesRow, char)
			}
		}
		runesData = append(runesData, runesRow)
	}

	for _, sym := range symbols {
		fmt.Printf("S:%c X:%d Y:%d\n", sym.Char, sym.X, sym.Y)
		sym.getNeighbors(runesData)

		for _, neighbor := range sym.Neighbors {
			if isNumber(neighbor.Char) {
				partNumbers = append(partNumbers, getPartNum(neighbor, runesData))
			}
		}

	}

	uniqueKeys := make(map[string]struct{})
	solution := 0

	for _, pn := range partNumbers {
		if _, exists := uniqueKeys[pn.Key]; !exists {
			uniqueKeys[pn.Key] = struct{}{}
			solution += pn.Value
		}
	}

	fmt.Println(solution)
}

func (sym *Symbol) getNeighbors(runeData [][]rune) {
	appendNeighbor := func(x, y int) {
		neighbor := checkNeighbor(x, y, runeData)
		sym.Neighbors = append(sym.Neighbors, Neighbor{
			Char: neighbor,
			X:    x,
			Y:    y,
		})
	}

	appendNeighbor(sym.X-1, sym.Y)   // Left
	appendNeighbor(sym.X+1, sym.Y)   // Right
	appendNeighbor(sym.X, sym.Y-1)   // Up
	appendNeighbor(sym.X, sym.Y+1)   // Down
	appendNeighbor(sym.X-1, sym.Y-1) // TopLeft
	appendNeighbor(sym.X+1, sym.Y-1) // TopRight
	appendNeighbor(sym.X-1, sym.Y+1) // BottomLeft
	appendNeighbor(sym.X+1, sym.Y+1) // BottomRight

}

func checkNeighbor(x, y int, runeData [][]rune) (neighbor rune) {
	neighbor = '.'
	if x >= 0 && x < len(runeData[0]) && y >= 0 && y < len(runeData) {
		neighbor = runeData[y][x]
	}

	return
}

func isNumber(char rune) bool {
	var numbers = map[rune]int{
		'0': 0,
		'1': 1,
		'2': 2,
		'3': 3,
		'4': 4,
		'5': 5,
		'6': 6,
		'7': 7,
		'8': 8,
		'9': 9,
	}
	_, isNumber := numbers[char]
	return isNumber
}

func getPartNum(neighbor Neighbor, runeData [][]rune) PartNumber {
	resString := ""
	resString += string(neighbor.Char)
	resKey := fmt.Sprintf("(%d,%d)", neighbor.X, neighbor.Y)

	x := neighbor.X
	y := neighbor.Y

	for i := x - 1; i >= 0; i-- {
		if isNumber(runeData[y][i]) {
			key := fmt.Sprintf("(%d,%d)", i, y)
			resString = string(runeData[y][i]) + resString
			resKey = key + resKey
		} else {
			break
		}
	}

	for i := x + 1; i < len(runeData[y]); i++ {
		if isNumber(runeData[y][i]) {
			key := fmt.Sprintf("(%d,%d)", i, y)
			resString += string(runeData[y][i])
			resKey = resKey + key
		} else {
			break
		}
	}

	resInt, err := strconv.Atoi(resString)
	if err != nil {
		fmt.Println(err)
	}

	partNum := PartNumber{
		Value: resInt,
		Key:   resKey,
	}

	return partNum
}
