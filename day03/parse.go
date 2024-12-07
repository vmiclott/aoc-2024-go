package day03

import (
	"os"
)

func parse(inputFileName string) (string, error) {
	content, err := os.ReadFile(inputFileName)
	if err != nil {
		return "", err
	}
	return string(content), nil
}
