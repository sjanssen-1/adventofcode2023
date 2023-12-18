package main

import (
	"adventofcode2023/util"
	"fmt"
	"log"
	"time"
)

// https://adventofcode.com/2023/day/14
func main() {
	dish := util.ScannerToStringSlice(*util.ReadFile("/Users/stefjanssens/git/adventofcode2023/day14/input.txt"))
	parsedDish := parseDish(dish)
	parsedDish2 := parseDish(dish)

	defer util.TimeTrack(time.Now(), "main")

	log.Default().Printf("P1: %d", part1(parsedDish))
	log.Default().Printf("P2: %d", part2(parsedDish2))
}

func parseDish(dish []string) [][]byte {
	parsedDish := make([][]byte, 0)

	for y := range dish {
		line := make([]byte, 0)
		for x := range dish[0] {
			line = append(line, dish[y][x])
		}
		parsedDish = append(parsedDish, line)
	}
	return parsedDish
}

var roundRock byte = 79
var empty byte = 46

func tiltNorth(dish *[][]byte) {
	for y := 1; y < len(*dish); y++ {
		for x := 0; x < len((*dish)[0]); x++ {
			if (*dish)[y][x] == roundRock {
				for m := y - 1; m >= 0; m-- {
					if (*dish)[m][x] == empty {
						(*dish)[m][x] = roundRock // set current
						(*dish)[m+1][x] = empty   // set previous
					} else {
						break
					}
				}
			}
		}
	}
}

func tiltEast(dish *[][]byte) {
	for x := len((*dish)[0]) - 2; x >= 0; x-- {
		for y := 0; y < len((*dish)); y++ {
			if (*dish)[y][x] == roundRock {
				for m := x + 1; m < len((*dish)[0]); m++ {
					if (*dish)[y][m] == empty {
						(*dish)[y][m] = roundRock // move rock
						(*dish)[y][m-1] = empty   //
					} else {
						break
					}
				}
			}
		}
	}
}

func tiltSouth(dish *[][]byte) {
	for y := len(*dish) - 2; y >= 0; y-- {
		for x := 0; x < len((*dish)[0]); x++ {
			if (*dish)[y][x] == roundRock {
				for m := y + 1; m < len((*dish)); m++ {
					if (*dish)[m][x] == empty {
						(*dish)[m][x] = roundRock // set current
						(*dish)[m-1][x] = empty   // set previous
					} else {
						break
					}
				}
			}
		}
	}
}

func tiltWest(dish *[][]byte) {
	for x := 1; x < len((*dish)[0]); x++ {
		for y := 0; y < len((*dish)); y++ {
			if (*dish)[y][x] == roundRock {
				for m := x - 1; m >= 0; m-- {
					if (*dish)[y][m] == empty {
						(*dish)[y][m] = roundRock // move rock
						(*dish)[y][m+1] = empty   //
					} else {
						break
					}
				}
			}
		}
	}
}
func calculateLoad(dish [][]byte) int {
	load := 0
	for y := 0; y < len(dish); y++ {
		for x := 0; x < len(dish[0]); x++ {
			if dish[y][x] == roundRock {
				load += len(dish) - y
			}
		}
	}
	return load
}

func part1(dish [][]byte) int {
	tiltNorth(&dish)

	return calculateLoad(dish)
}

func debug(dish [][]byte) {
	for y := 0; y < len(dish); y++ {
		for x := 0; x < len(dish[0]); x++ {
			fmt.Print(string(dish[y][x]))
		}
		fmt.Println("")
	}
	fmt.Println("---")
}

func part2(dish [][]byte) int {
	calculations := make([]int, 0)
	for i := 0; i < 1000; i++ {
		tiltNorth(&dish)
		tiltWest(&dish)
		tiltSouth(&dish)
		tiltEast(&dish)

		calculations = append(calculations, calculateLoad(dish))
	}

	turtle, hare := 0, 0
	for {
		if turtle != 0 && calculations[turtle] == calculations[hare] {
			log.Printf("Cycle after %d", turtle)
			break
		} else {
			turtle++
			hare += 2
		}
	}
	totalIterations := 1000000000
	remaining := totalIterations - (turtle + 1)
	cycle := calculations[turtle:hare]
	index := remaining % (hare - turtle)
	return cycle[index]
}
