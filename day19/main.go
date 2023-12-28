package main

import (
	"adventofcode2023/util"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/Goldziher/go-utils/sliceutils"
)

// https://adventofcode.com/2023/day/19
func main() {
	theSystem := util.ReadFile("C:\\Users\\janss\\git\\adventofcode2023\\day19\\demo_input.txt")
	defer util.TimeTrack(time.Now(), "main")

	workflows := make(map[string]Workflow)
	parts := make([]Part, 0)

	isParts := false
	for theSystem.Scan() {
		systemLine := theSystem.Text()
		if systemLine == "" {
			isParts = true
			continue
		}

		if !isParts {
			split := strings.Split(systemLine, "{")
			workflows[split[0]] = NewWorkFlow(split[1][:len(split[1])-1])
		} else {
			parts = append(parts, NewPart(systemLine))
		}
	}

	// log.Print(workflows)
	// log.Print(parts)

	log.Default().Printf("P1: %d", part1(workflows, parts))
	log.Default().Printf("P2: %d", part2(workflows))
}

type Workflow struct {
	conditions []Condition
}

func NewWorkFlow(s string) Workflow {
	split := strings.Split(s, ",")

	conditions := make([]Condition, 0)

	for _, c := range split {
		csplit := strings.Split(c, ":")
		if len(csplit) == 1 {
			conditions = append(conditions, Condition{next: csplit[0], isFinal: true})
		} else {
			next := csplit[1]
			category := string(csplit[0][0])
			operand := string(csplit[0][1])
			value, _ := strconv.Atoi(csplit[0][2:])
			conditions = append(conditions, Condition{category, operand, value, next, false})
		}
	}
	return Workflow{conditions}
}

type Condition struct {
	category string
	operator string
	value    int
	next     string
	isFinal  bool
}

type Part struct {
	x int
	m int
	a int
	s int
}

func NewPart(part string) Part {
	var x, m, a, s int
	for _, c := range strings.Split(part[1:len(part)-1], ",") {
		csplit := strings.Split(c, "=")
		switch csplit[0] {
		case "x":
			x, _ = strconv.Atoi(csplit[1])
		case "m":
			m, _ = strconv.Atoi(csplit[1])
		case "a":
			a, _ = strconv.Atoi(csplit[1])
		case "s":
			s, _ = strconv.Atoi(csplit[1])
		}
	}
	return Part{x, m, a, s}
}

func getCategory(category string, part Part) int {
	switch category {
	case "x":
		return part.x
	case "m":
		return part.m
	case "a":
		return part.a
	case "s":
		return part.s
	}
	panic("bogus category")
}

func calculatePartScore(part Part) int {
	return part.x + part.m + part.a + part.s
}

func processPart(workflows map[string]Workflow, part Part) int {
	start := workflows["in"]

flows:
	for {
		for _, condition := range start.conditions {
			if condition.isFinal {
				if condition.next == "A" {
					return calculatePartScore(part)
				} else if condition.next == "R" {
					return 0
				} else {
					start = workflows[condition.next]
					continue flows
				}
			}

			var passed bool
			switch condition.operator {
			case ">":
				passed = getCategory(condition.category, part) > condition.value
			case "<":
				passed = getCategory(condition.category, part) < condition.value
			}

			if passed {
				if condition.next == "A" {
					return calculatePartScore(part)
				} else if condition.next == "R" {
					return 0
				} else {
					start = workflows[condition.next]
					continue flows
				}
			}
		}
	}
}

type QueuePart struct {
	name           string
	conditionIndex int
	ranges         Ranges
}

type Ranges struct {
	xRange Range
	mRange Range
	aRange Range
	sRange Range
}

type Range struct {
	low  int
	high int
}

func calculateCombinations(ranges Ranges) int {
	return (ranges.xRange.high - ranges.xRange.low + 1) *
		(ranges.mRange.high - ranges.mRange.low + 1) *
		(ranges.aRange.high - ranges.aRange.low + 1) *
		(ranges.sRange.high - ranges.sRange.low + 1)
}

func isCompletelyInRange(low int, high int, operator string, value int) bool {
	switch operator {
	case ">":
		return low > value
	case "<":
		return high < value
	}

	panic("dumb operator")
}

