package utils

import (
	"log"
	"os"
)

func GetData(day string, example bool) string {
	path := "input/day_" + day + "/"
	if example {
		path += "example.txt"
	} else {
		path += "input.txt"
	}

	res, err := os.ReadFile(path)

	if err != nil {
		log.Println("Error while loading data: ")
		log.Println(err)
	}

	return string(res)
}
