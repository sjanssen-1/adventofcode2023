package main

import (
	"adventofcode2023/util"
	"log"
	"slices"
	"strings"
	"time"

	"github.com/Goldziher/go-utils/maputils"
)

// https://adventofcode.com/2023/day/20
func main() {
	moduleConfiguration := util.ScannerToStringSlice(*util.ReadFile("/Users/stefjanssens/git/adventofcode2023/day20/input.txt"))
	defer util.TimeTrack(time.Now(), "main")

	log.Default().Printf("P1: %d", part1(parseModules(moduleConfiguration)))
	log.Default().Printf("P2: %d", part2(parseModules(moduleConfiguration)))
}

func parseModules(moduleConfiguration []string) map[string]Module {
	modules := make(map[string]Module)

	for _, moduleConfig := range moduleConfiguration {
		split := strings.Split(moduleConfig, " -> ")
		if split[0] == "broadcaster" {
			modules["broadcaster"] = NewModule(split[0], split[1])
		} else {
			modules[split[0][1:]] = NewModule(split[0], split[1])
		}
	}

	for conjunctionName, conjunctionModule := range modules {
		if conjunctionModule.isConjunction {
			for sourceName, sourceModule := range modules {
				if slices.Contains(sourceModule.destinations, conjunctionName) {
					conjunctionModule.sourcesMemory[sourceName] = false
				}
			}
			modules[conjunctionName] = conjunctionModule
		}
	}
	return modules
}

type Module struct {
	isFlipFlop    bool
	isConjunction bool
	isOn          bool
	destinations  []string
	sourcesMemory map[string]bool
}

func NewModule(name string, destinations string) Module {

	module := Module{}
	module.destinations = strings.Split(destinations, ", ")
	module.sourcesMemory = make(map[string]bool)

	if name[0] == '%' {
		// flipflop
		module.isFlipFlop = true
	} else {
		// conjunction
		module.isConjunction = true
	}
	return module
}

type Pulse struct {
	isHigh      bool
	source      string
	destination string
}

func pressButton(modules map[string]Module) (int, int, map[string]bool) {
	lows := 0
	highs := 0

	// hardcoded to my input idgaf
	bqSources := make(map[string]bool)
	for key, _ := range modules["bq"].sourcesMemory {
		bqSources[key] = false
	}

	buffer := []Pulse{{false, "me", "broadcaster"}}

	for len(buffer) > 0 {
		pulse := buffer[0]
		buffer = buffer[1:]

		if pulse.isHigh {
			highs++
		} else {
			lows++
		}

		module, exists := modules[pulse.destination]
		if !exists {
			continue
		}

		if pulse.destination == "broadcaster" {
			for _, destination := range module.destinations {
				buffer = append(buffer, Pulse{false, pulse.destination, destination})
			}
		} else if module.isFlipFlop {
			if pulse.isHigh {
				continue
			}
			if module.isOn {
				module.isOn = false
				for _, destination := range module.destinations {
					buffer = append(buffer, Pulse{false, pulse.destination, destination})
				}
			} else {
				module.isOn = true
				for _, destination := range module.destinations {
					buffer = append(buffer, Pulse{true, pulse.destination, destination})
				}
			}
		} else {
			if pulse.isHigh {
				module.sourcesMemory[pulse.source] = true
			} else {
				module.sourcesMemory[pulse.source] = false
			}

			if !slices.Contains(maputils.Values(module.sourcesMemory), false) {
				// all highs ; send low
				for _, destination := range module.destinations {
					buffer = append(buffer, Pulse{false, pulse.destination, destination})
				}
			} else {
				// send high
				for _, destination := range module.destinations {
					// hardcoded to my input idgaf.
					if destination == "bq" {
						bqSources[pulse.destination] = true
					}
					buffer = append(buffer, Pulse{true, pulse.destination, destination})
				}
			}
		}
		modules[pulse.destination] = module
	}

	return lows, highs, bqSources
}

func part1(modules map[string]Module) int {
	totalLows, totalHighs := 0, 0

	for i := 0; i < 1000; i++ {
		lows, highs, _ := pressButton(modules)
		totalLows += lows
		totalHighs += highs
	}
	return totalLows * totalHighs
}

func part2(modules map[string]Module) int {
	counters := make([]int, 0)
	counter := 0
	for {
		counter++
		// hardcoded to my input idgaf.
		_, _, bqSources := pressButton(modules)
		if slices.Contains(maputils.Values(bqSources), true) {
			counters = append(counters, counter)
		}

		if len(counters) == len(bqSources) {
			break
		}

	}
	return LCM(counters[0], counters[1], counters[2:])
}

// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers []int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[0], integers[1:])
	}

	return result
}
