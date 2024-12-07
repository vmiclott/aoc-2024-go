package day01

import (
	"fmt"
	"math"
	"sort"
)

var debug bool

func solve(locationIdSlices [][]int) int {
	var result int
	for _, locationIdSlice := range locationIdSlices {
		sort.Ints(locationIdSlice)
	}
	for i := range locationIdSlices[0] {
		result += int(math.Abs(float64(locationIdSlices[0][i] - locationIdSlices[1][i])))
	}
	return result
}

func solvePartTwo(locationIdSlices [][]int) int {
	var result int
	counter := make(map[int]int)
	for _, locationId := range locationIdSlices[1] {
		counter[locationId]++
	}
	for _, locationId := range locationIdSlices[0] {
		result += locationId * counter[locationId]
	}
	return result
}

func Solve(inputFileName string, d bool) error {
	debug = d
	locationIdSlices, err := parse(inputFileName)
	if err != nil {
		return err
	}
	fmt.Println("Part 1:", solve(locationIdSlices))
	fmt.Println("Part 2:", solvePartTwo(locationIdSlices))
	return nil
}
