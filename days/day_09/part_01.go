package day_09

import (
	"fmt"
	"strings"

	"github.com/nico-mayer/aoc_2023/utils"
)

func Part01() {
	data := strings.Split(utils.GetData("09", true), "\n")

	solution := 0

	for _, line := range data {
		line = strings.TrimSpace(line)
		solution += extrapolate(line)
		fmt.Println(line)
	}
}

func extrapolate(line string)int{

	return 0 
}