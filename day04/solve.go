package day04

import (
	"fmt"
)

var debug bool

func Solve(inputFileName string, d bool) error {
	debug = d
	content, err := parse(inputFileName)
	if err != nil {
		return err
	}
	fmt.Println("Part 1:", partOne(content))
	fmt.Println("Part 2:", partTwo(content))
	return nil
}
