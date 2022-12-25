package main

import (
	"aoc/utils"
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

func getPossibleTranslations(input [][]string, previousTranslations []string) []string {
	curr := make([]string, 0)
	for _, el := range input[0] {
		if !utils.ContainsString(previousTranslations, el) {
			curr = append(curr, el)
		}
	}
	for i := 1; i < len(input); i++ {
		ingredients := input[i]
		newCurr := make([]string, 0)
		for _, el := range curr {
			if utils.ContainsString(ingredients, el) && !utils.ContainsString(previousTranslations, el) {
				newCurr = append(newCurr, el)
			}
		}
		curr = newCurr
	}
	return curr
}

func main() {
	ingredients := make([]string, 0)
	allergenCandidates := make(map[string][][]string, 0)
	allergenTranslation := make(map[string]string, 0)
	dat, err := ioutil.ReadFile("input")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(strings.TrimSpace(string(dat)), "\n")

	for _, line := range lines {
		splitLine := strings.Split(line, " (contains ")
		newCandidates := strings.Fields(splitLine[0])
		ingredients = append(ingredients, newCandidates...)
		allergens := strings.Split(splitLine[1][:len(splitLine[1])-1], ", ")
		for _, allergen := range allergens {
			alCan, found := allergenCandidates[allergen]

			if !found {
				alCan = make([][]string, 0)
			}
			alCan = append(alCan, newCandidates)
			allergenCandidates[allergen] = alCan
		}
	}

	translatedSet := make([]string, 0)
	allergenSet := make([]string, 0)
	for len(allergenCandidates) > 0 {
		for k, v := range allergenCandidates {
			trans := getPossibleTranslations(v, translatedSet)
			if len(trans) == 1 {
				translatedSet = append(translatedSet, trans[0])
				allergenSet = append(allergenSet, k)
				allergenTranslation[k] = trans[0]
				delete(allergenCandidates, k)
			}
		}
	}

	count := 0
	for _, ing := range ingredients {
		if !utils.ContainsString(translatedSet, ing) {
			count++
		}
	}
	fmt.Println(count)

	sort.Strings(allergenSet)
	result := ""
	for _, el := range allergenSet {
		result += allergenTranslation[el] + ","
	}
	fmt.Println(result[:len(result)-1])

}
