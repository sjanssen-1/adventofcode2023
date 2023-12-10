package main

import (
	"adventofcode2023/util"
	"fmt"
	"log"
	"reflect"
	"time"
)

// https://adventofcode.com/2023/day/10
func main() {
	sketch := util.ReadFile("C:\\Users\\janss\\git\\adventofcode2023\\day10\\demo_input3.txt")
	tiles := make([][]interface{}, 0)
	line := 0
	x, y := 0, 0
	for sketch.Scan() {
		sketchLine := sketch.Text()
		parsedTiles, animalX, animalY := parseTiles(sketchLine, line)
		if animalX != -1 && animalY != -1 {
			x = animalX
			y = animalY
		}
		tiles = append(tiles, parsedTiles)
		line++
	}

	defer util.TimeTrack(time.Now(), "main")

	// log.Print(tiles)
	// log.Printf("Animal start = %d, %d", x, y)

	log.Default().Printf("P1: %d", part1(tiles, x, y))
	log.Default().Printf("P2: %d", part2(tiles))
}

func parseTiles(sketchLine string, row int) ([]interface{}, int, int) {
	tiles := make([]interface{}, 0)
	x, y := -1, -1
	for i, sketchTile := range sketchLine {
		switch string(sketchTile) {
		case "|":
			tiles = append(tiles, Pipe{
				Tile:       Tile{x: i, y: row, value: "|"},
				direction1: Direction{x: i, y: row + 1}, // south
				direction2: Direction{x: i, y: row - 1}, // north
			})
		case "-":
			tiles = append(tiles, Pipe{
				Tile:       Tile{x: i, y: row, value: "-"},
				direction1: Direction{x: i + 1, y: row}, // east
				direction2: Direction{x: i - 1, y: row}, // west
			})
		case "L":
			tiles = append(tiles, Pipe{
				Tile:       Tile{x: i, y: row, value: "L"},
				direction1: Direction{x: i + 1, y: row}, // east
				direction2: Direction{x: i, y: row - 1}, // north
			})
		case "J":
			tiles = append(tiles, Pipe{
				Tile:       Tile{x: i, y: row, value: "J"},
				direction1: Direction{x: i - 1, y: row}, // west
				direction2: Direction{x: i, y: row - 1}, // north
			})
		case "7":
			tiles = append(tiles, Pipe{
				Tile:       Tile{x: i, y: row, value: "7"},
				direction1: Direction{x: i - 1, y: row}, // west
				direction2: Direction{x: i, y: row + 1}, // south
			})
		case "F":
			tiles = append(tiles, Pipe{
				Tile:       Tile{x: i, y: row, value: "F"},
				direction1: Direction{x: i + 1, y: row}, // east
				direction2: Direction{x: i, y: row + 1}, // south
			})
		case ".":
			tiles = append(tiles, Ground{
				Tile: Tile{x: i, y: row, value: "."},
			})
		case "S":
			x = i
			y = row
			tiles = append(tiles, Animal{
				Tile: Tile{x: i, y: row, value: "S"},
			})
		}
	}
	return tiles, x, y
}

type Tile struct {
	x     int
	y     int
	value string
}

type Ground struct {
	Tile
}

type DummyGround struct {
	Tile
}

type Animal struct {
	Tile
}

type Pipe struct {
	Tile
	direction1 Direction
	direction2 Direction
}

type Direction struct {
	x int
	y int
}

func loop(tiles [][]interface{}, x int, y int) int {
	previousX, previousY := x, y
	currX, currY := findNextAfterStart(tiles, x, y)

	steps := 1

	for currX != x || currY != y {
		currentPipe := tiles[currY][currX].(Pipe)
		var newX, newY int
		if currentPipe.direction1.x == previousX && currentPipe.direction1.y == previousY {
			newX = currentPipe.direction2.x
			newY = currentPipe.direction2.y
		} else {
			newX = currentPipe.direction1.x
			newY = currentPipe.direction1.y
		}

		previousX = currX
		previousY = currY
		currX = newX
		currY = newY

		steps++
	}
	return steps
}

func findNextAfterStart(tiles [][]interface{}, startX int, startY int) (int, int) {
	if startX-1 >= 0 && isConnected(tiles[startY][startX-1], startX, startY) {
		log.Print("left!")
		return tiles[startY][startX-1].(Pipe).x, tiles[startY][startX-1].(Pipe).y
	} else if startX+1 < len(tiles[0]) && isConnected(tiles[startY][startX+1], startX, startY) {
		log.Print("right!")
		return tiles[startY][startX+1].(Pipe).x, tiles[startY][startX+1].(Pipe).y
	} else if startY-1 >= 0 && isConnected(tiles[startY-1][startX], startX, startY) {
		log.Print("up!")
		return tiles[startY-1][startX].(Pipe).x, tiles[startY-1][startX].(Pipe).y
	} else if startY+1 < len(tiles) && isConnected(tiles[startY+1][startX], startX, startY) {
		log.Print("down!")
		return tiles[startY+1][startX].(Pipe).x, tiles[startY+1][startX].(Pipe).y
	}
	panic("No start found...")
}

func isConnected(tile interface{}, x int, y int) bool {
	if reflect.TypeOf(tile) == reflect.TypeOf(Pipe{}) {
		if (tile.(Pipe).direction1.x == x && tile.(Pipe).direction1.y == y) ||
			(tile.(Pipe).direction2.x == x && tile.(Pipe).direction2.y == y) {
			return true
		}
	}
	return false
}

func part1(tiles [][]interface{}, x int, y int) int {
	return loop(tiles, x, y) / 2
}

func part2(tiles [][]interface{}) int {
	debugPrintTiles(tiles)

	expandedTiles := make([][]interface{}, 0)
	for y := range tiles {
		expandedTilesRow := make([]interface{}, 0)
		for x := range tiles[0] {
			expandedTilesRow = append(expandedTilesRow, tiles[y][x])
			if reflect.TypeOf(tiles[y][x]) == reflect.TypeOf(Ground{}) {
				expandedTilesRow = append(expandedTilesRow, DummyGround{Tile{value: "*"}})
			}
		}
		expandedTiles = append(expandedTiles, expandedTilesRow)
	}
	debugPrintTiles(expandedTiles)

	return 0
}

func debugPrintTiles(tiles [][]interface{}) {
	for y := range tiles {
		for x := range tiles[0] {
			if reflect.TypeOf(tiles[y][x]) == reflect.TypeOf(Ground{}) {
				fmt.Print(tiles[y][x].(Ground).value)
			} else if reflect.TypeOf(tiles[y][x]) == reflect.TypeOf(Pipe{}) {
				fmt.Print(tiles[y][x].(Pipe).value)
			} else if reflect.TypeOf(tiles[y][x]) == reflect.TypeOf(DummyGround{}) {
				fmt.Print(tiles[y][x].(DummyGround).value)
			} else {
				fmt.Print(tiles[y][x].(Animal).value)
			}
		}
		fmt.Println("")
	}
}
