package day_01

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/nico-mayer/aoc_2023/utils"
)

func Part01() {
	data := strings.Split(utils.GetData("01", false), "\n")
	var result int

	for _, line := range data {
		firstDigitFound := false
		firstDigit := -1
		lastDigit := -1
		for _, c := range line {
			v, err := strconv.Atoi(string(c))

			if err != nil {
				continue
			}

			if !firstDigitFound {
				firstDigit = v
				firstDigitFound = true
			}

			lastDigit = v
		}
		resString := fmt.Sprintf("%d%d", firstDigit, lastDigit)
		num, _ := strconv.Atoi(resString)
		result += num
		fmt.Printf("First Digit: %d\n", firstDigit)
		fmt.Printf("Last Digit: %d\n", lastDigit)
		fmt.Println()
	}

	fmt.Println(result)
}