func processAllParts(workflows map[string]Workflow) int {
	queue := []QueuePart{QueuePart{"in", 0, Ranges{Range{1, 4000}, Range{1, 4000}, Range{1, 4000}, Range{1, 4000}}}}
	combinations := 0

queue:
	for len(queue) > 0 {
		currQueue := queue[:1][0]
		currFlow := workflows[currQueue.name]
		queue = queue[1:]

		for i := currQueue.conditionIndex; i < len(currFlow.conditions); i++ {
			condition := currFlow.conditions[i]

			if condition.isFinal {
				if condition.next == "A" {
					combinations += calculateCombinations(currQueue.ranges)
				} else if condition.next == "R" {
					continue queue
				} else {
					// bug: should continue to the next condition ; remove for loop
					// put work on queue
					continue
				}
			}

			var low int
			var high int

			switch condition.category {
			case "x":
				high = currQueue.ranges.xRange.high
				low = currQueue.ranges.xRange.low
			case "m":
				high = currQueue.ranges.mRange.high
				low = currQueue.ranges.mRange.low
			case "a":
				high = currQueue.ranges.aRange.high
				low = currQueue.ranges.aRange.low
			case "s":
				high = currQueue.ranges.sRange.high
				low = currQueue.ranges.sRange.low
			}

			// TODO figure out how to check if true or not --> check if completely satisfies condition

			if isCompletelyInRange(low, high, condition.operator, condition.value) && condition.next == "A" {
				// calculate the combinations and continue with the queue
				combinations += calculateCombinations(currQueue.ranges)
			} else if isCompletelyInRange(low, high, condition.operator, condition.value) && condition.next == "R" {
				// just continue with the queue
				continue queue
			} else {
				switch condition.category {
				case "x":
					// check the possible splits and put new work on queue on the current condition index
					switch condition.operator {
					case "<":
						// true branch
						queue = append(queue, QueuePart{condition.next, 0, Ranges{Range{currQueue.ranges.xRange.low, condition.value - 1}, currQueue.ranges.mRange, currQueue.ranges.aRange, currQueue.ranges.sRange}})
						// false branch
						queue = append(queue, QueuePart{currQueue.name, currQueue.conditionIndex + 1, Ranges{Range{condition.value, currQueue.ranges.xRange.high}, currQueue.ranges.mRange, currQueue.ranges.aRange, currQueue.ranges.sRange}})
					case ">":
						// false branch
						queue = append(queue, QueuePart{currQueue.name, currQueue.conditionIndex + 1, Ranges{Range{currQueue.ranges.xRange.low, condition.value}, currQueue.ranges.mRange, currQueue.ranges.aRange, currQueue.ranges.sRange}})
						// true branch
						queue = append(queue, QueuePart{condition.next, 0, Ranges{Range{condition.value + 1, currQueue.ranges.xRange.high}, currQueue.ranges.mRange, currQueue.ranges.aRange, currQueue.ranges.sRange}})
					}
					continue queue
				case "m":
					// check the possible splits and put new work on queue on the current condition index
					switch condition.operator {
					case "<":
						// true branch
						queue = append(queue, QueuePart{condition.next, 0, Ranges{currQueue.ranges.xRange, Range{currQueue.ranges.mRange.low, condition.value - 1}, currQueue.ranges.aRange, currQueue.ranges.sRange}})
						// false branch
						queue = append(queue, QueuePart{currQueue.name, currQueue.conditionIndex + 1, Ranges{currQueue.ranges.xRange, Range{condition.value, currQueue.ranges.mRange.high}, currQueue.ranges.aRange, currQueue.ranges.sRange}})
					case ">":
						// false branch
						queue = append(queue, QueuePart{currQueue.name, currQueue.conditionIndex + 1, Ranges{currQueue.ranges.xRange, Range{currQueue.ranges.mRange.low, condition.value}, currQueue.ranges.aRange, currQueue.ranges.sRange}})
						// true branch
						queue = append(queue, QueuePart{condition.next, 0, Ranges{currQueue.ranges.xRange, Range{condition.value + 1, currQueue.ranges.mRange.high}, currQueue.ranges.aRange, currQueue.ranges.sRange}})
					}
					continue queue
				case "a":
					// check the possible splits and put new work on queue on the current condition index
					switch condition.operator {
					case "<":
						// true branch
						queue = append(queue, QueuePart{condition.next, 0, Ranges{currQueue.ranges.xRange, currQueue.ranges.mRange, Range{currQueue.ranges.sRange.low, condition.value - 1}, currQueue.ranges.sRange}})
						// false branch
						queue = append(queue, QueuePart{currQueue.name, currQueue.conditionIndex + 1, Ranges{currQueue.ranges.xRange, currQueue.ranges.mRange, Range{condition.value, currQueue.ranges.sRange.high}, currQueue.ranges.sRange}})
					case ">":
						// false branch
						queue = append(queue, QueuePart{currQueue.name, currQueue.conditionIndex + 1, Ranges{currQueue.ranges.xRange, currQueue.ranges.mRange, Range{currQueue.ranges.aRange.low, condition.value}, currQueue.ranges.sRange}})
						// true branch
						queue = append(queue, QueuePart{condition.next, 0, Ranges{currQueue.ranges.xRange, currQueue.ranges.mRange, Range{condition.value + 1, currQueue.ranges.aRange.high}, currQueue.ranges.sRange}})
					}
					continue queue
				case "s":
					// check the possible splits and put new work on queue on the current condition index
					switch condition.operator {
					case "<":
						// true branch
						queue = append(queue, QueuePart{condition.next, 0, Ranges{currQueue.ranges.xRange, currQueue.ranges.mRange, currQueue.ranges.aRange, Range{currQueue.ranges.sRange.low, condition.value - 1}}})
						// false branch
						queue = append(queue, QueuePart{currQueue.name, currQueue.conditionIndex + 1, Ranges{currQueue.ranges.xRange, currQueue.ranges.mRange, currQueue.ranges.aRange, Range{condition.value, currQueue.ranges.sRange.high}}})
					case ">":
						// true branch
						queue = append(queue, QueuePart{currQueue.name, currQueue.conditionIndex + 1, Ranges{currQueue.ranges.xRange, currQueue.ranges.mRange, currQueue.ranges.aRange, Range{currQueue.ranges.sRange.low, condition.value}}})
						// false branch
						queue = append(queue, QueuePart{condition.next, 0, Ranges{currQueue.ranges.xRange, currQueue.ranges.mRange, currQueue.ranges.aRange, Range{condition.value + 1, currQueue.ranges.sRange.high}}})
					}
					continue queue
				}
			}

		}
	}

	return combinations
}

func part1(workflows map[string]Workflow, parts []Part) int {
	return sliceutils.Reduce(parts, func(acc int, value Part, index int, slice []Part) int {
		return acc + processPart(workflows, value)
	}, 0)
}

func part2(workflows map[string]Workflow) int {
	return processAllParts(workflows)
}
