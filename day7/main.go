package main

import (
	"adventofcode2023/util"
	"log"
	"slices"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/Goldziher/go-utils/sliceutils"
)

// https://adventofcode.com/2023/day/7
func main() {
	handsInput := util.ReadFile("day7/input.txt")
	defer util.TimeTrack(time.Now(), "main")

	hands := make([]Hand, 0)

	for handsInput.Scan() {
		hand := handsInput.Text()
		hands = append(hands, toHand(hand))
	}

	log.Default().Printf("P1: %d", part1(hands))
	log.Default().Printf("P2: %d", part2(hands))
}

type HandType int

const (
	HIGH  HandType = iota
	ONE   HandType = iota
	TWO   HandType = iota
	THREE HandType = iota
	FULL  HandType = iota
	FOUR  HandType = iota
	FIVE  HandType = iota
)

// for joker
func nextBestHandType(handType HandType) HandType {
	if handType == ONE {
		return THREE
	} else if handType == THREE {
		return FOUR
	} else if handType == FIVE {
		return FIVE
	} else if handType == TWO {
		return FULL
	} else {
		return handType + 1
	}
}

type Hand struct {
	hand       string
	handType   HandType
	handTypeP2 HandType
	bet        int
}

func getPossibleCardValues() []string {
	return []string{"A", "K", "Q", "J", "T", "9", "8", "7", "6", "5", "4", "3", "2"}
}

func getPossibleCardValuesP2() []string {
	return []string{"A", "K", "Q", "T", "9", "8", "7", "6", "5", "4", "3", "2", "J"}
}

func toHand(hand string) Hand {
	handSplit := strings.Split(hand, " ")
	bet, _ := strconv.Atoi(handSplit[1])

	return Hand{
		hand:       handSplit[0],
		handType:   detectType(handSplit[0]),
		handTypeP2: detectTypeP2(handSplit[0]),
		bet:        bet,
	}
}

func detectType(hand string) HandType {
	duplicates := make([]int, 0)
	for _, cardValue := range getPossibleCardValues() {
		duplicates = append(duplicates, strings.Count(hand, cardValue))
	}
	max := slices.Max(duplicates)
	if max == 5 {
		return FIVE
	} else if max == 4 {
		return FOUR
	} else if max == 3 && slices.Contains(duplicates, 2) {
		return FULL
	} else if max == 3 {
		return THREE
	} else if max == 2 && len(sliceutils.FindIndexes(duplicates, func(value int, index int, slice []int) bool {
		return value == 2
	})) == 2 {
		return TWO
	} else if max == 2 {
		return ONE
	} else {
		return HIGH
	}
}

func detectTypeP2(hand string) HandType {
	duplicates := make([]int, 0)
	for _, cardValue := range getPossibleCardValuesP2()[0:12] {
		duplicates = append(duplicates, strings.Count(hand, cardValue))
	}
	jokers := strings.Count(hand, "J")
	max := slices.Max(duplicates)
	var handType HandType
	if max == 5 {
		handType = FIVE
	} else if max == 4 {
		handType = FOUR
	} else if max == 3 && slices.Contains(duplicates, 2) {
		handType = FULL
	} else if max == 3 {
		handType = THREE
	} else if max == 2 && len(sliceutils.FindIndexes(duplicates, func(value int, index int, slice []int) bool {
		return value == 2
	})) == 2 {
		handType = TWO
	} else if max == 2 {
		handType = ONE
	} else {
		handType = HIGH
	}

	for i := 0; i < jokers; i++ {
		handType = nextBestHandType(handType)
	}
	return handType
}

func part1(hands []Hand) int {
	sort.Slice(hands, func(i, j int) bool {
		leftHand := hands[i]
		rightHand := hands[j]

		if leftHand.handType == rightHand.handType {
			possibleCardValues := getPossibleCardValues()
			for c := 0; c < 5; c++ {
				if hands[i].hand[c] == hands[j].hand[c] {
					continue
				}
				return slices.Index(possibleCardValues, string(leftHand.hand[c])) >
					slices.Index(possibleCardValues, string(rightHand.hand[c]))
			}
			return false
		}
		return leftHand.handType < rightHand.handType
	})

	return sliceutils.Reduce(
		hands,
		func(acc int, cur Hand, index int, slice []Hand) int {
			return acc + (cur.bet * (index + 1))
		},
		0,
	)
}

func part2(hands []Hand) int {
	handsCopy := make([]Hand, len(hands))
	copy(handsCopy, hands)

	// log.Println(handsCopy)
	sort.Slice(handsCopy, func(i, j int) bool {
		leftHand := handsCopy[i]
		rightHand := handsCopy[j]

		if leftHand.handTypeP2 == rightHand.handTypeP2 {
			possibleCardValues := getPossibleCardValuesP2()
			for c := 0; c < 5; c++ {
				if handsCopy[i].hand[c] == handsCopy[j].hand[c] {
					continue
				}
				return slices.Index(possibleCardValues, string(leftHand.hand[c])) >
					slices.Index(possibleCardValues, string(rightHand.hand[c]))
			}
			return false
		}
		return leftHand.handTypeP2 < rightHand.handTypeP2
	})

	return sliceutils.Reduce(
		handsCopy,
		func(acc int, cur Hand, index int, slice []Hand) int {
			return acc + (cur.bet * (index + 1))
		},
		0,
	)
}
