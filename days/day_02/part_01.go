package day_02

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/nico-mayer/aoc_2023/utils"
)

const (
	MAX_RED   = 12
	MAX_GREEN = 13
	MAX_BLUE  = 14
)

type Game struct {
	ID   int
	Sets []Set
}
type Set struct {
	Rolls []Roll
}
type Roll struct {
	Count int
	Color string
}

func (game *Game) isValid() bool {
	valid := true
	r := 0
	g := 0
	b := 0

	for _, set := range game.Sets {
		r = 0
		g = 0
		b = 0

		for _, roll := range set.Rolls {

			switch roll.Color {
			case "red":
				r += roll.Count
			case "green":
				g += roll.Count
			case "blue":
				b += roll.Count
			}
		}

		if r > MAX_RED || g > MAX_GREEN || b > MAX_BLUE {
			valid = false
			break
		}

	}

	return valid
}

func Part01() {
	data := strings.Split(utils.GetData("02", false), "\n")
	var solution int

	for _, line := range data {
		game := parseLine(line)
		if game.isValid() {
			solution += game.ID
		}

	}
	fmt.Println(solution)
}

func parseLine(line string) (game Game) {
	parts := strings.Split(line, ":")

	id, _ := strconv.Atoi(strings.ReplaceAll(parts[0], "Game ", ""))
	game.ID = id

	setArr := strings.Split(parts[1], ";")

	for _, set := range setArr {
		set = strings.Trim(set, " ")
		rolls := strings.Split(set, ",")

		var newSet Set

		for _, roll := range rolls {
			roll = strings.Trim(roll, " ")
			parts := strings.Split(roll, " ")

			count, _ := strconv.Atoi(parts[0])
			color := strings.TrimSpace(parts[1])

			var roll Roll = Roll{
				Count: count,
				Color: color,
			}

			newSet.Rolls = append(newSet.Rolls, roll)
		}

		game.Sets = append(game.Sets, newSet)

	}

	return
}
