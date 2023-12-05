package main

import (
	"adventofcode2023/util"
	"bufio"
	"log"
	"math"
	"strings"
	"time"
)

// https://adventofcode.com/2023/day/5
func main() {
	almanac := util.ReadFile("./day5/input.txt")
	defer util.TimeTrack(time.Now(), "main")

	seeds := parseSeeds(almanac)
	log.Println(seeds)
	categoryMap := parseMapping(almanac)
	log.Println(categoryMap)

	log.Default().Printf("P1: %d", part1(seeds, categoryMap))
	log.Default().Printf("P2: %d", part2(seeds, categoryMap))
}

type CategoryMapping struct {
	destinationStart int
	sourceStart      int
	rangeLength      int
}

func parseSeeds(almanac *bufio.Scanner) []int {
	almanac.Scan()
	return util.SpaceDelimitedStringToIntSlice(strings.Split(almanac.Text(), ": ")[1])
}

func parseMapping(almanac *bufio.Scanner) map[string][]CategoryMapping {
	categoryMap := make(map[string][]CategoryMapping)

	// currentCategoryMappings := make([]CategoryMapping, 0)
	currentCategory := ""
	for almanac.Scan() {
		almanacLine := almanac.Text()

		if almanacLine == "" {
			continue
		}

		category := getCategory(almanacLine)
		if category != "" {
			currentCategory = category
			continue
		}

		mappingLine := util.SpaceDelimitedStringToIntSlice(almanac.Text())
		categoryMap[currentCategory] = append(categoryMap[currentCategory], CategoryMapping{
			destinationStart: mappingLine[0],
			sourceStart:      mappingLine[1],
			rangeLength:      mappingLine[2],
		})
	}
	return categoryMap
}

func getCategory(almanacLine string) string {
	switch almanacLine {
	case "seed-to-soil map:":
		return "soil"
	case "soil-to-fertilizer map:":
		return "fertilizer"
	case "fertilizer-to-water map:":
		return "water"
	case "water-to-light map:":
		return "light"
	case "light-to-temperature map:":
		return "temperature"
	case "temperature-to-humidity map:":
		return "humidity"
	case "humidity-to-location map:":
		return "location"
	default:
		return ""
	}
}

func nextCategory(category string) string {
	switch category {
	case "soil":
		return "fertilizer"
	case "fertilizer":
		return "water"
	case "water":
		return "light"
	case "light":
		return "temperature"
	case "temperature":
		return "humidity"
	case "humidity":
		return "location"
	default:
		return ""
	}
}

func part1(seeds []int, categoryMap map[string][]CategoryMapping) int {
	lowest := math.MaxInt
	for _, seed := range seeds {
		lowest = int(math.Min(float64(lowest), float64(traverse("soil", seed, categoryMap))))
	}
	return lowest
}

func traverse(category string, value int, categoryMap map[string][]CategoryMapping) int {
	categoryMapping := categoryMap[category]
	found := value
	for _, c := range categoryMapping {
		if value >= c.sourceStart && value < c.sourceStart+c.rangeLength {
			found = c.destinationStart + value - c.sourceStart
		}
	}
	if category == "location" {
		return found
	}
	return traverse(nextCategory(category), found, categoryMap)
}

func part2(seeds []int, categoryMap map[string][]CategoryMapping) int {
	lowest := math.MaxInt
	var seedStart int
	for i, v := range seeds {
		if i%2 == 0 {
			seedStart = v
		} else {
			for j := seedStart; j < seedStart+v; j++ {
				lowest = int(math.Min(float64(lowest), float64(traverse("soil", j, categoryMap))))
			}
		}
	}
	return lowest
}
