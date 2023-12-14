package main

import (
	"adventofcode2023/util"
	"log"
	"strings"
	"time"
)

// https://adventofcode.com/2023/day/12
func main() {
	conditionRecords := util.ScannerToStringSlice(*util.ReadFile("C:\\Users\\janss\\git\\adventofcode2023\\day12\\demo_input.txt"))

	defer util.TimeTrack(time.Now(), "main")

	log.Default().Printf("P1: %d", part1(conditionRecords))
	log.Default().Printf("P2: %d", part2(conditionRecords))
}

func part1(conditionRecords []string) int {
	return calculateSumOfArrangements(conditionRecords, 1)
}

func calculatePossibleArrangements(conditionRecord string, fold int) int {
	conditionRecordSplit := strings.Split(conditionRecord, " ")
	record := conditionRecordSplit[0]
	info := conditionRecordSplit[1]

	foldedRecord, _ := strings.CutSuffix(strings.Repeat(record+"?", fold), "?")
	foldedInfo, _ := strings.CutSuffix(strings.Repeat(info+",", fold), ",")

	infoInts := util.SpaceDelimitedStringToIntSlice(strings.ReplaceAll(foldedInfo, ",", " "))

	possibleArrangements := findPossibleArrangements([]byte(foldedRecord), 0, infoInts)
	log.Printf("%s produced %d arrangements", conditionRecord, possibleArrangements)
	return possibleArrangements
}

var memo = make(map[string]int)

func findPossibleArrangements(s []byte, index int, info []int) int {
	key := string(s)
	if result, found := memo[key]; found {
		return result
	}

	if index == len(s) {
		if checkPossibleArrangement(string(s), info) {
			return 1
		} else {
			return 0
		}
	}

	possibleArrangements := 0
	if s[index] == '?' {
		s[index] = '#'
		possibleArrangements += findPossibleArrangements(s, index+1, info)
		s[index] = '.'
		possibleArrangements += findPossibleArrangements(s, index+1, info)
		s[index] = '?' // revert back to the original state
	} else {
		possibleArrangements += findPossibleArrangements(s, index+1, info)
	}

	memo[key] = possibleArrangements

	return possibleArrangements
}

func checkPossibleArrangement(arrangement string, info []int) bool {
	matchIndex := 0
	groupSize := 0
	for i := range arrangement {
		if arrangement[i] == '#' {
			groupSize++
		} else {
			if groupSize > 0 && matchIndex == len(info) {
				return false
			} else if groupSize > 0 && info[matchIndex] == groupSize {
				groupSize = 0
				matchIndex++
			} else if groupSize > 0 && info[matchIndex] != groupSize {
				return false
			}
		}
	}
	return (matchIndex == len(info) && groupSize == 0) || (matchIndex == len(info)-1 && info[matchIndex] == groupSize)
}

func calculateSumOfArrangements(conditionRecords []string, fold int) int {
	sumOfArrangements := 0
	for i := range conditionRecords {
		test := calculatePossibleArrangements(conditionRecords[i], fold)
		log.Print(test)
		sumOfArrangements += test
	}
	return sumOfArrangements
}

func part2(conditionRecords []string) int {
	return calculateSumOfArrangements(conditionRecords, 5)
}
