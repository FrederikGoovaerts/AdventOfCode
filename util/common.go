package util

import "fmt"

func Serialize[T int | string](x T, y T) string {
	return fmt.Sprint(x) + " " + fmt.Sprint(y)
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
