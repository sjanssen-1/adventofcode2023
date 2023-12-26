package main

import (
	"adventofcode2023/util"
	"fmt"
	"log"
	"slices"
	"strconv"
	"strings"
	"time"

	"github.com/Goldziher/go-utils/sliceutils"
)

// https://adventofcode.com/2023/day/18
func main() {
	digInstructions := util.ScannerToStringSlice(*util.ReadFile("C:\\Users\\janss\\git\\adventofcode2023\\day18\\input.txt"))
	defer util.TimeTrack(time.Now(), "main")

	parsedDigInstructions := sliceutils.Map(digInstructions, func(value string, index int, slice []string) DigInstruction {
		return NewDigInstruction(value)
	})

	parsedDigInstructions2 := sliceutils.Map(digInstructions, func(value string, index int, slice []string) DigInstruction {
		return NewDigInstruction2(value)
	})

	log.Default().Printf("P1: %d", part1(parsedDigInstructions))
	log.Default().Printf("P2: %d", part2(parsedDigInstructions2))
}

type DigInstruction struct {
	direction string
	amount    int
}

type Point struct {
	x, y int
}

func NewDigInstruction(digInstruction string) DigInstruction {
	split := strings.Split(digInstruction, " ")
	amount, _ := strconv.Atoi(split[1])
	return DigInstruction{split[0], amount}
}

func NewDigInstruction2(digInstruction string) DigInstruction {
	split := strings.Split(digInstruction, " ")
	encoded := split[2]

	amount, _ := strconv.ParseInt(encoded[2:len(encoded)-2], 16, 64)
	switch encoded[len(encoded)-2 : len(encoded)-1] {
	case "0":
		return DigInstruction{"R", int(amount)}
	case "1":
		return DigInstruction{"D", int(amount)}
	case "2":
		return DigInstruction{"L", int(amount)}
	case "3":
		return DigInstruction{"U", int(amount)}
	}
	panic("Weird direction")
}

func calculateArea(digInstructions []DigInstruction) int {
	vertices := make([]Point, 0)
	currentX := 0
	currentY := 0
	boundary := 0

	vertices = append(vertices, Point{0, 0})

	for _, digInstruction := range digInstructions {
		switch digInstruction.direction {
		case "R":
			currentX += digInstruction.amount
		case "L":
			currentX -= digInstruction.amount
		case "U":
			currentY -= digInstruction.amount
		case "D":
			currentY += digInstruction.amount
		}
		vertices = append(vertices, Point{currentX, currentY})
		boundary += digInstruction.amount
	}

	return picksTheorem(shoelace(vertices), boundary) + boundary
}

func debug(vertices []Point) {
	i := 0
	for y := -500; y < 500; y++ {
		for x := -500; x < 500; x++ {
			if slices.ContainsFunc(vertices, func(p Point) bool {
				return p.x == x && p.y == y
			}) {
				i++
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println("")
	}
	fmt.Printf("Amount vs printed %d/%d", i, len(vertices))
}

// https://www.101computing.net/the-shoelace-algorithm/
func shoelace(vertices []Point) int {
	n := len(vertices)
	area := 0

	for i := 0; i < n; i++ {
		j := (i + 1) % n
		area += (vertices[i].x * vertices[j].y) - (vertices[j].x * vertices[i].y)
	}

	return area / 2
}

// https://en.wikipedia.org/wiki/Pick%27s_theorem
func picksTheorem(area int, boundary int) int {
	// A = i + (b/2) - 1
	// i = A - (b/2) + 1
	return area - (boundary / 2) + 1
}

func part1(digInstructions []DigInstruction) int {
	return calculateArea(digInstructions)
}

func part2(digInstructions []DigInstruction) int {
	return calculateArea(digInstructions)
}
