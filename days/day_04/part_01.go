package day_04

import (
	"fmt"
	"strings"

	"github.com/nico-mayer/aoc_2023/utils"
)

func Part01() {
	data := strings.Split(utils.GetData("04", true), "\n")

	fmt.Printf("%#v", data)
}
