package main

import (
	"adventofcode2023/util"
	"log"
	"time"

	"github.com/Goldziher/go-utils/sliceutils"
)

var raceDetails1 = []Race{
	{
		time:     7,
		distance: 9,
	},
	{
		time:     15,
		distance: 40,
	},
	{
		time:     30,
		distance: 200,
	},
}

var raceDetails2 = []Race{
	{
		time:     42,
		distance: 308,
	},
	{
		time:     89,
		distance: 1170,
	},
	{
		time:     91,
		distance: 1291,
	},
	{
		time:     89,
		distance: 1467,
	},
}

var raceDetails3 = []Race{
	{
		time:     71530,
		distance: 940200,
	},
}

var raceDetails4 = []Race{
	{
		time:     42899189,
		distance: 308117012911467,
	},
}

type Race struct {
	time     int
	distance int
}

// https://adventofcode.com/2023/day/6
func main() {
	defer util.TimeTrack(time.Now(), "main")

	part1()
	part2()
}

func simulateRace(raceDetails []Race) int {
	return sliceutils.Reduce(
		raceDetails,
		func(acc int, cur Race, index int, slice []Race) int {
			var possibleWins int
			for pushButtonTime := 1; pushButtonTime < cur.time-1; pushButtonTime++ {
				if (pushButtonTime * (cur.time - pushButtonTime)) > cur.distance {
					possibleWins++
				}
			}
			return acc * possibleWins
		},
		1,
	)
}

func part1() {
	log.Default().Printf("P1 (sample): %d", simulateRace(raceDetails1))
	log.Default().Printf("P1: %d", simulateRace(raceDetails2))
}

func part2() {
	log.Default().Printf("P2 (sample): %d", simulateRace(raceDetails3))
	log.Default().Printf("P2: %d", simulateRace(raceDetails4))
}
