package main

type Block string

const (
	Flat   Block = "-"
	Plus   Block = "+"
	Hook   Block = "hook hook ay look at them moving eyes"
	Long   Block = "|"
	Square Block = "o"
)

var blockOrder = []Block{Flat, Plus, Hook, Long, Square}

func getBlockCoords(block Block) []string {
	switch block {
	case Flat:
		return []string{"0 0", "1 0", "2 0", "3 0"}
	case Plus:
		return []string{"0 1", "1 0", "1 1", "1 2", "2 1"}
	case Hook:
		return []string{"0 0", "1 0", "2 0", "2 1", "2 2"}
	case Long:
		return []string{"0 0", "0 1", "0 2", "0 3"}
	case Square:
		return []string{"0 0", "0 1", "1 0", "1 1"}
	}
	panic("Not a valid block name")
}

func blockWidth(block Block) int {
	if block == Flat {
		return 4
	} else if block == Plus || block == Hook {
		return 3
	} else if block == Square {
		return 2
	} else {
		return 1
	}
}

func blockHeight(block Block) int {
	if block == Long {
		return 4
	} else if block == Plus || block == Hook {
		return 3
	} else if block == Square {
		return 2
	} else {
		return 1
	}
}

func getNthBlock(turn int) ([]string, Block) {
	name := blockOrder[turn%len(blockOrder)]
	return getBlockCoords(name), name
}
