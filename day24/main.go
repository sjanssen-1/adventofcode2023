package main

import (
	"adventofcode2023/util"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/Goldziher/go-utils/sliceutils"
)

// https://adventofcode.com/2023/day/24
func main() {
	noteDemo := util.ScannerToStringSlice(*util.ReadFile("C:\\Users\\janss\\git\\adventofcode2023\\day24\\demo_input.txt"))
	note := util.ScannerToStringSlice(*util.ReadFile("C:\\Users\\janss\\git\\adventofcode2023\\day24\\input.txt"))
	defer util.TimeTrack(time.Now(), "main")

	hailStonesDemo := parseHailStones(noteDemo)
	hailStones := parseHailStones(note)

	log.Default().Printf("P1 demo: %d", part1(hailStonesDemo, 7, 27))
	log.Default().Printf("P1: %d", part1(hailStones, 200000000000000, 400000000000000))
	log.Default().Printf("P2 demo: %d", part2())
}

// ax + by = c
type HailStone struct {
	sx, sy, sz float64
	vx, vy, vz float64
	a, b, c    float64
}

func NewHailStone(s string) HailStone {
	split := strings.Split(s, " @ ")
	positions := sliceutils.Map(strings.Split(split[0], ", "), func(value string, index int, slice []string) int {
		position, _ := strconv.Atoi(value)
		return position
	})

	velocities := sliceutils.Map(strings.Split(split[1], ", "), func(value string, index int, slice []string) int {
		velocity, _ := strconv.Atoi(value)
		return velocity
	})

	sx := positions[0]
	sy := positions[1]
	sz := positions[2]

	vx := velocities[0]
	vy := velocities[1]
	vz := velocities[2]

	a := vy
	b := -vx
	c := vy*sx - vx*sy

	return HailStone{float64(sx), float64(sy), float64(sz), float64(vx), float64(vy), float64(vz), float64(a), float64(b), float64(c)}
}

func parseHailStones(notes []string) []HailStone {
	hailstones := make([]HailStone, 0)
	for _, note := range notes {
		hailstones = append(hailstones, NewHailStone(note))
	}
	return hailstones
}

func part1(hailStones []HailStone, min, max float64) int {
	total := 0

	for i, hs1 := range hailStones {
		for _, hs2 := range hailStones[:i] {
			a1, b1, c1 := hs1.a, hs1.b, hs1.c
			a2, b2, c2 := hs2.a, hs2.b, hs2.c
			if a1*b2 == b1*a2 {
				continue
			}
			x := (c1*b2 - c2*b1) / (a1*b2 - a2*b1)
			y := (c2*a1 - c1*a2) / (a1*b2 - a2*b1)
			if x >= min && x <= max && y >= min && y <= max {
				if ((x-hs1.sx)*hs1.vx >= 0 && (y-hs1.sy)*hs1.vy >= 0) &&
					((x-hs2.sx)*hs2.vx >= 0 && (y-hs2.sy)*hs2.vy >= 0) {
					total += 1
				}
			}
		}
	}
	return total
}

func part2() int {
	return 0
}
