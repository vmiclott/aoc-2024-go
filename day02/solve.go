package day02

import (
	"fmt"
)

var debug bool

func reportIsSafe(report []int) bool {
	if len(report) < 2 {
		return true
	}
	previousLevel := report[0]
	level := report[1]
	direction := sign(level - previousLevel)
	if !transitionIsSafe(previousLevel, level, direction) {
		return false
	}
	for i := 2; i < len(report); i++ {
		previousLevel = level
		level = report[i]
		if !transitionIsSafe(previousLevel, level, direction) {
			return false
		}
	}
	return true
}

func transitionIsSafe(previousLevel int, level int, direction int) bool {
	diff := level - previousLevel
	return (direction > 0 && diff >= 1 && diff <= 3) || (direction < 0 && diff <= -1 && diff >= -3)
}

func sign(number int) int {
	if number < 0 {
		return -1
	}
	return +1
}

func solve(reports [][]int) int {
	var result int
	for _, report := range reports {
		if reportIsSafe(report) {
			result++
		} else {
		}
	}
	return result
}

func solvePartTwo(reports [][]int) int {
	var result int
	for _, report := range reports {
		if reportIsSafe(report) {
			result++
		} else {
			for i := range report {
				reportWithoutI := append([]int{}, report[:i]...)
				reportWithoutI = append(reportWithoutI, report[i+1:]...)
				if reportIsSafe(reportWithoutI) {
					result++
					break
				}
			}
		}
	}
	return result
}

func Solve(inputFileName string, d bool) error {
	debug = d
	reports, err := parse(inputFileName)
	if err != nil {
		return err
	}
	fmt.Println("Part 1:", solve(reports))
	fmt.Println("Part 2:", solvePartTwo(reports))
	return nil
}
