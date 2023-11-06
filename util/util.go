package util

import (
	"bufio"
	"log"
	"os"
)

func ReadFile(path string) *bufio.Scanner {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	return scanner
}
