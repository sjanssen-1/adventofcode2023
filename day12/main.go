package main

import (
	"adventofcode2023/util"
	"log"
	"strings"
	"time"
)

// https://adventofcode.com/2023/day/12
func main() {
	conditionRecords := util.ScannerToStringSlice(*util.ReadFile("/Users/stefjanssens/git/adventofcode2023/day12/input.txt"))

	defer util.TimeTrack(time.Now(), "main")

	log.Default().Printf("P1: %d", part1(conditionRecords))
	log.Default().Printf("P2: %d", part2())
}

func part1(conditionRecords []string) int {
	sumOfArrangements := 0
	for i := range conditionRecords {
		sumOfArrangements += calculatePossibleArrangements(conditionRecords[i])
	}
	return sumOfArrangements
}

func calculatePossibleArrangements(conditionRecord string) int {
	conditionRecordSplit := strings.Split(conditionRecord, " ")
	record := conditionRecordSplit[0]
	info := util.SpaceDelimitedStringToIntSlice(strings.ReplaceAll(conditionRecordSplit[1], ",", " "))

	possibleArrangements := 0
	generatePermutations([]byte(record), 0, &possibleArrangements, info)
	log.Printf("%s produced %d arrangements", conditionRecord, possibleArrangements)
	return possibleArrangements
}

func generatePermutations(s []byte, index int, count *int, info []int) {
	if index == len(s) {
		// fmt.Println(string(s))
		if checkPossibleArrangement(string(s), info) {
			log.Printf("%s is true", string(s))
			*count++
		}
		return
	}

	if s[index] == '?' {
		s[index] = '#'
		generatePermutations(s, index+1, count, info)
		s[index] = '.'
		generatePermutations(s, index+1, count, info)
		s[index] = '?' // revert back to the original state
	} else {
		generatePermutations(s, index+1, count, info)
	}
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

func part2() int {
	return 0
}
