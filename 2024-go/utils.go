package main

import (
	"os"
	"log"
	"strings"
)

func check(err error) {
	if err != nil {
		log.Fatal("Error: %s", err)
		panic(err)
	}
}

func abs(i int) int {
	if (i < 0) {
		return -1 * i
	}
	return i
}

func readFileLines(fileLocation string) []string {
	bytes, err := os.ReadFile(fileLocation)
	check(err)
	return strings.Split(string(bytes), "\n") 
}
