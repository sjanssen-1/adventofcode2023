package util

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/Goldziher/go-utils/sliceutils"
)

func ReadFile(path string) *bufio.Scanner {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	return scanner
}

func TimeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}

func SpaceDelimitedStringToIntSlice(s string) []int {
	return sliceutils.Map(strings.Split(s, " "), func(value string, index int, slice []string) int {
		number, _ := strconv.Atoi(value)
		return number
	})
}

func ScannerToStringSlice(scanner bufio.Scanner) []string {
	stringSlice := make([]string, 0)
	for scanner.Scan() {
		stringSlice = append(stringSlice, scanner.Text())
	}
	return stringSlice
}
