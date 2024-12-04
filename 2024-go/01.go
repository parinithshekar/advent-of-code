package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func one() {
	left := []int{}
	right := []int{}
	count := make(map[int]int)
	
	lines := readFileLines("inputs/01")
	for _, line := range lines {
		if (line == "") {
			continue
		}
		locationIds := strings.Fields(line)
		leftId, _ := strconv.Atoi(locationIds[0])
		rightId, _ := strconv.Atoi(locationIds[1])
		
		left = append(left, leftId)
		right = append(right, rightId)
		count[rightId] += 1
	}

	slices.Sort(left)
	slices.Sort(right)

	answer1 := 0
	answer2 := 0
	for i, id := range left {
		answer1 += abs(left[i] - right[i])
		answer2 += id * count[id]
	}

	fmt.Println("Day 1 Answer 1 -> ", answer1)
	fmt.Println("Day 1 Answer 2 -> ", answer2)
}
