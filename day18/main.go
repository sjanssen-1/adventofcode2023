package main

import (
	"adventofcode2023/util"
	"log"
	"strconv"
	"strings"
	"time"
)

// https://adventofcode.com/2023/day/18
func main() {
	digInstructions := util.ScannerToStringSlice(*util.ReadFile("C:\\Users\\janss\\git\\adventofcode2023\\day17\\input.txt"))
	defer util.TimeTrack(time.Now(), "main")

	log.Default().Printf("P1: %d", part1())
	log.Default().Printf("P2: %d", part2())
}

type DigInstruction struct {
	direction string
	amount    int
	color     string
}

func NewDigInstruction(digInstruction string) DigInstruction {
	split := strings.Split(digInstruction, " ")
	amount, _ := strconv.Atoi(split[1])
	return DigInstruction{split[0], amount, split[2]}
}

func part1() int {
	return 0
}

func part2() int {
	return 0
}
