package day_06

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/nico-mayer/aoc_2023/utils"
)

type Race struct {
	ID       int
	Time     int
	Distance int
}

func Part01() {
	data := strings.Split(utils.GetData("06", false), "\n")
	fmt.Println("Day 6 Part 1:")
	start := time.Now()

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

	solution := 1
	for _, race := range races {
		pw := race.calcPossibleWins()
		solution *= pw
	}
	end := time.Since(start)
	fmt.Printf("Solution: %d", solution)
	fmt.Printf(" Took: %v\n", end)
}

func (r *Race) calcPossibleWins() int {
	var possibleWins int

	for i := 0; i <= r.Time; i++ {
		speed := i
		traveled := speed * (r.Time - i)

		if traveled > r.Distance && r.Time-i > 0 {
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
