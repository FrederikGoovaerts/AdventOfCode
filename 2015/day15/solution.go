package main

import (
	"aoc/util"
	"fmt"
	"regexp"
)

type Ingredient struct {
	name string
	cap  int
	dur  int
	flav int
	tex  int
	cal  int
}

func getTotal(ingr []Ingredient, amount map[string]int, checkCalories bool) int {
	cap := 0
	dur := 0
	flav := 0
	tex := 0
	cal := 0

	for _, i := range ingr {
		am, present := amount[i.name]
		if present {
			cap += i.cap * am
			dur += i.dur * am
			flav += i.flav * am
			tex += i.tex * am
			cal += i.cal * am
		}
	}

	if checkCalories && cal != 500 {
		return 0
	}
	return util.MaxInt(0, cap) * util.MaxInt(0, dur) * util.MaxInt(0, flav) * util.MaxInt(0, tex)
}

var inputRegex = regexp.MustCompile("^(.*): capacity (.*), durability (.*), flavor (.*), texture (.*), calories (.*)$")

func parse(lines []string) []Ingredient {
	ingredients := make([]Ingredient, 0, len(lines))

	for _, line := range lines {
		matches := inputRegex.FindStringSubmatch(line)
		capacity := util.StringToInt(matches[2])
		durability := util.StringToInt(matches[3])
		flavor := util.StringToInt(matches[4])
		texture := util.StringToInt(matches[5])
		calories := util.StringToInt(matches[6])

		ingredients = append(ingredients, Ingredient{matches[1], capacity, durability, flavor, texture, calories})
	}

	return ingredients
}

func getCopy(theMap map[string]int) map[string]int {
	copy := make(map[string]int, len(theMap))
	for k, v := range theMap {
		copy[k] = v
	}
	return copy
}

func getHighest(scoopsRemaining int, ingrRemaining []string, ratio map[string]int, ingr []Ingredient, checkCalories bool) int {
	if len(ingrRemaining) == 1 {
		newRatio := getCopy(ratio)
		newRatio[ingrRemaining[0]] = scoopsRemaining
		return getTotal(ingr, newRatio, checkCalories)
	} else if scoopsRemaining == 0 {
		return getTotal(ingr, ratio, checkCalories)
	} else {
		best := 0
		nextIngr, restIngr := ingrRemaining[0], ingrRemaining[1:]
		for amount := 0; amount <= scoopsRemaining; amount++ {
			newRatio := getCopy(ratio)
			newRatio[nextIngr] = amount
			best = util.MaxInt(best, getHighest(scoopsRemaining-amount, restIngr, newRatio, ingr, checkCalories))
		}
		return best
	}
}

func part1(ingredients []Ingredient) int {
	names := make([]string, 0, len(ingredients))
	for _, i := range ingredients {
		names = append(names, i.name)
	}

	return getHighest(100, names, make(map[string]int), ingredients, false)
}

func part2(ingredients []Ingredient) int {
	names := make([]string, 0, len(ingredients))
	for _, i := range ingredients {
		names = append(names, i.name)
	}

	return getHighest(100, names, make(map[string]int), ingredients, true)
}
func main() {
	// input := util.FileAsLines("ex1")
	input := util.FileAsLines("input")
	ingredients := parse(input)

	fmt.Println(part1(ingredients))
	fmt.Println(part2(ingredients))
}
