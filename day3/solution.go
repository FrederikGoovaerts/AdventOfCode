package main

import (
	"fmt"
	"os"
	"strings"
)

func priority(a rune) int {
	code := int(a)
	if code > 96 {
		return code - 96
	}
	// ASCII offset AND A starts at 27 prio
	return code - 38
}

type Pack struct {
	contents string
}

func (p Pack) first() string {
	return p.contents[0 : len(p.contents)/2]
}

func (p Pack) second() string {
	return p.contents[len(p.contents)/2:]
}

func (p Pack) getDuplicate() rune {
	for _, inFirst := range p.first() {
		for _, inSecond := range p.second() {
			if inFirst == inSecond {
				return inFirst
			}
		}
	}
	panic("No duplicate found.")
}

func getBadge(pack1 Pack, pack2 Pack, pack3 Pack) rune {
	for _, ref := range pack1.contents {
		inPack2 := false
		for _, el := range pack2.contents {
			if ref == el {
				inPack2 = true
				break
			}
		}
		if inPack2 {
			for _, el := range pack3.contents {
				if ref == el {
					return ref
				}
			}
		}

	}
	panic("No badge found.")
}

func part1(packs []Pack) int {
	result := 0
	for _, pack := range packs {
		result += priority(pack.getDuplicate())
	}
	return result
}

func part2(packs []Pack) int {
	result := 0
	for i := 0; i < len(packs); i += 3 {
		badge := getBadge(packs[i], packs[i+1], packs[i+2])
		result += priority(badge)
	}
	return result
}

func main() {
	dat, err := os.ReadFile("input")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(strings.TrimSpace(string(dat)), "\n")
	packs := []Pack{}

	for _, line := range lines {
		if line != "" {
			packs = append(packs, Pack{line})
		}
	}
	fmt.Println(part1(packs))
	fmt.Println(part2(packs))
}
