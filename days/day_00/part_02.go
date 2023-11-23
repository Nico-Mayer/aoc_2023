package day_00

import (
	"fmt"
	"strings"

	"github.com/nico-mayer/aoc_2023/utils"
)

func Part02() {
	data := strings.Split(utils.GetData("00", false), "\n")

	fmt.Printf("%#v", data)
}
