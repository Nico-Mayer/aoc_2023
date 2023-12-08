package day_08

import (
	"fmt"
	"strings"

	"github.com/nico-mayer/aoc_2023/utils"
)

func Part01() {
	data := strings.Split(utils.GetData("08", false), "\n")
	instructions := ""
	nodes := make(map[string][]string)

	currentPos := "AAA"
	steps := 0

	for i, line := range data {
		line = strings.TrimSpace(line)

		if i == 0 {
			instructions = line
			continue
		}
		if i == 1 {
			continue
		}

		parseNode(line, nodes)

	}

	for currentPos != "ZZZ" {
		for _, instruction := range instructions {
			pos := nodes[currentPos]
			if instruction == 'L' {
				currentPos = pos[0]
			} else if instruction == 'R' {
				currentPos = pos[1]
			}
			steps++

			if currentPos == "ZZZ" {
				break
			}

		}
	}
	fmt.Printf("Solutions P1 %d", steps)
}

func parseNode(line string, nodes map[string][]string) {
	parts := strings.Split(line, "=")
	pos := strings.TrimSpace(parts[0])

	parts[1] = strings.TrimSpace(parts[1])
	parts[1] = strings.ReplaceAll(parts[1], "(", "")
	parts[1] = strings.ReplaceAll(parts[1], ")", "")

	leftRight := strings.Split(parts[1], ",")
	left := strings.TrimSpace(leftRight[0])
	right := strings.TrimSpace(leftRight[1])

	nodes[pos] = []string{left, right}
}
