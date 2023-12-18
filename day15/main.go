package main

import (
	"adventofcode2023/util"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/Goldziher/go-utils/sliceutils"
	orderedmap "github.com/wk8/go-ordered-map/v2"
)

// https://adventofcode.com/2023/day/15
func main() {

	scanner := util.ReadFile("/Users/stefjanssens/git/adventofcode2023/day15/input.txt")

	scanner.Scan()
	initializationSequence := scanner.Text()

	defer util.TimeTrack(time.Now(), "main")

	log.Default().Printf("P1: %d", part1(initializationSequence))
	log.Default().Printf("P2: %d", part2(initializationSequence))
}

func calculateHash(s string) int {
	currentValue := 0
	for i := range s {
		currentValue += int(s[i])
		currentValue *= 17
		currentValue %= 256
	}
	return currentValue
}

func part1(initializationSequence string) int {
	return sliceutils.Reduce(strings.Split(initializationSequence, ","), func(acc int, value string, index int, slice []string) int {
		return acc + calculateHash(value)
	}, 0)
}

type Step struct {
	label       string
	operation   byte
	focalLength int
}

func part2(initializationSequence string) int {
	steps := strings.Split(initializationSequence, ",")

	boxes := make(map[int]*orderedmap.OrderedMap[string, int], 0)
	for i := 0; i < 256; i++ {
		boxes[i] = orderedmap.New[string, int]()
	}

	for i := range steps {
		step := parseStep(steps[i])
		boxNumber := calculateHash(step.label)
		if step.operation == '=' {
			_, present := boxes[boxNumber].Set(step.label, step.focalLength)
			if !present {
				boxes[boxNumber].MoveToBack(step.label)
			}
		} else {
			boxes[boxNumber].Delete(step.label)
		}
	}

	focussingPower := 0
	for i := 0; i < 256; i++ {
		if boxes[i].Len() > 0 {
			boxValue := i + 1
			j := 1
			for pair := boxes[i].Oldest(); pair != nil; pair = pair.Next() {
				focussingPower += boxValue * j * pair.Value
				j++
			}
		}
	}
	return focussingPower
}

func parseStep(step string) Step {
	if strings.Contains(step, "=") {
		split := strings.Split(step, "=")
		focalLength, _ := strconv.Atoi(split[1])
		return Step{label: split[0], operation: '=', focalLength: focalLength}
	} else {
		return Step{label: step[:len(step)-1], operation: '-'}
	}
}
