package day_03

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"

	"github.com/nico-mayer/aoc_2023/utils"
)

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

	var solution int
	var partNumbers = make(map[string]int)
	var symbols []Symbol

	for i, line := range data {
		for j, char := range line {
			if !unicode.IsDigit(char) && char != '.' && char != '\r' {
				symbols = append(symbols, Symbol{
					Char: char,
					Y:    i,
					X:    j,
				})
			}

		}
	}

	for _, sym := range symbols {
		sym.getNeighbors(data)
		for _, n := range sym.Neighbors {
			if unicode.IsDigit(n.Char) {
				key, value := getPartNum(n, data)
				partNumbers[key] = value
			}
		}
	}

	partNumbers = utils.RemoveDuplicates(partNumbers)

	for _, pn := range partNumbers {
		solution += pn
	}

	fmt.Printf("Solution P1: %d\n", solution)
}

func (sym *Symbol) getNeighbors(data []string) {
	appendNeighbor := func(x, y int) {
		neighbor := checkNeighbor(x, y, data)
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

func checkNeighbor(x, y int, data []string) (neighbor rune) {
	neighbor = '.'
	inBounds := (x >= 0 && x < len(data[0]) && y >= 0 && y < len(data))
	if inBounds {
		neighbor = rune(data[y][x])
	}
	return
}

func getPartNum(neighbor Neighbor, data []string) (key string, value int) {
	valueString := ""
	valueString += string(neighbor.Char)
	key = fmt.Sprintf("(%d,%d)", neighbor.X, neighbor.Y)

	x := neighbor.X
	y := neighbor.Y

	for i := x - 1; i >= 0; i-- {
		if unicode.IsDigit(rune(data[y][i])) {
			pos := fmt.Sprintf("(%d,%d)", i, y)
			valueString = string(data[y][i]) + valueString
			key = pos + key
		} else {
			break
		}
	}

	for i := x + 1; i < len(data[y]); i++ {
		if unicode.IsDigit(rune(data[y][i])) {
			pos := fmt.Sprintf("(%d,%d)", i, y)
			valueString += string(data[y][i])
			key = key + pos
		} else {
			break
		}
	}

	value, _ = strconv.Atoi(valueString)

	return key, value
}
