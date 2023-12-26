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
	theSystem := util.ReadFile("C:\\Users\\janss\\git\\adventofcode2023\\day19\\input.txt")
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
	log.Default().Printf("P2: %d", part2())
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
	operand  string
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
			switch condition.operand {
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

func part1(workflows map[string]Workflow, parts []Part) int {
	return sliceutils.Reduce(parts, func(acc int, value Part, index int, slice []Part) int {
		return acc + processPart(workflows, value)
	}, 0)
}

func part2() int {
	return 0
}
