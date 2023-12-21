package main

import (
	"adventofcode2023/util"
	"fmt"
	"log"
	"math"
	"reflect"
	"slices"
	"strings"
	"time"

	"github.com/Goldziher/go-utils/maputils"
)

// https://adventofcode.com/2023/day/16
func main() {
	contraption := util.ScannerToStringSlice(*util.ReadFile("C:\\Users\\janss\\git\\adventofcode2023\\day16\\input.txt"))
	parsedContraption := parseContraption(contraption)
	defer util.TimeTrack(time.Now(), "main")

	log.Default().Printf("P1: %d", part1(parsedContraption))
	log.Default().Printf("P2: %d", part2(parsedContraption))
}

type ContraptionPart struct {
	x, y      int
	energized bool
}

func NewContraptionPart(x int, y int) ContraptionPart {
	return ContraptionPart{x, y, false}
}

type Empty struct {
	ContraptionPart
}

func NewEmpty(x int, y int) Empty {
	return Empty{NewContraptionPart(x, y)}
}

type Mirror struct {
	ContraptionPart
	form byte
}

func NewMirror(x int, y int, form byte) Mirror {
	return Mirror{NewContraptionPart(x, y), form}
}

type Splitter struct {
	ContraptionPart
	form byte
}

func NewSplitter(x int, y int, form byte) Splitter {
	return Splitter{NewContraptionPart(x, y), form}
}

func parseContraption(contraption []string) [][]interface{} {
	parsedContraption := make([][]interface{}, 0)
	for y := range contraption {
		contraptionLine := make([]interface{}, 0)
		for x := range contraption[y] {
			var part interface{}
			switch contraption[y][x] {
			case '.':
				part = NewEmpty(x, y)
			case '/', '\\':
				part = NewMirror(x, y, contraption[y][x])
			case '|', '-':
				part = NewSplitter(x, y, contraption[y][x])
			}
			contraptionLine = append(contraptionLine, part)
		}
		parsedContraption = append(parsedContraption, contraptionLine)
	}
	return parsedContraption
}

func part1(contraption [][]interface{}) int {
	visited = make(map[string]string)
	copy := copyContraption(contraption)
	beam(copy, 0, 0, "east")
	// debug(contraption)
	return countEnergized(copy)
}

var visited map[string]string

func visit(x int, y int, direction string) bool {
	key := string(x) + string(y)
	if slices.Contains(maputils.Keys(visited), key) {
		if strings.Contains(visited[key], direction) {
			return true
		} else {
			visited[key] += direction
			return false
		}
	} else {
		visited[key] = direction
		return false
	}
}

