package day04

import (
	"bufio"
	"os"
)

func parse(inputFileName string) ([][]rune, error) {
	file, err := os.Open(inputFileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	var result [][]rune
	var line []rune
	for {
		r, _, err := reader.ReadRune()
		if err != nil {
			break
		}
		if r == '\n' {
			result = append(result, line)
			line = make([]rune, 0)
		} else {
			line = append(line, r)
		}
	}
	if len(line) > 0 {
		result = append(result, line)
	}
	return result, nil
}
