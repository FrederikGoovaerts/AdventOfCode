package util

import (
	"os"
	"strings"
)

func FileAsString(fileName string) string {
	dat, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	return string(dat)
}

func FileAsLines(fileName string) []string {
	return strings.Split(FileAsString(fileName), "\n")
}

func FileAsNumbers(fileName string) []int {
	lines := FileAsLines(fileName)
	result := make([]int, 0, len(lines))

	for _, line := range lines {
		result = append(result, StringToInt(line))
	}
	return result
}
