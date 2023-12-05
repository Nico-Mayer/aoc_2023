package day_05

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/nico-mayer/aoc_2023/utils"
)

type SrcToDst struct {
	Source      string
	Destination string
	Maps        []LocationMap
}

type LocationMap struct {
	Source      int
	Destination int
	Range       int
}

func Part01() {
	data := strings.Split(utils.GetData("05", false), "\n")

	var solution int

	chunks := parseIntoChunks(data)
	seedsToPlant := getSeedsToPlant(chunks[0][0])

	seedToSoil := chunkToStd(chunks[1])
	soilToFertilizer := chunkToStd(chunks[2])
	fertilizerToWater := chunkToStd(chunks[3])
	waterToLight := chunkToStd(chunks[4])
	lightToTemperature := chunkToStd(chunks[5])
	temperatureToHumidity := chunkToStd(chunks[6])
	humidityToLocation := chunkToStd(chunks[7])

	for _, seed := range seedsToPlant {
		location := seedToSoil.convert(seed)
		location = soilToFertilizer.convert(location)
		location = fertilizerToWater.convert(location)
		location = waterToLight.convert(location)
		location = lightToTemperature.convert(location)
		location = temperatureToHumidity.convert(location)
		location = humidityToLocation.convert(location)

		if solution == 0 {
			solution = location
		} else if location < solution {
			solution = location
		}

	}
	fmt.Printf("Solution: %d", solution)

}

func (srcToDst *SrcToDst) convert(value int) int {
	for _, m := range srcToDst.Maps {
		if value >= m.Source && value < m.Source+m.Range {
			convertedValue := m.Destination + (value - m.Source)

			return convertedValue
		}
	}
	return value
}

func chunkToStd(chunk []string) SrcToDst {
	var srcToDst SrcToDst

	for j, line := range chunk {
		if j == 0 {
			src, dest := parseSrcAndDest(line)
			srcToDst.Source = src
			srcToDst.Destination = dest
			continue
		}

		vMap := parseMap(line)

		srcToDst.Maps = append(srcToDst.Maps, vMap)
	}

	return srcToDst
}

func parseMap(line string) (locationMap LocationMap) {
	parts := strings.Split(line, " ")
	for i, v := range parts {
		vInt, err := strconv.Atoi(v)
		if err != nil {
			log.Fatal(err)
		}

		if i == 0 {
			locationMap.Destination = vInt
		} else if i == 1 {
			locationMap.Source = vInt
		} else {
			locationMap.Range = vInt
		}

	}
	return locationMap
}

func getSeedsToPlant(line string) (seeds []int) {
	parts := strings.Split(line, ":")
	parts[1] = strings.TrimSpace(parts[1])

	seedsStr := strings.Split(parts[1], " ")

	for _, seed := range seedsStr {
		seedInt, err := strconv.Atoi(seed)
		if err != nil {
			log.Fatal(err)
		}
		seeds = append(seeds, seedInt)
	}

	return seeds
}

func parseSrcAndDest(line string) (src, dest string) {
	line = strings.ReplaceAll(line, " map:", "")
	parts := strings.Split(line, "-")

	src = parts[0]
	dest = parts[2]

	return src, dest
}

func parseIntoChunks(data []string) (chunks [][]string) {
	var chunk []string

	for _, line := range data {
		line = strings.TrimSpace(line)

		if len(line) == 0 {
			chunks = append(chunks, chunk)
			chunk = nil
			continue
		}

		chunk = append(chunk, line)
	}

	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}

	return chunks
}
