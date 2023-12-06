package day_06

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/nico-mayer/aoc_2023/utils"
)

type Race struct {
	ID       int
	Time     int
	Distance int
}

func Part01() {
	data := strings.Split(utils.GetData("06", false), "\n")

	times := parseLine(data[0])
	distances := parseLine(data[1])
	var races []Race

	for i := range times {
		races = append(races, Race{
			ID:       i,
			Time:     times[i],
			Distance: distances[i],
		})
	}

	races[2].calcPossibleWins()

	solution := 1
	for _, race := range races {
		pw := race.calcPossibleWins()
		solution *= pw
	}
	fmt.Println(solution)
}

func (r *Race) calcPossibleWins() int {
	var possibleWins int
	goalDistance := r.Distance
	maxTime := r.Time

	for i := 0; i <= r.Time; i++ {
		speed := i
		traveled := speed * (maxTime - i)

		if traveled > goalDistance && maxTime-i > 0 {
			possibleWins++
		}
	}

	return possibleWins
}

func parseLine(line string) []int {
	line = strings.TrimSpace(line)

	parts := strings.Split(line, ":")

	parts[1] = strings.TrimSpace(parts[1])
	fields := strings.Fields(parts[1])

	var values []int

	for _, v := range fields {
		vInt, err := strconv.Atoi(v)
		if err != nil {
			log.Fatal(err)
		}
		values = append(values, vInt)
	}

	return values
}
