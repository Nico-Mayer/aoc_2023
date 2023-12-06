package day_06

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/nico-mayer/aoc_2023/utils"
)

func Part02() {
	data := strings.Split(utils.GetData("06", false), "\n")

	race := parseInput(data)

	fmt.Println(race)

	pw := race.calcPossibleWins()

	fmt.Println(pw)
}

func parseInput(data []string) (r Race) {
	times := parseLine(data[0])
	distances := parseLine(data[1])

	timeStr := ""
	distancesStr := ""

	for _, t := range times {
		tStr := strconv.Itoa(t)
		timeStr += tStr
	}

	for _, d := range distances {
		dStr := strconv.Itoa(d)
		distancesStr += dStr
	}

	time, err := strconv.Atoi(timeStr)
	if err != nil {
		log.Fatal(err)
	}
	r.Time = time

	distance, err := strconv.Atoi(distancesStr)
	if err != nil {
		log.Fatal(err)
	}
	r.Distance = distance

	return r
}
