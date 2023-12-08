package day_08

import (
	"fmt"
	"strings"

	"github.com/nico-mayer/aoc_2023/utils"
)

var startPositions []string

func Part02() {
	data := strings.Split(utils.GetData("08", false), "\n")
	instructions := ""
	nodes := make(map[string][]string)

	for i, line := range data {
		line = strings.TrimSpace(line)

		if i == 0 {
			instructions = line
			continue
		}
		if i == 1 {
			continue
		}

		parseNode2(line, nodes)
	}

	var minStepsPerNode []int
	fmt.Println(instructions)

	for _, pos := range startPositions {
		steps := 0
		// fmt.Printf("Start: %s\n", pos)

		for pos[2] != 'Z' {
			for _, instruction := range instructions {
				steps++

				if instruction == 'L' {
					pos = nodes[pos][0]
				} else if instruction == 'R' {
					pos = nodes[pos][1]
				}

				//fmt.Printf("From:%s To:%s\n", old, pos)
				if pos[2] == 'Z' {
					break
				}
			}
		}

		minStepsPerNode = append(minStepsPerNode, steps)
		//fmt.Printf("Took %d Steps\n", steps)
		//fmt.Println()
	}
	fmt.Println(minStepsPerNode)

	fmt.Println(lcm(minStepsPerNode))

}

func lcm(numbers []int) int {
	result := numbers[0]
	for i := 1; i < len(numbers); i++ {
		result = (result * numbers[i]) / gcd(result, numbers[i])
	}
	return result
}

func gcd(a, b int) int {
	if a == 0 {
		return b
	}
	return gcd(b%a, a)
}

func parseNode2(line string, nodes map[string][]string) {
	parts := strings.Split(line, "=")
	pos := strings.TrimSpace(parts[0])

	if pos[2] == 'A' {
		startPositions = append(startPositions, pos)
	}

	parts[1] = strings.TrimSpace(parts[1])
	parts[1] = strings.ReplaceAll(parts[1], "(", "")
	parts[1] = strings.ReplaceAll(parts[1], ")", "")

	leftRight := strings.Split(parts[1], ",")
	left := strings.TrimSpace(leftRight[0])
	right := strings.TrimSpace(leftRight[1])

	nodes[pos] = []string{left, right}
}
