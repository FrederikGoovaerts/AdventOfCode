package main

import (
	"aoc/util"
	"fmt"
	"math"
)

type Boss struct {
	hp  int
	dmg int
}

type GameState struct {
	hp            int
	mana          int
	bossHp        int
	bossDmg       int
	shieldCount   int
	poisonCount   int
	rechargeCount int
	manaSpent     int
}

func (s GameState) getSpellList() []Spell {
	list := make([]Spell, 0, 5)

	if s.mana >= 53 {
		list = append(list, MagicMissile)
	}
	if s.mana >= 73 {
		list = append(list, Drain)
	}
	if s.mana >= 113 && s.shieldCount == 0 {
		list = append(list, Shield)
	}
	if s.mana >= 173 && s.poisonCount == 0 {
		list = append(list, Poison)
	}
	if s.mana >= 229 && s.rechargeCount == 0 {
		list = append(list, Recharge)
	}

	return list
}

func (s GameState) didPlayerWin() bool {
	return s.bossHp <= 0
}

func (s GameState) didBossWin() bool {
	return s.hp <= 0 && !s.didPlayerWin()
}

type Spell int

const (
	MagicMissile Spell = iota
	Drain
	Shield
	Poison
	Recharge
)

func stepUntilPlayerAction(state GameState, hard bool) GameState {
	// Effects before boss turn
	if state.poisonCount > 0 {
		state.bossHp -= 3
		state.poisonCount--
	}
	if state.rechargeCount > 0 {
		state.mana += 101
		state.rechargeCount--
	}
	if state.shieldCount > 0 {
		state.shieldCount--
	}

	// Boss turn
	dmgToPlayer := state.bossDmg
	if state.shieldCount > 0 {
		dmgToPlayer = util.MaxInt(1, dmgToPlayer-7)
	}
	state.hp -= dmgToPlayer

	// Effects before player turn
	if hard {
		state.hp--
	}
	if state.hp > 0 {
		// Only if player hasn't lost, otherwise we could count this as a false win
		if state.poisonCount > 0 {
			state.bossHp -= 3
			state.poisonCount--
		}
	}
	if state.rechargeCount > 0 {
		state.mana += 101
		state.rechargeCount--
	}
	if state.shieldCount > 0 {
		state.shieldCount--
	}
	return state
}

func getLowestManaSpent(state GameState, currentBest int, hard bool) int {
	if state.didPlayerWin() {
		return util.MinInt(currentBest, state.manaSpent)
	}
	if state.manaSpent >= currentBest {
		return currentBest
	}
	if state.didBossWin() {
		return currentBest
	}

	spellList := state.getSpellList()
	best := currentBest

	for _, spell := range spellList {
		// cast spell, advance turn and recurse
		newState := state
		switch spell {
		case MagicMissile:
			newState.bossHp -= 4
			newState.mana -= 53
			newState.manaSpent += 53
		case Drain:
			newState.bossHp -= 2
			newState.hp += 2
			newState.mana -= 73
			newState.manaSpent += 73
		case Shield:
			newState.shieldCount = 6
			newState.mana -= 113
			newState.manaSpent += 113
		case Poison:
			newState.poisonCount = 6
			newState.mana -= 173
			newState.manaSpent += 173
		case Recharge:
			newState.rechargeCount = 5
			newState.mana -= 229
			newState.manaSpent += 229
		}

		if newState.didPlayerWin() {
			best = util.MinInt(best, newState.manaSpent)
			continue
		}
		newState = stepUntilPlayerAction(newState, hard)
		best = util.MinInt(best, getLowestManaSpent(newState, best, hard))
	}

	return best
}

func part1(state GameState) int {
	return getLowestManaSpent(state, math.MaxInt, false)
}

func part2(state GameState) int {
	// First turn penalty
	state.hp = state.hp - 1
	return getLowestManaSpent(state, math.MaxInt, true)
}

func main() {
	fmt.Println(part1(ex1State))
	fmt.Println(part1(ex2State))
	fmt.Println(part1(inputState))
	fmt.Println(part2(inputState))
}
