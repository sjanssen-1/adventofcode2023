package main

import (
	"adventofcode2023/util"
	"log"
	"reflect"
	"time"
)

// https://adventofcode.com/2023/day/16
func main() {
	contraption := util.ScannerToStringSlice(*util.ReadFile("/Users/stefjanssens/git/adventofcode2023/day16/demo_input.txt"))
	parsedContraption := parseContraption(contraption)
	defer util.TimeTrack(time.Now(), "main")

	log.Default().Printf("P1: %d", part1(parsedContraption))
	log.Default().Printf("P2: %d", part2())
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
	return 0
}

func beam(contraption [][]interface{}, startX int, startY int, startDirection string) {
	x, y := startX, startY
	direction := startDirection

	contraptionPart := contraption[y][x]

	energize(&contraptionPart)

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
				// TODO do the rest and fix out of bounds
				if y-1 < 0 {
					return
				}
				beam(contraption, x+1, y, "east")
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
		case '\\':

		}
	case reflect.TypeOf(Splitter{}):
	}

	switch direction {
	case "north":
		if y-1 < 0 {
			return
		}
		y--
	case "south":
		if y+1 == len(contraption) {
			return
		}
		y++
	case "east":
		if x+1 == len(contraption[0]) {
			return
		}
		x++
	case "west":
		if x-1 < 0 {
			return
		}
		x--
	}

}

func energize(part interface{}) {
	part.(*ContraptionPart).energized = true
}

func part2() int {
	return 0
}
