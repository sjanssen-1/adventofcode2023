package main

import (
	"adventofcode2023/util"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/Goldziher/go-utils/sliceutils"
)

// https://adventofcode.com/2023/day/18
func main() {
	digInstructions := util.ScannerToStringSlice(*util.ReadFile("/Users/stefjanssens/git/adventofcode2023/day18/demo_input.txt"))
	defer util.TimeTrack(time.Now(), "main")

	parsedDigInstructions := sliceutils.Map(digInstructions, func(value string, index int, slice []string) DigInstruction {
		return NewDigInstruction(value)
	})

	log.Default().Printf("P1: %d", part1(parsedDigInstructions))
	log.Default().Printf("P2: %d", part2())
}

type DigInstruction struct {
	direction string
	amount    float64
	color     string
}

type Point struct {
	x, y float64
}

func NewDigInstruction(digInstruction string) DigInstruction {
	split := strings.Split(digInstruction, " ")
	amount, _ := strconv.Atoi(split[1])
	return DigInstruction{split[0], float64(amount), split[2]}
}

func calculateArea(digInstructions []DigInstruction) int {
	vertices := make([]Point, 0)
	currentX := 0.0
	currentY := 0.0

	for _, digInstruction := range digInstructions {
		switch digInstruction.direction {
		case "R":
			for x := currentX; x < currentX+digInstruction.amount; x++ {
				vertices = append(vertices, Point{x, currentY})
			}
			currentX += digInstruction.amount
		case "L":
			for x := currentX; x > currentX-digInstruction.amount; x-- {
				vertices = append(vertices, Point{x, currentY})
			}
			currentX -= digInstruction.amount
		case "U":
			for y := currentY; y > currentY-digInstruction.amount; y-- {
				vertices = append(vertices, Point{currentX, y})
			}
			currentY -= digInstruction.amount
		case "D":
			for y := currentY; y < currentY+digInstruction.amount; y++ {
				vertices = append(vertices, Point{currentX, y})
			}
			currentY += digInstruction.amount
		}
	}

	return int(shoelace(vertices))
}

// https://www.101computing.net/the-shoelace-algorithm/
func shoelace(vertices []Point) float64 {
	n := len(vertices)
	area := 0.0

	for i := 0; i < n; i++ {
		j := (i + 1) % n
		area += (vertices[i].x * vertices[j].y) - (vertices[j].x * vertices[i].y)
	}

	area = 0.5 * area
	if area < 0 {
		area = -area
	}

	return area
}

func part1(digInstructions []DigInstruction) int {
	return calculateArea(digInstructions)
}

func part2() int {
	return 0
}
