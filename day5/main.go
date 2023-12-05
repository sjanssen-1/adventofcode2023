package main

import (
	"adventofcode2023/util"
	"bufio"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/Goldziher/go-utils/sliceutils"
)

// https://adventofcode.com/2023/day/5
func main() {
	almanac := util.ReadFile("./day5/input.txt")
	defer util.TimeTrack(time.Now(), "main")

	log.Default().Printf("P1: %d", part1())
	log.Default().Printf("P2: %d", part2())
}

type CategoryMap struct {
	category         string
	destinationStart int64
	sourceStart      int64
}

func parseSeeds(almanac bufio.Scanner) []int {
	return sliceutils.Map(strings.Split(strings.Split(almanac.Text(), "seeds: ")[1], " "), func(value string, index int, slice []string) int {
		number, _ := strconv.Atoi(value)
		return number
	})
}

func parseMapping(almanac bufio.Scanner) map[string]CategoryMap {
	categoryMaps := make(map[string]CategoryMap)

	currentCategory := ""
	for almanac.Scan() {
		almanacLine := almanac.Text()
	}

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

	}
}

func part1() int {
	return 0
}

func part2() int {
	return 0
}
