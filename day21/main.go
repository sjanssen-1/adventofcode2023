package main

import (
	"adventofcode2023/util"
	"log"
	"math"
	"time"
)

// https://adventofcode.com/2023/day/21
func main() {
	garden := util.ScannerToStringSlice(*util.ReadFile("/Users/stefjanssens/git/adventofcode2023/day21/demo_input.txt"))
	defer util.TimeTrack(time.Now(), "main")

	gardenGraph, start := parseGarden(garden)

	log.Default().Printf("P1: %d", part1(gardenGraph, start))
	log.Default().Printf("P2: %d", part2())
}

type Coordinate struct {
	x, y int
}

func parseGarden(garden []string) (map[Coordinate][]Coordinate, Coordinate) {
	gardenGraph := make(map[Coordinate][]Coordinate)
	var start Coordinate

	for y := 0; y < len(garden); y++ {
		for x := 0; x < len(garden[0]); x++ {
			gardenPlot := garden[y][x]
			switch gardenPlot {
			case '.', 'S':
				gardenGraph[Coordinate{x, y}] = getNeighbours(garden, x, y)
			}

			if gardenPlot == 'S' {
				start = Coordinate{x, y}
			}
		}
	}

	return gardenGraph, start
}

func getNeighbours(garden []string, x int, y int) []Coordinate {
	neighbours := make([]Coordinate, 0)

	if x-1 > 0 && (garden[y][x-1] == '.' || garden[y][x-1] == 'S') {
		neighbours = append(neighbours, Coordinate{x - 1, y})
	}

	if x+1 < len(garden[0]) && (garden[y][x+1] == '.' || garden[y][x+1] == 'S') {
		neighbours = append(neighbours, Coordinate{x + 1, y})
	}

	if y-1 > 0 && (garden[y-1][x] == '.' || garden[y-1][x] == 'S') {
		neighbours = append(neighbours, Coordinate{x, y - 1})
	}

	if y+1 < len(garden) && (garden[y+1][x] == '.' || garden[y+1][x] == 'S') {
		neighbours = append(neighbours, Coordinate{x, y + 1})
	}

	return neighbours
}

func part1(garden map[Coordinate][]Coordinate, start Coordinate) int {

	reachedEdges := make([]Coordinate, 0)

	for i := 0; i < 6; i++ {

	}
}

func part2() int {
	return 0
}

func dijkstra(garden map[Coordinate][]Coordinate, start Coordinate) {
	distances := map[Coordinate]int{}

	for key, _ := range garden {
		distances[key] = math.MaxInt
	}

	distances[start] = 0

	visited := map[Coordinate]bool{}

	toVisit := []Coordinate{start}

	for len(toVisit) > 0 {
		current := toVisit[0]
		toVisit = toVisit[1:]

		for _, neighbour := range garden[current] {
			if !visited[neighbour] && distances[neighbour] > distances[current]+1 {
				distances[neighbour] = distances[current] + 1
				toVisit = append(toVisit, neighbour)
			}
		}
		visited[current] = true
	}

	log.Print(distances)
}