func beam(contraption [][]interface{}, x int, y int, direction string) {
	// debug(contraption)
	contraptionPart := contraption[y][x]

	if visit(x, y, direction) {
		return
	}

	if part, ok := contraptionPart.(Empty); ok {
		energize(&part.ContraptionPart)
		contraption[y][x] = part
	}

	if part, ok := contraptionPart.(Mirror); ok {
		energize(&part.ContraptionPart)
		contraption[y][x] = part
	}

	if part, ok := contraptionPart.(Splitter); ok {
		energize(&part.ContraptionPart)
		contraption[y][x] = part
	}

	switch reflect.TypeOf(contraptionPart) {
	case reflect.TypeOf(Empty{}):
		switch direction {
		case "north":
			if y-1 < 0 {
				return
			}
			beam(contraption, x, y-1, "north")
		case "south":
			if y+1 == len(contraption) {
				return
			}
			beam(contraption, x, y+1, "south")
		case "east":
			if x+1 == len(contraption[0]) {
				return
			}
			beam(contraption, x+1, y, "east")
		case "west":
			if x-1 < 0 {
				return
			}
			beam(contraption, x-1, y, "west")
		}
	case reflect.TypeOf(Mirror{}):
		switch contraptionPart.(Mirror).form {
		case '/':
			switch direction {
			case "north":
				if x+1 == len(contraption[0]) {
					return
				}
				beam(contraption, x+1, y, "east")
			case "south":
				if x-1 < 0 {
					return
				}
				beam(contraption, x-1, y, "west")
			case "east":
				if y-1 < 0 {
					return
				}
				beam(contraption, x, y-1, "north")
			case "west":
				if y+1 == len(contraption) {
					return
				}
				beam(contraption, x, y+1, "south")
			}
		case '\\':
			switch direction {
			case "north":
				if x-1 < 0 {
					return
				}
				beam(contraption, x-1, y, "west")
			case "south":
				if x+1 == len(contraption[0]) {
					return
				}
				beam(contraption, x+1, y, "east")
			case "west":
				if y-1 < 0 {
					return
				}
				beam(contraption, x, y-1, "north")
			case "east":
				if y+1 == len(contraption) {
					return
				}
				beam(contraption, x, y+1, "south")
			}
		}
	case reflect.TypeOf(Splitter{}):
		switch contraptionPart.(Splitter).form {
		case '|':
			switch direction {
			case "north":
				if y-1 < 0 {
					return
				}
				beam(contraption, x, y-1, "north")
			case "south":
				if y+1 == len(contraption) {
					return
				}
				beam(contraption, x, y+1, "south")
			case "east", "west":
				if y-1 >= 0 {
					beam(contraption, x, y-1, "north")
				}
				if y+1 < len(contraption) {
					beam(contraption, x, y+1, "south")
				}
			}
		case '-':
			switch direction {
			case "east":
				if x+1 == len(contraption[0]) {
					return
				}
				beam(contraption, x+1, y, "east")
			case "west":
				if x-1 < 0 {
					return
				}
				beam(contraption, x-1, y, "west")
			case "north", "south":
				if x-1 >= 0 {
					beam(contraption, x-1, y, "west")
				}
				if x+1 < len(contraption[0]) {
					beam(contraption, x+1, y, "east")
				}
			}
		}
	}
}

func energize(part *ContraptionPart) {
	part.energized = true
}

func isEnergized(part interface{}) bool {
	switch reflect.TypeOf(part) {
	case reflect.TypeOf(Empty{}):
		return part.(Empty).energized
	case reflect.TypeOf(Mirror{}):
		return part.(Mirror).energized
	case reflect.TypeOf(Splitter{}):
		return part.(Splitter).energized
	}
	return false
}

func countEnergized(contraption [][]interface{}) int {
	energized := 0
	for y := range contraption {
		for x := range contraption[y] {
			if isEnergized(contraption[y][x]) {
				energized++
			}
		}
	}
	return energized
}

func debug(contraption [][]interface{}) {
	for y := range contraption {
		for x := range contraption[y] {
			if isEnergized(contraption[y][x]) {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
	fmt.Println("---")
}

func part2(contraption [][]interface{}) int {
	maxEnergized := 0
	for y := range contraption {
		log.Printf("y1: %d", y)
		visited = make(map[string]string)
		copy := copyContraption(contraption)
		beam(copy, 0, y, "east")
		maxEnergized = int(math.Max(float64(maxEnergized), float64(countEnergized(copy))))

		log.Printf("y2: %d", y)
		visited = make(map[string]string)
		copy = copyContraption(contraption)
		beam(copy, len(contraption[0])-1, y, "west")
		maxEnergized = int(math.Max(float64(maxEnergized), float64(countEnergized(copy))))
	}

	for x := range contraption[0] {
		log.Printf("x1: %d", x)
		visited = make(map[string]string)
		copy := copyContraption(contraption)
		beam(copy, x, 0, "south")
		maxEnergized = int(math.Max(float64(maxEnergized), float64(countEnergized(copy))))

		log.Printf("x2: %d", x)
		visited = make(map[string]string)
		copy = copyContraption(contraption)
		beam(copy, x, len(contraption)-1, "north")
		maxEnergized = int(math.Max(float64(maxEnergized), float64(countEnergized(copy))))
	}
	return maxEnergized
}

func copyContraption(original [][]interface{}) [][]interface{} {
	copyArray := make([][]interface{}, len(original))

	for i := range original {
		copyArray[i] = make([]interface{}, len(original[i]))
		copy(copyArray[i], original[i])
	}

	return copyArray
}
