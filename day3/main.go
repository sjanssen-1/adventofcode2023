package main

import (
	"adventofcode2023/util"
	"bufio"
	"log"
	"strconv"
	"time"
	"unicode"

	"github.com/Goldziher/go-utils/sliceutils"
)

// https://adventofcode.com/2023/day/3
func main() {
	engineSchematic := util.ReadFile("./day3/input.txt")
	defer util.TimeTrack(time.Now(), "main")

	engine := parseEngine(*engineSchematic)

	p1Result, foundEngineParts := part1(engine)
	log.Default().Printf("P1: %d", p1Result)
	log.Default().Printf("P2: %d", part2(engine, foundEngineParts))
}

func parseEngine(engineSchematic bufio.Scanner) [][]rune {
	engine := make([][]rune, 0)

	engineLineNumber := 0
	for engineSchematic.Scan() {
		engineLine := engineSchematic.Text()
		engineSlice := make([]rune, 0)
		for _, char := range engineLine {
			engineSlice = append(engineSlice, char)
		}
		engine = append(engine, engineSlice)
		engineLineNumber++
	}
	// log.Printf("%v", engine)
	return engine
}

type EnginePart struct {
	number int
	length int
	x      int
	y      int
}

func part1(engine [][]rune) (int, []EnginePart) {
	foundEngineParts := make([]EnginePart, 0)

	for y, row := range engine {
		var foundNumber string
		var isEnginePart bool
		for x := range row {
			// log.Printf("Current foundNumber = %v", foundNumber)
			// log.Printf("Current character = %v", string(engine[y][x]))

			if unicode.IsDigit(engine[y][x]) {
				foundNumber += string(engine[y][x])

				if !isEnginePart {
					isEnginePart = hasAdjacent(x, y, engine, isSpecial)
				}
			} else if isEnginePart && foundNumber != "" {
				// log.Printf("Found number %v", foundNumber)
				foundNumberInt, _ := strconv.Atoi(foundNumber)
				foundEngineParts = append(foundEngineParts, EnginePart{
					number: foundNumberInt,
					length: len(foundNumber),
					x:      x - len(foundNumber),
					y:      y,
				})
				foundNumber = ""
				isEnginePart = false
			} else {
				// log.Printf("Throwing away number %v", foundNumber)
				foundNumber = ""
				isEnginePart = false
			}
		}

		if isEnginePart && foundNumber != "" {
			// log.Printf("Found number %v", foundNumber)
			foundNumberInt, _ := strconv.Atoi(foundNumber)
			foundEngineParts = append(foundEngineParts, EnginePart{
				number: foundNumberInt,
				length: len(foundNumber),
				x:      len(row) - len(foundNumber),
				y:      y,
			})
		}
	}

	return sliceutils.Reduce(
		foundEngineParts,
		func(acc int, value EnginePart, index int, slice []EnginePart) int {
			return acc + value.number
		},
		0,
	), foundEngineParts
}

func part2(engine [][]rune, engineParts []EnginePart) int {
	// log.Printf("%#v", engineParts)
	foundGearRatios := make([]int, 0)
	for y, row := range engine {
		for x := range row {
			if isGear(engine[y][x]) { // cog x=3 y=1  | x=4 y=1 | x=2 y=2
				// log.Printf("Found gear at x=%d and y=%d", x, y)
				filteredParts := sliceutils.Filter(engineParts, func(value EnginePart, index int, slice []EnginePart) bool {
					// return ((y == value.y+1 || y == value.y-1) && (x >= value.x && x <= value.x+value.length-1)) ||
					// 	(y == value.y && ((value.x+value.length-1 == x-1) || (value.x == x+1)))
					return (y == value.y && ((value.x+value.length-1 == x-1) || (value.x == x+1))) ||
						(y == value.y-1 && ((value.x+value.length-1 == x-1) || (value.x+value.length-1 == x) || (value.x == x) || (value.x == x+1) || (value.x == x-1) || (x >= value.x && x <= value.x+value.length-1))) ||
						(y == value.y+1 && ((value.x+value.length-1 == x-1) || (value.x+value.length-1 == x) || (value.x == x) || (value.x == x+1) || (value.x == x-1) || (x >= value.x && x <= value.x+value.length-1)))
				})
				if len(filteredParts) == 2 {
					foundGearRatios = append(foundGearRatios, filteredParts[0].number*filteredParts[1].number)
				} else {
					// log.Printf("Filtered parts: %#v", filteredParts)
				}
			}
		}
	}

	return sliceutils.Reduce(foundGearRatios,
		func(acc int, value int, index int, slice []int) int {
			return acc + value
		},
		0,
	)
}

type checkCharacter func(character rune) bool

func isSpecial(character rune) bool {
	// log.Printf("Is character %s special? => result %v", string(character), !unicode.IsDigit(character) && string(character) != ".")
	return !unicode.IsDigit(character) && string(character) != "."
}

func isGear(character rune) bool {
	return string(character) == "*"
}

func hasAdjacent(x int, y int, engine [][]rune, doCheck checkCharacter) bool {
	// log.Printf("Checking neighbours of character %v", string(engine[y][x]))

	if x == 0 && y == 0 {
		return doCheck(engine[y][x+1]) || // right
			doCheck(engine[y+1][x]) || // down
			doCheck(engine[y+1][x+1]) // down right
	}

	if x == 0 && y == len(engine)-1 {
		return doCheck(engine[y-1][x]) || // up
			doCheck(engine[y-1][x+1]) || // up right
			doCheck(engine[y][x+1]) // right
	}

	if x == len(engine[0])-1 && y == 0 {
		return doCheck(engine[y][x-1]) || // left
			doCheck(engine[y+1][x-1]) || // down left
			doCheck(engine[y+1][x]) // down
	}

	if x == len(engine[0])-1 && y == len(engine)-1 {
		return doCheck(engine[y][x-1]) || // left
			doCheck(engine[y-1][x]) || // up
			doCheck(engine[y-1][x-1]) // up left
	}

	if x == 0 {
		return doCheck(engine[y+1][x]) || // down
			doCheck(engine[y-1][x]) || // up
			doCheck(engine[y-1][x+1]) || // up right
			doCheck(engine[y][x+1]) || // right
			doCheck(engine[y+1][x+1]) // down right
	}

	if y == 0 {
		return doCheck(engine[y][x-1]) || // left
			doCheck(engine[y][x+1]) || // right
			doCheck(engine[y+1][x-1]) || // down left
			doCheck(engine[y+1][x]) || // down
			doCheck(engine[y+1][x+1]) // down right
	}

	if x == len(engine[0])-1 {
		return doCheck(engine[y-1][x]) || // up
			doCheck(engine[y+1][x]) || // down
			doCheck(engine[y][x-1]) || // left
			doCheck(engine[y-1][x-1]) || // up left
			doCheck(engine[y+1][x-1]) // down left
	}

	if y == len(engine)-1 {
		return doCheck(engine[y][x-1]) || // left
			doCheck(engine[y][x+1]) || // right
			doCheck(engine[y-1][x]) || // up
			doCheck(engine[y-1][x-1]) || // up left
			doCheck(engine[y-1][x+1]) // up right
	}

	return doCheck(engine[y][x+1]) || // right
		doCheck(engine[y][x-1]) || // left
		doCheck(engine[y-1][x]) || // up
		doCheck(engine[y+1][x]) || // down
		doCheck(engine[y-1][x-1]) || // up left
		doCheck(engine[y+1][x-1]) || // down left
		doCheck(engine[y-1][x+1]) || // up right
		doCheck(engine[y+1][x+1]) // down right
}
