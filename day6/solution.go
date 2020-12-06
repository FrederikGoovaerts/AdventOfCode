package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	dat, err := ioutil.ReadFile("input")
	if err != nil {
		panic(err)
	}

	groups := strings.Split(strings.TrimSpace(string(dat)), "\n\n")
	answerList := make([](map[string]int), 0)
	peopleCount := make([]int, 0)
	for _, group := range groups {
		answers := make(map[string]int, 0)
		appended := strings.ReplaceAll(group, "\n", "")
		for _, answer := range appended {
			answers[string(answer)] += 1
		}
		answerList = append(answerList, answers)
		peopleCount = append(peopleCount, len(strings.Split(group, "\n")))
	}

	var varyingAnswerSum int
	var sameAnswerSum int
	for groupIndex, answers := range answerList {
		varyingAnswerSum += len(answers)
		for _, v := range answers {
			if peopleCount[groupIndex] == v {
				sameAnswerSum++
			}
		}
	}
	fmt.Println(varyingAnswerSum)
	fmt.Println(sameAnswerSum)
}
