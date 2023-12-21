package main

import (
	"adventofcode2023/util"
	"log"
	"math"
	"time"

	"gonum.org/v1/gonum/graph/simple"
)

// https://adventofcode.com/2023/day/17
func main() {
	// contraption := util.ScannerToStringSlice(*util.ReadFile("C:\\Users\\janss\\git\\adventofcode2023\\day16\\input.txt"))
	defer util.TimeTrack(time.Now(), "main")

	log.Default().Printf("P1: %d", part1())
	log.Default().Printf("P2: %d", part2())
}

func part1() int {
	edges := []simple.WeightedEdge{
		{F: simple.Node('a'), T: simple.Node('f'), W: -1},
		{F: simple.Node('b'), T: simple.Node('a'), W: 1},
		{F: simple.Node('b'), T: simple.Node('c'), W: -1},
		{F: simple.Node('b'), T: simple.Node('d'), W: 1},
		{F: simple.Node('c'), T: simple.Node('b'), W: 0},
		{F: simple.Node('e'), T: simple.Node('a'), W: 1},
		{F: simple.Node('f'), T: simple.Node('e'), W: -1},
	}
	g := simple.NewWeightedDirectedGraph(0, math.Inf(1))
	for _, e := range edges {
		g.SetWeightedEdge(e)
	}

	return 0
}

func part2() int {
	return 0
}
