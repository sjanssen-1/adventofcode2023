package main

import (
	"adventofcode2023/util"
	"bufio"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/Goldziher/go-utils/sliceutils"
)

// https://adventofcode.com/2023/day/2
func main() {
	gamesRecord := util.ReadFile("./day2/demo_input.txt")
	defer util.TimeTrack(time.Now(), "main")

	games := parseGames(gamesRecord)

	log.Default().Printf("P1: %d", part1(games))
	log.Default().Printf("P2: %d", part2(games))
}

type Game struct {
	number        int
	maxRedCubes   int
	maxGreenCubes int
	maxBlueCubes  int
}

func newGame(number int) *Game {
	game := Game{number: number}
	return &game
}

func parseGames(gamesRecord *bufio.Scanner) []Game {
	games := make([]Game, 0)

	for gamesRecord.Scan() {
		if gamesRecord.Text() != "" {
			gameLine := gamesRecord.Text()
			// log.Default().Printf("Game line: %s", gameLine)

			gameLineSplit := strings.Split(gameLine, ": ")
			gameNumber, _ := strconv.Atoi(strings.Split(gameLineSplit[0], " ")[1])
			game := newGame(gameNumber)
			gameRounds := strings.Split(gameLineSplit[1], "; ")
			for _, gameRound := range gameRounds {
				cubes := strings.Split(gameRound, ", ")
				for _, cube := range cubes {
					throw := strings.Split(cube, " ")
					throwValue, _ := strconv.Atoi(throw[0])
					throwColor := throw[1]
					switch throwColor {
					case "red":
						game.maxRedCubes = max(game.maxRedCubes, throwValue)
					case "green":
						game.maxGreenCubes = max(game.maxGreenCubes, throwValue)
					case "blue":
						game.maxBlueCubes = max(game.maxBlueCubes, throwValue)
					}
				}
			}
			games = append(games, *game)
		}
	}
	return games
}

func part1(games []Game) int {
	redCubes := 12
	greenCubes := 13
	blueCubes := 14

	validGames := sliceutils.Filter(games, func(value Game, index int, slice []Game) bool {
		return value.maxRedCubes <= redCubes && value.maxGreenCubes <= greenCubes && value.maxBlueCubes <= blueCubes
	})

	return sliceutils.Reduce(
		validGames,
		func(acc int, value Game, index int, slice []Game) int {
			return acc + value.number
		},
		0,
	)
}

func part2(games []Game) int {
	// log.Printf("%#v", games)
	return sliceutils.Reduce(
		games,
		func(acc int, value Game, index int, slice []Game) int {
			return acc + (value.maxRedCubes * value.maxGreenCubes * value.maxBlueCubes)
		},
		0,
	)
}
