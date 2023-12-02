package day_01

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/nico-mayer/aoc_2023/utils"
)

func Part02() {
	data := strings.Split(utils.GetData("01", false), "\n")

	var result int
	replaceValues := map[string]string{
		"one":   "o1e",
		"two":   "t2o",
		"three": "th3ee",
		"four":  "fo4r",
		"five":  "fi5e",
		"six":   "s6x",
		"seven": "se7en",
		"eight": "ei8ht",
		"nine":  "ni9e",
	}

	for i, line := range data {
		for oldStr, newStr := range replaceValues {
			line = strings.ReplaceAll(line, oldStr, newStr)
		}
		data[i] = line
	}

	for _, modifiedLine := range data {
		firstDigitFound := false
		firstDigit := -1
		lastDigit := -1
		for _, c := range modifiedLine {
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
