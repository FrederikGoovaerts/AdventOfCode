package util

import (
	"fmt"
	"strconv"
)

func Serialize(items ...any) string {
	result := fmt.Sprint(items[0])
	for i := 1; i < len(items); i++ {
		result += " " + fmt.Sprint(items[i])
	}
	return result
}

func StringToInt(in string) int {
	res, _ := strconv.Atoi(in)
	return res
}

func Contains[T int | string](list []T, search T) bool {
	for _, el := range list {
		if el == search {
			return true
		}
	}
	return false
}

func Remove[T int | string](list []T, search ...T) []T {
	result := make([]T, 0)
	for _, el := range list {
		if !Contains(search, el) {
			result = append(result, el)
		}
	}
	return result
}
