package main

import (
	"adventofcode2023/util"
	"log"
	"strings"
	"time"

	"github.com/Goldziher/go-utils/maputils"
	"github.com/Goldziher/go-utils/sliceutils"
)

// https://adventofcode.com/2023/day/8
func main() {
	maps := util.ReadFile("C:\\Users\\janss\\git\\adventofcode2023\\day8\\input.txt")
	defer util.TimeTrack(time.Now(), "main")

	maps.Scan()
	path := maps.Text()

	// empty line
	maps.Scan()

	nodes := make(map[string]Node, 0)
	for maps.Scan() {
		node := maps.Text()
		split := strings.Split(node, " = ")
		nodes[split[0]] = fromString(split[1])
	}

	log.Default().Printf("P1: %d", part1(path, nodes))
	log.Default().Printf("P2: %d", part2(path, nodes))
}

type Node struct {
	left  string
	right string
}

func fromString(node string) Node {
	split := strings.Split(strings.ReplaceAll(strings.ReplaceAll(node, "(", ""), ")", ""), ", ")
	return Node{
		left:  split[0],
		right: split[1],
	}
}

type stopCondition func(string) bool

func traversePath(currentNode string, steps int, path string, nodes map[string]Node, shouldStop stopCondition) int {
	for _, turn := range path {
		if turn == rune('L') {
			currentNode = nodes[currentNode].left
		} else {
			currentNode = nodes[currentNode].right
		}
		steps++

		if shouldStop(currentNode) {
			return steps
		}
	}
	return traversePath(currentNode, steps, path, nodes, shouldStop)
}

func part1(path string, nodes map[string]Node) int {
	return traversePath("AAA", 0, path, nodes, func(node string) bool {
		return node == "ZZZ"
	})
}

func part2(path string, nodes map[string]Node) int {
	currentNodes := sliceutils.Filter(maputils.Keys(nodes), func(value string, index int, slice []string) bool {
		return strings.HasSuffix(value, "A")
	})

	steps := sliceutils.Map(currentNodes, func(value string, index int, slice []string) int {
		return traversePath(value, 0, path, nodes, func(node string) bool {
			return strings.HasSuffix(node, "Z")
		})
	})

	return LCM(steps[0], steps[1], steps[2:])
}

// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers []int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[0], integers[1:])
	}

	return result
}
