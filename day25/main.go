package main

import (
	"adventofcode2023/util"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/daviddengcn/go-algs/maxflow"
)

// https://adventofcode.com/2023/day/25
func main() {
	wiringDiagramDemo := util.ScannerToStringSlice(*util.ReadFile("C:\\Users\\janss\\git\\adventofcode2023\\day25\\demo_input.txt"))
	// note := util.ScannerToStringSlice(*util.ReadFile("C:\\Users\\janss\\git\\adventofcode2023\\day25\\input.txt"))
	defer util.TimeTrack(time.Now(), "main")

	connectedParts := parseDiagram(wiringDiagramDemo)

	g := maxflow.NewGraph()

	nodes := make([]*maxflow.Node, 2)

	for i := range nodes {
		nodes[i] = g.AddNode()
	} // for i

	g.SetTweights(nodes[0], 1, 5)
	g.SetTweights(nodes[1], 2, 6)
	g.AddEdge(nodes[0], nodes[1], 3, 4)

	g.Run()

	// flow := g.Flow()

	if g.IsSource(nodes[0]) {
		fmt.Println("nodes 0 is SOURCE")
	} else {
		fmt.Println("nodes 0 is SINK")
	} // else

	log.Default().Printf("P1 demo: %d", part1(connectedParts))
	// log.Default().Printf("P1: %d", part1())
	log.Default().Printf("P2 demo: %d", part2())
}

func parseDiagram(diagram []string) map[string][]string {
	connectedParts := map[string][]string{}
	for _, line := range diagram {
		split := strings.Split(line, ": ")
		connectedParts[split[0]] = strings.Split(split[1], " ")
	}

	return connectedParts

	// connectedParts2 := map[string][]string{}
	// for key, values := range connectedParts {
	// 	for _, value := range values {
	// 		connectedParts2[value] = append(connectedParts2[value], key)
	// 	}
	// }

	// return maputils.Merge(connectedParts, connectedParts2)
}

func part1(connectedParts map[string][]string) int {
	// nrOfParts := len(connectedParts)

	return 0
}

func part2() int {
	return 0
}
