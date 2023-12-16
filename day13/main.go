package main

import (
	"adventofcode2023/util"
	"fmt"
	"log"
	"time"
)

// https://adventofcode.com/2023/day/13
func main() {
	valleyOfMirrors := util.ScannerToStringSlice(*util.ReadFile("C:\\Users\\janss\\git\\adventofcode2023\\day13\\input.txt"))
	valleyOfMirrorsParsed := parseValleyOfMirrors(valleyOfMirrors)

	defer util.TimeTrack(time.Now(), "main")

	log.Default().Printf("P1: %d", part1(valleyOfMirrorsParsed))
	log.Default().Printf("P2: %d", part2(valleyOfMirrorsParsed))
}

func parseValleyOfMirrors(valleyOfMirrors []string) [][]string {
	valleyOfMirrorsParsed := make([][]string, 0)

	currentValley := make([]string, 0)
	for _, valley := range valleyOfMirrors {
		if valley == "" {
			valleyOfMirrorsParsed = append(valleyOfMirrorsParsed, currentValley)
			currentValley = make([]string, 0)
		} else {
			currentValley = append(currentValley, valley)
		}
	}
	valleyOfMirrorsParsed = append(valleyOfMirrorsParsed, currentValley)
	return valleyOfMirrorsParsed
}

func part1(valleyOfMirrors [][]string) int {
	summary := 0
	for i := range valleyOfMirrors {
		horizontalResult := findHorizontalReflection(valleyOfMirrors[i], false)
		if horizontalResult != -1 {
			summary += horizontalResult * 100
		} else {
			summary += findVerticalReflection(valleyOfMirrors[i], false)
		}
	}
	return summary
}

func isSmudged(s1 string, s2 string) bool {
	if s1 == "" || s2 == "" {
		return false
	}
	differences := 0
	for i := range s1 {
		if s1[i] != s2[i] {
			differences++
		}
		if differences > 1 {
			return false
		}
	}
	return true
}

type Match struct {
	index      int
	withSmudge bool
}

func findHorizontalReflection(valley []string, fixSmudge bool) int {
	// horizontal
	previousHorizontal := ""
	matches := make([]Match, 0)
	for y := range valley {
		if previousHorizontal == valley[y] {
			matches = append(matches, Match{y, false})
		} else if fixSmudge && isSmudged(previousHorizontal, valley[y]) {
			matches = append(matches, Match{y, true})
		}
		previousHorizontal = valley[y]
	}

	if len(matches) != 0 {
	OUTER:
		for _, match := range matches {
			// matchIndex and matchIndex -1 are the OGs
			smudgeFixed := false
			if match.index > len(valley)/2 {
				// move down
				j := match.index - 2
				for i := match.index + 1; i < len(valley); i++ {
					if valley[i] != valley[j] {
						if fixSmudge {
							if match.withSmudge {
								continue OUTER
							} else if !smudgeFixed && isSmudged(valley[i], valley[j]) {
								smudgeFixed = true
								j--
								continue
							}
						}
						continue OUTER
					} else {
						j--
					}
				}
			} else {
				// move up
				j := match.index + 1
				for i := match.index - 2; i > -1; i-- {
					if valley[i] != valley[j] {
						if fixSmudge {
							if match.withSmudge {
								continue OUTER
							} else if !smudgeFixed && isSmudged(valley[i], valley[j]) {
								smudgeFixed = true
								j++
								continue
							}
						}
						continue OUTER
					} else {
						j++
					}
				}
			}
			if fixSmudge && (smudgeFixed || match.withSmudge) {
				return match.index
			} else if !fixSmudge {
				return match.index
			}
		}
	}
	return -1
}

func findVerticalReflection(valley []string, fixSmudge bool) int {
	return findHorizontalReflection(rotateValley(valley), fixSmudge)
}

func rotateValley(valley []string) []string {
	rotatedValley := make([]string, 0)

	for x := 0; x < len(valley[0]); x++ {
		s := ""
		for y := len(valley) - 1; y >= 0; y-- {
			s += string(valley[y][x])
		}
		rotatedValley = append(rotatedValley, s)
	}
	// debug(rotatedValley)
	return rotatedValley
}

// func debug(valley []string) {
// 	for _, v := range valley {
// 		fmt.Println(v)
// 	}
// }

func part2(valleyOfMirrors [][]string) int {
	summary := 0
	for i := range valleyOfMirrors {
		horizontalResult := 100 * findHorizontalReflection(valleyOfMirrors[i], true)
		if horizontalResult > 0 {
			fmt.Println("horizontal ", horizontalResult)
			summary += horizontalResult
		} else {
			verticalResult := findVerticalReflection(valleyOfMirrors[i], true)
			fmt.Println("vertical ", verticalResult)
			summary += verticalResult
		}
	}
	return summary
}
