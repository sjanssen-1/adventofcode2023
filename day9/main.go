package main

import (
	"adventofcode2023/util"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/Goldziher/go-utils/sliceutils"
)

// https://adventofcode.com/2023/day/9
func main() {
	historiesScanner := util.ReadFile("C:\\Users\\janss\\git\\adventofcode2023\\day9\\input.txt")
	histories := make([]string, 0)
	for historiesScanner.Scan() {
		histories = append(histories, historiesScanner.Text())
	}

	defer util.TimeTrack(time.Now(), "main")

	log.Default().Printf("P1: %d", part1(histories))
	log.Default().Printf("P2: %d", part2(histories))
}

func parseHistory(history string) []int {
	splitHistory := strings.Split(history, " ")
	historyInt := make([]int, len(splitHistory))
	for i, s := range splitHistory {
		historyInt[i], _ = strconv.Atoi(s)
	}
	return historyInt
}

func diffHistory(history string) [][]int {
	parsedHistory := parseHistory(history)

	sequence := make([][]int, 0)
	sequence = append(sequence, parsedHistory)

	for !is0History(sequence[len(sequence)-1]) {
		diffedHistory := make([]int, len(sequence[len(sequence)-1])-1)
		for i := 0; i < len(diffedHistory); i++ {
			diffedHistory[i] = sequence[len(sequence)-1][i+1] - sequence[len(sequence)-1][i]
		}
		sequence = append(sequence, diffedHistory)
	}
	return sequence
}

func is0History(history []int) bool {
	return sliceutils.Every(history, func(value int, index int, slice []int) bool {
		return value == 0
	})
}

func extrapolate(sequence [][]int) int {
	toIncrement := 0
	for i := len(sequence) - 1; i > 0; i-- {
		toIncrement += sequence[i-1][len(sequence[i-1])-1]
	}
	return toIncrement
}

func extrapolateBackwards(sequence [][]int) int {
	toDecrement := 0
	for i := len(sequence) - 1; i > 0; i-- {
		toDecrement = sequence[i-1][0] - toDecrement
	}
	return toDecrement
}

func part1(histories []string) int {
	extrapolatedValues := make([]int, 0)
	for _, history := range histories {
		diffedHistory := diffHistory(history)
		extrapolatedValues = append(extrapolatedValues, extrapolate(diffedHistory))
	}
	return sliceutils.Sum(extrapolatedValues)
}

func part2(histories []string) int {
	extrapolatedValues := make([]int, 0)
	for _, history := range histories {
		diffedHistory := diffHistory(history)
		extrapolatedValues = append(extrapolatedValues, extrapolateBackwards(diffedHistory))
	}
	return sliceutils.Sum(extrapolatedValues)
}
