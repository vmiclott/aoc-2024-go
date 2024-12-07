package day01

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

	var locationIdSlices = [][]int{}
	for scanner.Scan() {
		line := scanner.Text()
		numbers, err := parseLine(line)
		if err != nil {
			return nil, err
		}
		for len(locationIdSlices) < len(numbers) {
			locationIdSlices = append(locationIdSlices, []int{})
		}
		for index, number := range numbers {
			locationIdSlices[index] = append(locationIdSlices[index], number)
		}
	}

	return locationIdSlices, nil
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
