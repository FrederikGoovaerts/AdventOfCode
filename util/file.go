package util

import (
	"os"
	"strings"
)

func FileAsLines(fileName string) []string {
	dat, err := os.ReadFile("input")
	if err != nil {
		panic(err)
	}
	return strings.Split(string(dat), "\n")
}
