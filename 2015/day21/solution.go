package main

import (
	"aoc/util"
	"fmt"
	"math"
)

func playerWins(boss Boss, items []Item, playerHp int) bool {
	playerArmor := 0
	playerDmg := 0

	for _, item := range items {
		playerArmor += item.armor
		playerDmg += item.dmg
	}

	bossHit := util.MaxInt(1, boss.dmg-playerArmor)
	playerHit := util.MaxInt(1, playerDmg-boss.armor)

	deathTurns := math.Ceil(float64(playerHp) / float64(bossHit))
	winTurns := math.Ceil(float64(boss.hp) / float64(playerHit))
	return winTurns <= deathTurns
}

func without(list []Item, removed Item) []Item {
	result := make([]Item, 0)
	for _, el := range list {
		if el != removed {
			result = append(result, el)
		}
	}
	return result
}

func part1(boss Boss) int {
	cheapest := math.MaxInt

	armorChoices := append(armor, Item{0, 0, 0})
	ringChoices := append(rings, Item{0, 0, 0}, Item{0, 0, 0})
	for _, weapon := range weapons {
		for _, armor := range armorChoices {
			for _, ring1 := range ringChoices {
				for _, ring2 := range without(ringChoices, ring1) {
					cost := weapon.cost + armor.cost + ring1.cost + ring2.cost
					if cost < cheapest {
						if playerWins(boss, []Item{weapon, armor, ring1, ring2}, 100) {
							cheapest = cost
						}
					}
				}
			}
		}
	}

	return cheapest
}

func part2(boss Boss) int {
	mostExpensive := 0

	armorChoices := append(armor, Item{0, 0, 0})
	ringChoices := append(rings, Item{0, 0, 0}, Item{0, 0, 0})
	for _, weapon := range weapons {
		for _, armor := range armorChoices {
			for _, ring1 := range ringChoices {
				for _, ring2 := range without(ringChoices, ring1) {
					cost := weapon.cost + armor.cost + ring1.cost + ring2.cost
					if cost > mostExpensive {
						if !playerWins(boss, []Item{weapon, armor, ring1, ring2}, 100) {
							mostExpensive = cost
						}
					}
				}
			}
		}
	}

	return mostExpensive
}

func main() {
	fmt.Println(part1(boss))
	fmt.Println(part2(boss))
}
