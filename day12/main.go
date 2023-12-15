package main

import (
	"adventofcode2023/util"
	"fmt"
	"log"
	"strings"
	"time"
)

// https://adventofcode.com/2023/day/12
func main() {
	conditionRecords := util.ScannerToStringSlice(*util.ReadFile("C:\\Users\\janss\\git\\adventofcode2023\\day12\\input.txt"))

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

	possibleArrangements := findPossibleArrangements(ConditionRecord{foldedRecord, infoInts})
	// log.Printf("%s %s produced %d arrangements", foldedRecord, foldedInfo, possibleArrangements)
	return possibleArrangements
}

type ConditionRecord struct {
	record string
	groups []int
}

var memo = make(map[string]int)

func createCacheKey(conditionRecord ConditionRecord) string {
	return fmt.Sprintf("%s_%v", conditionRecord.record, conditionRecord.groups)
}

func pound(conditionRecord ConditionRecord, nextGroup int) int {
	var thisGroup string
	if nextGroup >= len(conditionRecord.record) {
		thisGroup = conditionRecord.record[:len(conditionRecord.record)]
	} else {
		thisGroup = conditionRecord.record[:nextGroup]
	}

	thisGroup = strings.ReplaceAll(thisGroup, "?", "#")

	if strings.Count(thisGroup, "#") != nextGroup {
		return 0
	}

	if len(conditionRecord.record) == nextGroup {
		if len(conditionRecord.groups) == 1 {
			return 1
		} else {
			return 0
		}
	}

	if conditionRecord.record[nextGroup] == '.' || conditionRecord.record[nextGroup] == '?' {
		return findPossibleArrangements(ConditionRecord{conditionRecord.record[nextGroup+1:], conditionRecord.groups[1:]})
	}
	return 0
}

func dot(conditionRecord ConditionRecord) int {
	return findPossibleArrangements(ConditionRecord{conditionRecord.record[1:], conditionRecord.groups})
}

func findPossibleArrangements(conditionRecord ConditionRecord) int {
	if result, found := memo[createCacheKey(conditionRecord)]; found {
		// log.Print("CACHE HIT!")
		return result
	}

	record := conditionRecord.record
	groups := conditionRecord.groups

	if len(groups) == 0 {
		if !strings.Contains(string(record), "#") {
			memo[createCacheKey(conditionRecord)] = 1
			return 1
		} else {
			memo[createCacheKey(conditionRecord)] = 0
			return 0
		}
	}

	if len(record) == 0 {
		memo[createCacheKey(conditionRecord)] = 0
		return 0
	}

	nextCharacter := record[0]
	nextGroup := groups[0]

	var out int
	switch nextCharacter {
	case '#':
		out = pound(conditionRecord, nextGroup)
	case '.':
		out = dot(conditionRecord)
	case '?':
		out = dot(conditionRecord) + pound(conditionRecord, nextGroup)
	}
	memo[createCacheKey(conditionRecord)] = out
	return out
}

func calculateSumOfArrangements(conditionRecords []string, fold int) int {
	sumOfArrangements := 0
	for i := range conditionRecords {
		test := calculatePossibleArrangements(conditionRecords[i], fold)
		sumOfArrangements += test
	}
	return sumOfArrangements
}

func part2(conditionRecords []string) int {
	return calculateSumOfArrangements(conditionRecords, 5)
}
