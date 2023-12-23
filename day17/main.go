package main

import (
	"adventofcode2023/util"
	"log"
	"math"
	"strconv"
	"time"
)

// https://adventofcode.com/2023/day/17
func main() {
	city := util.ScannerToStringSlice(*util.ReadFile("C:\\Users\\janss\\git\\adventofcode2023\\day17\\input.txt"))
	defer util.TimeTrack(time.Now(), "main")

	parsedCity := parseCity(city)

	log.Default().Printf("P1: %d", part1(parsedCity))
	log.Default().Printf("P2: %d", part2(parsedCity))
}

func parseCity(city []string) [][]int {
	cityBlocks := make([][]int, 0)

	for y := range city {
		cityBlockLine := make([]int, 0)
		for x := range city[0] {
			cityBlockValue, _ := strconv.Atoi(string(city[y][x]))
			cityBlockLine = append(cityBlockLine, cityBlockValue)
		}
		cityBlocks = append(cityBlocks, cityBlockLine)
	}

	return cityBlocks

}

type Node struct {
	x, y int
}

type CrucibleState struct {
	pos, dir Node
	steps    int
}

func findLeastHeatLoss(city [][]int, start Node, end Node, minSteps int, maxSteps int) int {
	toCheck := []CrucibleState{{start, Node{1, 0}, 0}, {start, Node{0, 1}, 0}}
	visited := map[CrucibleState]int{{start, Node{0, 0}, 0}: 0}

	heatLoss := math.MaxInt

	for len(toCheck) > 0 {
		current := toCheck[0]
		toCheck = toCheck[1:]

		if current.pos == end && current.steps >= minSteps {
			heatLoss = min(heatLoss, visited[current])
		}

		// visit all neighbours
		for _, dir := range [3]Node{current.dir, turnLeft(current.dir), turnRight(current.dir)} {
			nextNode := Node{current.pos.x + dir.x, current.pos.y + dir.y}
			// check out of bounds
			if nextNode.x < 0 || nextNode.x == len(city[0]) || nextNode.y < 0 || nextNode.y == len(city) {
				continue // skip
			}

			currentHeatLoss := visited[current] + city[nextNode.y][nextNode.x]
			steps := 1
			if dir == current.dir {
				steps = current.steps + 1
			}

			if (dir == current.dir && current.steps < maxSteps) ||
				(dir != current.dir && current.steps >= minSteps) {
				nextState := CrucibleState{nextNode, dir, steps}
				hl, found := visited[nextState]
				if !found || hl > currentHeatLoss {
					visited[nextState] = currentHeatLoss
					toCheck = append(toCheck, nextState)
				}
			}
		}

	}

	return heatLoss
}

func turnLeft(currentDir Node) Node {
	return Node{currentDir.y, -currentDir.x}
}

func turnRight(currentDir Node) Node {
	return Node{-currentDir.y, currentDir.x}
}

func part1(city [][]int) int {
	start, end := Node{0, 0}, Node{len(city[0]) - 1, len(city) - 1}
	return findLeastHeatLoss(city, start, end, 1, 3)
}

func part2(city [][]int) int {
	start, end := Node{0, 0}, Node{len(city[0]) - 1, len(city) - 1}
	return findLeastHeatLoss(city, start, end, 4, 10)
}
