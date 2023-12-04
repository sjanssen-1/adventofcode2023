package main

import (
	"adventofcode2023/util"
	"bufio"
	"log"
	"math"
	"strconv"
	"strings"
	"time"

	"github.com/Goldziher/go-utils/sliceutils"
)

// https://adventofcode.com/2023/day/4
func main() {
	pileOfCards := util.ReadFile("./day4/input.txt")
	defer util.TimeTrack(time.Now(), "main")

	cards := parseCards(*pileOfCards)

	log.Default().Printf("P1: %d", part1(cards))
	log.Default().Printf("P2: %d", part2(cards))
}

type Card struct {
	winningNumbers []int
	yourNumbers    []int
}

func calculateScore(card Card) int {
	log.Println(sliceutils.Intersection(card.winningNumbers, card.yourNumbers))
	return int(math.Pow(2, float64(len(sliceutils.Intersection(card.winningNumbers, card.yourNumbers))-1)))
}

func parseCard(stringCard string) Card {
	cardParts := strings.Split(stringCard, " | ")
	winningNumbersString := strings.ReplaceAll(strings.TrimSpace(strings.Split(cardParts[0], ": ")[1]), "  ", " ")
	// log.Println(winningNumbersString)
	yourNumbersString := strings.ReplaceAll(strings.TrimSpace(cardParts[1]), "  ", " ")
	// log.Println(yourNumbersString)

	return Card{
		winningNumbers: sliceutils.Map(strings.Split(winningNumbersString, " "), func(value string, index int, slice []string) int {
			winningNumber, _ := strconv.Atoi(strings.TrimSpace(value))
			return winningNumber
		}),
		yourNumbers: sliceutils.Map(strings.Split(yourNumbersString, " "), func(value string, index int, slice []string) int {
			yourNumber, _ := strconv.Atoi(strings.TrimSpace(value))
			return yourNumber
		}),
	}
}

func parseCards(pileOfCards bufio.Scanner) []Card {
	cards := make([]Card, 0)
	for pileOfCards.Scan() {
		card := parseCard(pileOfCards.Text())
		cards = append(cards, card)
	}
	return cards
}

func part1(cards []Card) int {
	return sliceutils.Reduce(
		cards,
		func(acc int, cur Card, index int, slice []Card) int {
			return acc + calculateScore(cur)
		},
		0,
	)
}

func part2(cards []Card) int {
	cardAmounts := make([]int, len(cards))

	for i, card := range cards {
		// add original copy
		cardAmounts[i]++

		score := len(sliceutils.Intersection(card.winningNumbers, card.yourNumbers))
		log.Printf("Card %d has %d matches", i, score)
		for j := 1; j <= score; j++ {
			log.Printf("Adding %d to card number %d", score, i+j)
			cardAmounts[i+j] += 1 * cardAmounts[i]
		}
		log.Printf("Card amounts: %v", cardAmounts)
	}

	return sliceutils.Reduce(
		cardAmounts,
		func(acc int, cur int, index int, slice []int) int {
			return acc + cur
		},
		0,
	)
}
