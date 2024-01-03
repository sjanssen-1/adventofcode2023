package main

import (
	"adventofcode2023/util"
	"log"
	"slices"
	"time"
)

// https://adventofcode.com/2023/day/23
func main() {
	trailMapDemo := util.ScannerToStringSlice(*util.ReadFile("C:\\Users\\janss\\git\\adventofcode2023\\day23\\demo_input.txt"))
	trailMap := util.ScannerToStringSlice(*util.ReadFile("C:\\Users\\janss\\git\\adventofcode2023\\day23\\input.txt"))
	defer util.TimeTrack(time.Now(), "main")

	demoGraph := parseGraph(trailMapDemo)
	graph := parseGraph(trailMap)

	demoGraphNoIce := parseGraphNoIce(trailMapDemo)
	graphNoIce := parseGraphNoIce(trailMap)

	log.Default().Printf("P1 demo: %d", dfs(demoGraph, Point{1, 0}, Point{21, 22}))
	log.Default().Printf("P1: %d", dfs(graph, Point{1, 0}, Point{139, 140}))
	log.Default().Printf("P2 demo: %d", dfs2(demoGraphNoIce, Point{1, 0}, Point{21, 22}))
	log.Default().Printf("P2: %d", dfs2(graphNoIce, Point{1, 0}, Point{139, 140}))
}

func parseGraph(trailMap []string) map[Point][]Edge {
	graph := map[Point][]Edge{}

	toVisit := []Point{{1, 0}}

	for len(toVisit) > 0 {
		current := toVisit[0]
		toVisit = toVisit[1:]

		leftPath := current.x-1 >= 0 && trailMap[current.y][current.x-1] != '#' && trailMap[current.y][current.x-1] != '>' && trailMap[current.y][current.x] != '>'
		rightPath := current.x+1 < len(trailMap[0]) && trailMap[current.y][current.x+1] != '#'
		upPath := current.y-1 >= 0 && trailMap[current.y-1][current.x] != '#' && trailMap[current.y-1][current.x] != 'v' && trailMap[current.y][current.x] != 'v'
		downPath := current.y+1 < len(trailMap) && trailMap[current.y+1][current.x] != '#'

		edges := make([]Edge, 0)

		if leftPath {
			leftPoint := Point{current.x - 1, current.y}
			if _, exists := graph[leftPoint]; !exists {
				edges = append(edges, Edge{leftPoint})
				toVisit = append(toVisit, leftPoint)
			}
		}

		if rightPath {
			rightPoint := Point{current.x + 1, current.y}
			if _, exists := graph[rightPoint]; !exists {
				edges = append(edges, Edge{rightPoint})
				toVisit = append(toVisit, rightPoint)
			} else if trailMap[current.y][current.x] == '>' {
				edges = append(edges, Edge{rightPoint})
			}
		}

		if upPath {
			upPoint := Point{current.x, current.y - 1}
			if _, exists := graph[upPoint]; !exists {
				edges = append(edges, Edge{upPoint})
				toVisit = append(toVisit, upPoint)
			}
		}

		if downPath {
			downPoint := Point{current.x, current.y + 1}
			if _, exists := graph[downPoint]; !exists {
				edges = append(edges, Edge{downPoint})
				toVisit = append(toVisit, downPoint)
			} else if trailMap[current.y][current.x] == 'v' {
				edges = append(edges, Edge{downPoint})
			}
		}

		graph[current] = edges
	}

	return graph
}

func parseGraphNoIce(trailMap []string) map[Point][]Edge {
	graph := map[Point][]Edge{}

	toVisit := []Point{{1, 0}}

	for len(toVisit) > 0 {
		current := toVisit[0]
		toVisit = toVisit[1:]

		leftPath := current.x-1 >= 0 && trailMap[current.y][current.x-1] != '#'
		rightPath := current.x+1 < len(trailMap[0]) && trailMap[current.y][current.x+1] != '#'
		upPath := current.y-1 >= 0 && trailMap[current.y-1][current.x] != '#'
		downPath := current.y+1 < len(trailMap) && trailMap[current.y+1][current.x] != '#'

		edges := make([]Edge, 0)

		if leftPath {
			leftPoint := Point{current.x - 1, current.y}
			edges = append(edges, Edge{leftPoint})
			if _, exists := graph[leftPoint]; !exists {
				toVisit = append(toVisit, leftPoint)
			}
		}

		if rightPath {
			rightPoint := Point{current.x + 1, current.y}
			edges = append(edges, Edge{rightPoint})
			if _, exists := graph[rightPoint]; !exists {
				toVisit = append(toVisit, rightPoint)
			}
		}

		if upPath {
			upPoint := Point{current.x, current.y - 1}
			edges = append(edges, Edge{upPoint})
			if _, exists := graph[upPoint]; !exists {
				toVisit = append(toVisit, upPoint)
			}
		}

		if downPath {
			downPoint := Point{current.x, current.y + 1}
			edges = append(edges, Edge{downPoint})
			if _, exists := graph[downPoint]; !exists {
				toVisit = append(toVisit, downPoint)
			}
		}

		graph[current] = edges
	}

	return graph
}

func dfs(graph map[Point][]Edge, start, end Point) int {
	toVisit := []Point{start}
	visited := map[Point]int{}
	visited[start] = 0

	for len(toVisit) > 0 {
		current := toVisit[0]
		// log.Print(current)
		toVisit = toVisit[1:]

		edges := graph[current]
		for _, edge := range edges {
			visited[edge.to] = max(visited[current]+1, visited[edge.to])
			if edge.to == end {
				log.Printf("End found %d", visited[edge.to])
			}
			toVisit = slices.Insert(toVisit, 0, edge.to)
		}

	}

	return visited[end]
}

var seen []Point = []Point{}

func dfs2(graph map[Point][]Edge, start, end Point) int {
	if start == end {
		return 0
	}

	m := -1
	seen = append(seen, start)

	for _, edge := range graph[start] {
		if !slices.Contains(seen, edge.to) {
			m = max(m, dfs2(graph, edge.to, end)+1)
		}
	}

	seen = slices.DeleteFunc(seen, func(p Point) bool {
		return p == start
	})
	return m

}

type Edge struct {
	to Point
}

type Point struct {
	x, y int
}
