package main

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

func (b *Blueprint) canCreate(robot Robot, ore int, clay int, obsidian int) bool {
	switch robot {
	case GeodeRobot:
		return ore >= b.geoRobotOre && obsidian >= b.geoRobotObs
	case ObsidianRobot:
		return ore >= b.obsRobotOre && clay >= b.obsRobotClay
	case ClayRobot:
		return ore >= b.clayRobotOre
	case OreRobot:
		return ore >= b.oreRobotOre
	}
	return true
}
