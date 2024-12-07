package day02

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func parse(inputFileName string) ([][]int, error) {
	readFile, err := os.Open(inputFileName)
	if err != nil {
		return nil, err
	}
	defer readFile.Close()

	scanner := bufio.NewScanner(readFile)
	scanner.Split(bufio.ScanLines)

	var reports = [][]int{}
	for scanner.Scan() {
		line := scanner.Text()
		numbers, err := parseLine(line)
		if err != nil {
			return nil, err
		}
		reports = append(reports, numbers)
	}

	return reports, nil
}

func parseLine(line string) ([]int, error) {
	numbers := []int{}
	fields := strings.Fields(line)
	for _, field := range fields {
		number, err := strconv.Atoi(field)
		if err != nil {
			return numbers, err
		}
		numbers = append(numbers, number)
	}
	return numbers, nil
}
