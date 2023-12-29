package main

import (
	"adventofcode2023/util"
	"fmt"
	"log"
	"math"
	"time"

	"github.com/Goldziher/go-utils/sliceutils"
)

// https://adventofcode.com/2023/day/21
func main() {
	demoGarden := util.ScannerToStringSlice(*util.ReadFile("/Users/stefjanssens/git/adventofcode2023/day21/demo_input.txt"))
	garden := util.ScannerToStringSlice(*util.ReadFile("/Users/stefjanssens/git/adventofcode2023/day21/input.txt"))
	defer util.TimeTrack(time.Now(), "main")

	demoGardenGraph, startDemo := parseGarden(demoGarden)
	gardenGraph, start := parseGarden(garden)

	part1(gardenGraph, start, demoGardenGraph, startDemo)
	part2()
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

	if x-1 >= 0 && (garden[y][x-1] == '.' || garden[y][x-1] == 'S') {
		neighbours = append(neighbours, Coordinate{x - 1, y})
	}

	if x+1 < len(garden[0]) && (garden[y][x+1] == '.' || garden[y][x+1] == 'S') {
		neighbours = append(neighbours, Coordinate{x + 1, y})
	}

	if y-1 >= 0 && (garden[y-1][x] == '.' || garden[y-1][x] == 'S') {
		neighbours = append(neighbours, Coordinate{x, y - 1})
	}

	if y+1 < len(garden) && (garden[y+1][x] == '.' || garden[y+1][x] == 'S') {
		neighbours = append(neighbours, Coordinate{x, y + 1})
	}

	return neighbours
}

func part1(garden map[Coordinate][]Coordinate, start Coordinate, demoGarden map[Coordinate][]Coordinate, demoStart Coordinate) {
	log.Default().Printf("P1 (demo): %d", dfs(demoGarden, demoStart, 6))
	log.Default().Printf("P1: %d", dfs(garden, start, 64))
}

func debug(test []Coordinate) {
	for y := 0; y < 11; y++ {
		for x := 0; x < 11; x++ {
			if sliceutils.FindIndexOf(test, Coordinate{x, y}) > 0 {
				fmt.Print("O")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println("")
	}
}

func part2() {
	log.Default().Printf("P2: %d", 0)
}

func dfs(garden map[Coordinate][]Coordinate, start Coordinate, steps int) int {
	queue := []Coordinate{start}

	for i := 0; i < steps; i++ {
		newQueue := []Coordinate{}

		for len(queue) > 0 {
			current := queue[0]
			queue = queue[1:]
			newQueue = append(newQueue, garden[current]...)
		}

		queue = append(queue, sliceutils.Unique(newQueue)...)
	}

	return len(queue)
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
