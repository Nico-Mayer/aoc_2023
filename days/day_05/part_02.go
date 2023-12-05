package day_05

import (
	"fmt"
	"strings"

	"github.com/nico-mayer/aoc_2023/utils"
)

type SeedRange struct {
	Start int
	End   int
}

func Part02() {
	data := strings.Split(utils.GetData("05", false), "\n")

	chunks := parseIntoChunks(data)
	seedsToPlant := getSeedsToPlant(chunks[0][0])
	seedRanges := getSeedRanges(seedsToPlant)

	seedToSoil := chunkToStd(chunks[1])
	soilToFertilizer := chunkToStd(chunks[2])
	fertilizerToWater := chunkToStd(chunks[3])
	waterToLight := chunkToStd(chunks[4])
	lightToTemperature := chunkToStd(chunks[5])
	temperatureToHumidity := chunkToStd(chunks[6])
	humidityToLocation := chunkToStd(chunks[7])

	/* var locationMaps []LocationMap

	locationMaps = append(locationMaps, seedToSoil.Maps...)
	locationMaps = append(locationMaps, soilToFertilizer.Maps...)
	locationMaps = append(locationMaps, fertilizerToWater.Maps...)
	locationMaps = append(locationMaps, waterToLight.Maps...)
	locationMaps = append(locationMaps, lightToTemperature.Maps...)
	locationMaps = append(locationMaps, temperatureToHumidity.Maps...)
	locationMaps = append(locationMaps, humidityToLocation.Maps...) */

	var validSeeds []int
	for _, seedRange := range seedRanges {
		fmt.Printf("Range: %v\n", seedRange)

		validSeeds = seedToSoil.seedRangeToValidSeeds(seedRange)
		//fmt.Printf("valid: %v\n", validSeeds)

		val := 0

		for _, seed := range validSeeds {
			location := seedToSoil.convert(seed)
			location = soilToFertilizer.convert(location)
			location = fertilizerToWater.convert(location)
			location = waterToLight.convert(location)
			location = lightToTemperature.convert(location)
			location = temperatureToHumidity.convert(location)
			location = humidityToLocation.convert(location)

			if val == 0 {
				val = location
			}

			if location < val {
				val = location
			}
		}

		fmt.Printf("Value: %d\n", val)

	}

}

func (stp *SrcToDst) seedRangeToValidSeeds(seedRange SeedRange) []int {
	var validSeeds []int
	intersect := false

	for _, m := range stp.Maps {
		intersect = false

		mMin := m.Source
		mMax := m.Source + m.Range

		if (mMin <= seedRange.End && mMax >= seedRange.Start) || (seedRange.Start <= mMax && seedRange.End >= mMin) {
			intersect = true
		}
		// fmt.Printf("Map:%v SeedRange:%v Intersect:%v\n", m, seedRange, intersect)

		if intersect {
			overlapStart := max(mMin, seedRange.Start)
			overlapEnd := min(mMax, seedRange.End)

			for i := overlapStart; i <= overlapEnd; i++ {
				validSeeds = append(validSeeds, i)
			}
		}

	}
	return validSeeds
}

func getSeedRanges(arr []int) (seedRanges []SeedRange) {
	for i := range arr {
		if i%2 == 0 {
			seedRanges = append(seedRanges, SeedRange{
				Start: arr[i],
				End:   arr[i] + arr[i+1],
			})
		}
	}

	return seedRanges
}
