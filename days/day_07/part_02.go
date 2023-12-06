package day_07

import (
	"fmt"
	"strings"

	"github.com/nico-mayer/aoc_2023/utils"
)

func Part02() {
	data := strings.Split(utils.GetData("07", true), "\n")

	fmt.Printf("%#v", data)
}