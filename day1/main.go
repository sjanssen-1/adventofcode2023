package main

import (
	"adventofcode2023/util"
	"fmt"
	"log"
	"sort"
	"strconv"
)

type elf struct {
	calories int
}

func main() {
	fileScanner := util.ReadFile("./day1/input2.txt")

	var elves []elf

	var currentElf = elf{}

	for fileScanner.Scan() {
		if fileScanner.Text() != "" {
			foundCalories, err := strconv.Atoi(fileScanner.Text())
			if err != nil {
				log.Fatal(err)
			}
			currentElf.calories += foundCalories
		} else {
			elves = append(elves, currentElf)
			currentElf = elf{}
		}
	}

	sort.Slice(elves, func(i, j int) bool {
		return elves[i].calories > elves[j].calories
	})

	fmt.Println("P1: ", elves[0].calories)
	fmt.Println("P2: ", elves[0].calories+elves[1].calories+elves[2].calories)
}
