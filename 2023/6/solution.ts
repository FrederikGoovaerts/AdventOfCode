interface Race {
  time: number;
  dist: number;
}

const racesEx1: Race[] = [
  { time: 7, dist: 9 },
  { time: 15, dist: 40 },
  { time: 30, dist: 200 },
];
const racesInput: Race[] = [
  { time: 60, dist: 475 },
  { time: 94, dist: 2138 },
  { time: 78, dist: 1015 },
  { time: 82, dist: 1650 },
];
const racesEx1P2: Race[] = [{ time: 71530, dist: 940200 }];
const racesInputP2: Race[] = [{ time: 60947882, dist: 475213810151650 }];

function doRaces(races: Race[]) {
  const ways: number[] = [];

  function crosses(time: number, dist: number, speed: number): boolean {
    return time * speed > dist;
  }

  for (const race of races) {
    let currentWays = 0;
    for (let i = 1; i < race.time; i++) {
      if (crosses(race.time - i, race.dist, i)) {
        currentWays++;
      }
    }
    ways.push(currentWays);
  }

  console.log(ways.reduce((a, b) => a * b, 1));
}

// doRaces(racesEx1);
// doRaces(racesEx1P2);
doRaces(racesInput);
doRaces(racesInputP2);
