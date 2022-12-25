package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type player struct {
	name  string
	cards []int
}

type result struct {
	name  string
	score int
}

func (p player) countScore() int {
	sum := 0
	for i := 0; i < len(p.cards); i++ {
		sum += (p.cards[i] * (len(p.cards) - i))
	}
	return sum
}

func (p player) hasLost() bool {
	return len(p.cards) == 0
}

func (p *player) popCard() int {
	card, remaining := p.cards[0], p.cards[1:]
	p.cards = remaining
	return card
}

func (p *player) addCards(cards []int) {
	p.cards = append(p.cards, cards...)
}

func (p player) copy() player {
	newCards := make([]int, len(p.cards))
	copy(newCards, p.cards)
	return player{p.name, newCards}
}

func (p player) copyAmount(nb int) player {
	newPlayer := p.copy()
	newPlayer.cards = newPlayer.cards[0:nb]
	return newPlayer
}

func (p player) serializeCards() string {
	res := ""
	for _, el := range p.cards {
		res += fmt.Sprint(el) + ","
	}
	return res
}

var results map[string]result

func playGame(players map[string]player, recursive bool) (string, int) {

	getCurrentState := func() string {
		return players["Player 1"].serializeCards() + "|" + players["Player 2"].serializeCards()
	}

	initialGame := getCurrentState()
	res, found := results[initialGame]
	if found {
		return res.name, res.score
	}

	visitedStates := make(map[string]struct{}, 0)
	hasWinner := func() bool {
		for _, p := range players {
			if p.hasLost() {
				return true
			}
		}
		return false
	}

	for !hasWinner() {
		// Check infinite game
		if recursive {
			state := getCurrentState()
			_, visited := visitedStates[state]
			if visited {
				return "Player 1", players["Player 1"].countScore()
			}
			visitedStates[state] = struct{}{}
		}

		turnMap := make(map[string]int, 0)
		// Only used in recursive mode but computed nonetheless
		descend := true
		for _, p := range players {
			nb := (&p).popCard()
			turnMap[p.name] = nb
			players[p.name] = p
			if nb > len(p.cards) {
				descend = false
			}
		}

		winner := ""
		if !recursive || !descend {
			max := -1
			for name, score := range turnMap {
				if score > max {
					winner = name
					max = score
				}
			}

		} else {
			recursePlayers := make(map[string]player, 0)
			for k, v := range players {
				recursePlayers[k] = v.copyAmount(turnMap[k])
			}
			winner, _ = playGame(recursePlayers, true)
		}
		cards := make([]int, 0)
		if winner == "Player 1" {
			cards = append(cards, turnMap["Player 1"])
			cards = append(cards, turnMap["Player 2"])
		} else {
			cards = append(cards, turnMap["Player 2"])
			cards = append(cards, turnMap["Player 1"])
		}
		w := players[winner]
		(&w).addCards(cards)
		players[winner] = w
	}

	for _, p := range players {
		if !p.hasLost() {
			if recursive {
				results[initialGame] = result{p.name, p.countScore()}
			}
			return p.name, p.countScore()
		}
	}
	panic("No winner, but game ended.")
}

func main() {
	results = make(map[string]result)
	dat, err := ioutil.ReadFile("input")
	if err != nil {
		panic(err)
	}

	playerStrings := strings.Split(strings.TrimSpace(string(dat)), "\n\n")

	players := make(map[string]player, 0)
	for _, string := range playerStrings {
		splitString := strings.Split(string, "\n")
		playerName, cardStrings := splitString[0], splitString[1:]
		playerName = playerName[:len(playerName)-1]
		cards := make([]int, 0)
		for _, cardString := range cardStrings {
			val, _ := strconv.Atoi(cardString)
			cards = append(cards, val)
		}
		players[playerName] = player{playerName, cards}
	}

	part1Players := make(map[string]player, 0)
	part2Players := make(map[string]player, 0)
	for k, v := range players {
		part1Players[k] = v.copy()
		part2Players[k] = v.copy()
	}
	_, score := playGame(part1Players, false)
	fmt.Println(score)
	_, score = playGame(part2Players, true)
	fmt.Println(score)

}
