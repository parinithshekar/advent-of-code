package main

import (
	"fmt"
	"strconv"
	"strings"
)

func two() {
	// lines := readFileLines("inputs/02")

	// reports := reports(lines)

	answer1 := 0
	answer2 := 0
	// for _, report := range reports {
	// 	if (isSafe(report)) {
	// 		answer1++
	// 	}
	// 	if (isSafeWithTolerance(report)) {
	// 		answer2++
	// 	}
	// }
	test := []int{1, 2, 3, 4, 5, 6}
	isSafeWithTolerance(test)

	fmt.Println("2.1 -> ", answer1)
	fmt.Println("2.2 -> ", answer2)
}

func isSafe(report []int) bool {
	oldDiff := 0
	for i := 1 ; i < len(report) ; i++ {
		diff := report[i] - report[i-1]
		if (diff == 0) {return false}
		if (abs(diff)>3) {return false}
		if (oldDiff * diff < 0) {return false}
		oldDiff = diff
	}
	return true
}

func isSafeWithTolerance(report []int) bool {
	for i := 0 ; i < len(report) ; i++ {
		reducedReport := append(report[:i], report[i+1:]...)
		fmt.Println(reducedReport)
		// if (isSafe(reducedReport)) { return true}
	}
	return false
}

func reports(lines []string) [][]int {
	result := [][]int {}
	for _, line := range lines {
		if (line == "") {
			continue
		}
		report := strings.Fields(line)
		result = append(result, levels(report))
	}
	return result
}

func levels(report []string) []int {
	result := []int{}
	for _, levelString := range report {
		level, _ := strconv.Atoi(levelString)
		result = append(result, level)
	}
	return result
}