package main

import (
	"adventofcode2023/util"
	"log"
	"math"
	"strconv"
	"strings"
	"time"
	"unicode"

	"github.com/Goldziher/go-utils/maputils"
	"github.com/Goldziher/go-utils/sliceutils"
)

// https://adventofcode.com/2023/day/1
func main() {
	defer util.TimeTrack(time.Now(), "main")

	fileScanner := util.ReadFile("./day1/demo_input.txt")

	var calibrationValuesP1 []int
	var calibrationValuesP2 []int

	for fileScanner.Scan() {
		if fileScanner.Text() != "" {
			calibrationLine := fileScanner.Text()
			// log.Default().Printf("Calibration line: %s", calibrationLine)

			calibrationValuesP1 = append(calibrationValuesP1, part1(calibrationLine))
			calibrationValuesP2 = append(calibrationValuesP2, part2(calibrationLine))
		}
	}

	log.Default().Printf("P1: %d", sliceutils.Sum(calibrationValuesP1))
	log.Default().Printf("P2: %d", sliceutils.Sum(calibrationValuesP2))
}

func part1(calibrationLine string) int {
	var firstNumber string
	var lastNumber string

	for _, char := range calibrationLine {
		// log.Default().Print(string(char))
		if unicode.IsDigit(char) {
			firstNumber = string(char)
			break
		}
	}

	for _, char := range sliceutils.Reverse([]rune(calibrationLine)) {
		if unicode.IsDigit(char) {
			lastNumber = string(char)
			break
		}
	}

	calibrationValue, _ := strconv.Atoi(firstNumber + lastNumber)
	return calibrationValue
}

var digits = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func part2(calibrationLine string) int {
	var firstStringDigit int
	var lastStringDigit int

	var lowestStringDigitIndex int = math.MaxInt
	var highestStringDigitIndex int = -1

	for _, digit := range maputils.Keys(digits) {
		firstFoundIndex := strings.Index(calibrationLine, digit)
		if firstFoundIndex != -1 && firstFoundIndex < lowestStringDigitIndex {
			firstStringDigit = digits[digit]
			lowestStringDigitIndex = firstFoundIndex
		}

		lastFoundIndex := strings.LastIndex(calibrationLine, digit)
		if lastFoundIndex != -1 && lastFoundIndex > highestStringDigitIndex {
			lastStringDigit = digits[digit]
			highestStringDigitIndex = lastFoundIndex
		}
	}

	var firstNumber int
	var lastNumber int

	for i, char := range calibrationLine {
		// log.Default().Print(string(char))
		if unicode.IsDigit(char) && i < lowestStringDigitIndex {
			firstNumber, _ = strconv.Atoi(string(char))
			break
		}
	}

	for i, char := range sliceutils.Reverse([]rune(calibrationLine)) {
		if unicode.IsDigit(char) && len(calibrationLine)-1-i > highestStringDigitIndex {
			lastNumber, _ = strconv.Atoi(string(char))
			break
		}
	}

	var finalFirst int
	var finalLast int

	if firstNumber == 0 {
		finalFirst = firstStringDigit
	} else {
		finalFirst = firstNumber
	}

	if lastNumber == 0 {
		finalLast = lastStringDigit
	} else {
		finalLast = lastNumber
	}

	// log.Default().Printf("%d", finalFirst*10+finalLast)

	return finalFirst*10 + finalLast
}
