package main

import (
	"adventofcode2023/util"
	"fmt"
	"log"
	"math"
	"strings"
	"time"
)

// https://adventofcode.com/2023/day/11
func main() {
	galaxyFile := util.ReadFile("C:\\Users\\janss\\git\\adventofcode2023\\day11\\input.txt")

	defer util.TimeTrack(time.Now(), "main")

	galaxy := make([]string, 0)
	for galaxyFile.Scan() {
		galaxy = append(galaxy, galaxyFile.Text())
	}

	log.Default().Printf("P1: %d", part1(galaxy, 2))
	log.Default().Printf("P2: %d", part2(galaxy))
}

var space rune = rune('.')
var universe rune = rune('#')

func expandGalaxy(galaxy []string, expansion int) [][]rune {
	expandedGalaxy := make([][]rune, 0)
	for _, galaxyLine := range galaxy {
		expandedGalaxyRow := make([]rune, 0)
		for _, element := range galaxyLine {
			expandedGalaxyRow = append(expandedGalaxyRow, element)
		}
		expandedGalaxy = append(expandedGalaxy, expandedGalaxyRow)
		if !strings.Contains(galaxyLine, string(universe)) {
			for i := 0; i < expansion-1; i++ {
				expandedGalaxy = append(expandedGalaxy, expandedGalaxyRow)
			}
		}
	}

	addedVerticals := 0
	for x := range expandedGalaxy[0] {
		isAllSpace := true
		for y := range expandedGalaxy {
			if expandedGalaxy[y][x+addedVerticals] == universe {
				isAllSpace = false
				break
			}
		}
		if isAllSpace {
			for i := 0; i < expansion-1; i++ {
				for y := range expandedGalaxy {
					expandedGalaxy = insert(expandedGalaxy, x+addedVerticals, y, space)
				}
				addedVerticals++
			}
		}
		// debugGalaxy(expandedGalaxy)
	}
	return expandedGalaxy
}

func insert(a [][]rune, x int, y int, value rune) [][]rune {
	a[y] = append(a[y][:x+1], a[y][x:]...)
	a[y][x] = value
	return a
}

func calculateDistance(u1 Universe, u2 Universe) int {
	return int(math.Abs(float64(u2.x-u1.x))) + int(math.Abs(float64(u2.y-u1.y)))
}

type Universe struct {
	x int
	y int
}

func part1(galaxy []string, expansion int) int {
	expandedGalaxy := expandGalaxy(galaxy, expansion)

	universes := make([]Universe, 0)
	for y := range expandedGalaxy {
		for x := range expandedGalaxy[0] {
			if expandedGalaxy[y][x] == universe {
				universes = append(universes, Universe{x: x, y: y})
			}
		}
	}

	distances := 0
	for i := 0; i < len(universes); i++ {
		for j := i; j < len(universes); j++ {
			// log.Printf("universe %d|%d", i, j)
			distances += calculateDistance(universes[i], universes[j])
		}
	}
	return distances
}

func part2(galaxy []string) int {
	distances := part1(galaxy, 0)
	distance10times := part1(galaxy, 10)
	diff := distance10times - distances

	distances += diff * 1
	distances += diff * 10
	distances += diff * 100
	distances += diff * 1000
	distances += diff * 10000
	distances += diff * 100000
	return distances
}

func debugGalaxy(galaxy [][]rune) {
	for y := range galaxy {
		for x := range galaxy[0] {
			fmt.Print(string(galaxy[y][x]))
		}
		fmt.Println("")
	}
	fmt.Println("---")
}
