import { asIs } from "../utils/inputReader";
const input = asIs("input");
const inputParts = input.split("\n\n");

interface Interval {
  start: number;
  end: number;
}

const seedsPart1: number[] = inputParts[0].split(" ").map((n) => parseInt(n));
const seedsPart2: Interval[] = [];

for (let i = 0; i < seedsPart1.length - 1; i += 2) {
  seedsPart2.push({
    start: seedsPart1[i],
    end: seedsPart1[i] + seedsPart1[i + 1] - 1,
  });
}

const maps: { name: string; mappings: Mapping[] }[] = [];

interface Mapping {
  start: number;
  end: number;
  offset: number;
}

function fitsMapping(n: number, m: Mapping): boolean {
  return n >= m.start && n <= m.end;
}

function doMapping(n: number, m: Mapping): number {
  return n + m.offset;
}

for (const rawMap of inputParts.slice(1)) {
  const mapParts = rawMap.split("\n");
  const mapName = mapParts[0].split(" map:")[0];
  const mappings: Mapping[] = [];

  for (const mapLine of mapParts.slice(1)) {
    const mapLineParts = mapLine.split(" ");
    const destStart = parseInt(mapLineParts[0]);
    const origStart = parseInt(mapLineParts[1]);
    const rangeLength = parseInt(mapLineParts[2]);
    mappings.push({
      start: origStart,
      end: origStart + rangeLength - 1,
      offset: destStart - origStart,
    });
  }
  maps.push({ name: mapName, mappings });
}

let mappedSeedsPart1 = [...seedsPart1];

for (const map of maps) {
  const newSeeds = [];
  for (const seed of mappedSeedsPart1) {
    let mapped = false;
    for (const mapping of map.mappings) {
      if (fitsMapping(seed, mapping)) {
        newSeeds.push(doMapping(seed, mapping));
        mapped = true;
        break;
      }
    }
    if (!mapped) {
      newSeeds.push(seed);
    }
  }
  mappedSeedsPart1 = newSeeds;
}

console.log(Math.min(...mappedSeedsPart1));

let intervals = [...seedsPart2];

for (const map of maps) {
  const resultIntervals: Interval[] = [];

  let currentInterval = intervals.shift();
  while (currentInterval) {
    const matchedStartMapping = map.mappings.find((m) =>
      fitsMapping(currentInterval!.start, m)
    );

    if (!matchedStartMapping) {
      const matchedFirstMapping = map.mappings
        .filter(
          (m) =>
            m.start <= currentInterval!.end && m.start >= currentInterval!.start
        )
        .sort((a, b) => a.start - b.start)[0];

      if (!matchedFirstMapping) {
        resultIntervals.push(currentInterval);
      } else if (matchedFirstMapping.end >= currentInterval.end) {
        resultIntervals.push({
          start: currentInterval.start,
          end: matchedFirstMapping.start - 1,
        });
        resultIntervals.push({
          start: matchedFirstMapping.start + matchedFirstMapping.offset,
          end: currentInterval.end + matchedFirstMapping.offset,
        });
      } else {
        resultIntervals.push({
          start: currentInterval.start,
          end: matchedFirstMapping.start - 1,
        });
        resultIntervals.push({
          start: matchedFirstMapping.start + matchedFirstMapping.offset,
          end: matchedFirstMapping.end + matchedFirstMapping.offset,
        });
        intervals.push({
          start: matchedFirstMapping.end + 1,
          end: currentInterval.end,
        });
      }
    } else {
      if (matchedStartMapping.end >= currentInterval.end) {
        resultIntervals.push({
          start: currentInterval.start + matchedStartMapping.offset,
          end: currentInterval.end + matchedStartMapping.offset,
        });
      } else {
        resultIntervals.push({
          start: currentInterval.start + matchedStartMapping.offset,
          end: matchedStartMapping.end + matchedStartMapping.offset,
        });
        intervals.push({
          start: matchedStartMapping.end + 1,
          end: currentInterval.end,
        });
      }
    }

    currentInterval = intervals.shift();
  }

  intervals = resultIntervals;
}

console.log(Math.min(...intervals.map((i) => i.start)));
