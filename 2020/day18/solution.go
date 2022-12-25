package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func solveRightward(line string, startIndex int) (int, int) {
	index := startIndex
	runes := []rune(line)

	solution := -1

	operator := '+'
	acc := ""

	performFold := func() {
		if solution == -1 {
			solution, _ = strconv.Atoi(acc)
		} else {
			val, _ := strconv.Atoi(acc)
			if operator == '+' {
				solution += val
			} else {
				solution *= val
			}
		}
	}

	for index < len(line) {
		rune := runes[index]
		if rune == '(' {
			result, resultIndex := solveRightward(line, index+1)
			acc = fmt.Sprint(result)
			index = resultIndex
		} else if rune == ')' {
			performFold()
			return solution, index
		} else if rune == '*' || rune == '+' {
			performFold()
			operator = rune
			acc = ""
		} else {
			acc += string(rune)
		}
		index++
	}
	performFold()

	return solution, index

}

func solveAdvanced(line string, startIndex int) (int, int) {
	index := startIndex
	runes := []rune(line)

	solution := -1
	plusAcc := 0

	numAcc := ""

	foldPlus := func() {
		val, _ := strconv.Atoi(numAcc)
		plusAcc += val
		numAcc = ""
	}

	foldMult := func() {
		if solution == -1 {
			solution = plusAcc
		} else {
			solution *= plusAcc
		}
		plusAcc = 0
	}

	for index < len(line) {
		rune := runes[index]
		if rune == '(' {
			result, resultIndex := solveAdvanced(line, index+1)
			plusAcc += result
			index = resultIndex
		} else if rune == ')' {
			foldPlus()
			foldMult()
			return solution, index
		} else if rune == '*' {
			foldPlus()
			foldMult()
		} else if rune == '+' {
			foldPlus()
		} else {
			numAcc += string(rune)
		}
		index++
	}
	foldPlus()
	foldMult()

	return solution, index

}

func main() {
	dat, err := ioutil.ReadFile("input")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(strings.TrimSpace(string(dat)), "\n")
	sum := 0
	for _, line := range lines {
		line = strings.Replace(line, " ", "", -1)
		res, _ := solveRightward(line, 0)
		sum += res
	}
	fmt.Println(sum)

	sum = 0
	for _, line := range lines {
		line = strings.Replace(line, " ", "", -1)
		res, _ := solveAdvanced(line, 0)
		sum += res
	}
	fmt.Println(sum)

}
