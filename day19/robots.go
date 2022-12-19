package main

import "math"

type Robot int

const (
	OreRobot Robot = iota
	ClayRobot
	ObsidianRobot
	GeodeRobot
	None
)

type Blueprint struct {
	id           int
	oreRobotOre  int
	clayRobotOre int
	obsRobotOre  int
	obsRobotClay int
	geoRobotOre  int
	geoRobotObs  int
}

func (b Blueprint) getGeoRatio() float64 {
	return float64(b.geoRobotObs) / float64(b.geoRobotOre)
}

func (b Blueprint) getObsRatio() float64 {
	return float64(b.obsRobotClay) / float64(b.obsRobotOre)
}

type Inventory struct {
	oreRobot      int
	ore           int
	clayRobot     int
	clay          int
	obsidianRobot int
	obsidian      int
	geodeRobot    int
	geode         int
}

func (inv Inventory) mine() Inventory {
	inv.ore += inv.oreRobot
	inv.clay += inv.clayRobot
	inv.obsidian += inv.obsidianRobot
	inv.geode += inv.geodeRobot

	return inv
}

func (inv Inventory) makeRobot(bl Blueprint, robot Robot) Inventory {
	switch robot {
	case None:
		return inv
	case OreRobot:
		inv.oreRobot++
		inv.ore -= bl.oreRobotOre
		return inv
	case ClayRobot:
		inv.clayRobot++
		inv.ore -= bl.clayRobotOre
		return inv
	case ObsidianRobot:
		inv.obsidianRobot++
		inv.clay -= bl.obsRobotClay
		inv.ore -= bl.obsRobotOre
		return inv
	case GeodeRobot:
		inv.geodeRobot++
		inv.obsidian -= bl.geoRobotObs
		inv.ore -= bl.geoRobotOre
		return inv
	}
	panic("That's not a robot")
}

func (inv *Inventory) getGeoRatio(b Blueprint) float64 {
	return (float64(inv.obsidianRobot) / float64(inv.oreRobot)) / b.getGeoRatio()
}
func (inv *Inventory) getObsRatio(b Blueprint) float64 {
	return (float64(inv.clayRobot) / float64(inv.oreRobot)) / b.getObsRatio()
}
func (inv *Inventory) getClayRatio(b Blueprint) float64 {
	return (float64(inv.clayRobot) / float64(inv.oreRobot)) / (float64(b.clayRobotOre) / float64(b.oreRobotOre))
}

func (inv *Inventory) turnsUntilGeodeRobot(blueprint Blueprint) int {
	if inv.obsidianRobot == 0 {
		return math.MaxInt
	}
	turnsForOre := math.Ceil(float64(blueprint.geoRobotOre-inv.ore) / float64(inv.oreRobot))
	turnsForObs := math.Ceil(float64(blueprint.geoRobotObs-inv.obsidian) / float64(inv.obsidianRobot))
	return int(math.Max(turnsForOre, turnsForObs))
}
func (inv *Inventory) turnsUntilObsidianRobot(blueprint Blueprint) int {
	if inv.clayRobot == 0 {
		return math.MaxInt
	}
	turnsForOre := math.Ceil(float64(blueprint.obsRobotOre-inv.ore) / float64(inv.oreRobot))
	turnsForClay := math.Ceil(float64(blueprint.obsRobotClay-inv.clay) / float64(inv.clayRobot))
	return int(math.Max(turnsForOre, turnsForClay))
}

func (inv *Inventory) canCreate(robot Robot, blueprint Blueprint) bool {
	switch robot {
	case GeodeRobot:
		return inv.ore >= blueprint.geoRobotOre && inv.obsidian >= blueprint.geoRobotObs
	case ObsidianRobot:
		return inv.ore >= blueprint.obsRobotOre && inv.clay >= blueprint.obsRobotClay
	case ClayRobot:
		return inv.ore >= blueprint.clayRobotOre
	case OreRobot:
		return inv.ore >= blueprint.oreRobotOre
	}
	return true
}

// I don't even know what I'm doing

// func (inv *Inventory) shouldCreate(blueprint Blueprint) []Robot {
// 	if inv.canCreate(GeodeRobot, blueprint) {
// 		return []Robot{GeodeRobot}
// 	} else if inv.canCreate(ObsidianRobot, blueprint) {
// 		// hypotheticalInv := inv.makeRobot(blueprint, ObsidianRobot)
// 		// if inv.turnsUntilGeodeRobot(blueprint) >= hypotheticalInv.turnsUntilGeodeRobot(blueprint) && inv.getGeoRatio(blueprint) < 1 {
// 		// 	return []Robot{ObsidianRobot}
// 		// }
// 		return []Robot{None, ObsidianRobot}
// 	} else if inv.canCreate(ClayRobot, blueprint) {
// 		// hypotheticalInv := inv.makeRobot(blueprint, ClayRobot)
// 		// geodeTurnSame := inv.turnsUntilGeodeRobot(blueprint) >= hypotheticalInv.turnsUntilGeodeRobot(blueprint)
// 		// obsTurnSame := inv.turnsUntilObsidianRobot(blueprint) >= hypotheticalInv.turnsUntilObsidianRobot(blueprint)
// 		// obsRatio := hypotheticalInv.getObsRatio(blueprint)
// 		// if geodeTurnSame && obsTurnSame && obsRatio < 1 {
// 		// 	return []Robot{ClayRobot}
// 		// }
// 		return []Robot{None, ClayRobot}
// 	} else if inv.canCreate(OreRobot, blueprint) {
// 		// hypotheticalInv := inv.makeRobot(blueprint, OreRobot)
// 		// geodeTurnSame := inv.turnsUntilGeodeRobot(blueprint) >= hypotheticalInv.turnsUntilGeodeRobot(blueprint)
// 		// obsTurnSame := inv.turnsUntilObsidianRobot(blueprint) >= hypotheticalInv.turnsUntilObsidianRobot(blueprint)
// 		// if geodeTurnSame && obsTurnSame && hypotheticalInv.getClayRatio(blueprint) <= 1 {
// 		// 	return OreRobot
// 		// }
// 		return []Robot{None, OreRobot}
// 	}
// 	return []Robot{None}
// }

// func (inv *Inventory) canCreate(blueprint Blueprint) []Robot {

// 	if inv.ore >= blueprint.geoRobotOre && inv.obsidian >= blueprint.geoRobotObs {
// 		return GeodeRobot
// 	} else if inv.ore >= blueprint.obsRobotOre && inv.clay >= blueprint.obsRobotClay {
// 		return ObsidianRobot
// 	} else if inv.ore >= blueprint.clayRobotOre {
// 		return ClayRobot
// 	} else if inv.ore >= blueprint.oreRobotOre {
// 		return OreRobot
// 	} else {
// 		return None
// 	}
// }

func (inv *Inventory) shouldCreate(blueprint Blueprint) []Robot {
	result := make([]Robot, 0)
	if inv.canCreate(GeodeRobot, blueprint) {
		return []Robot{GeodeRobot}
	} else if inv.canCreate(ObsidianRobot, blueprint) {
		return []Robot{ObsidianRobot}
	} else if inv.canCreate(ClayRobot, blueprint) {
		result = append(result, ClayRobot)
	}
	if inv.canCreate(OreRobot, blueprint) {
		result = append(result, OreRobot)
	}
	result = append(result, None)
	return result
}
