import { cloneDeep, isEqual } from "lodash";
import { as2d } from "../utils/inputReader";
const input = as2d("input");
const CYCLE_TOTAL = 1_000_000_000;

const part1Map = cloneDeep(input);
const part2Map = cloneDeep(input);

function shiftNorth(map: string[][]): void {
  for (let x = 0; x < map[0].length; x++) {
    let currStart = 0;
    let currEnd = 0;
    let count = 0;
    while (currStart < map.length) {
      let currEl = map[currEnd][x];
      while (currEl !== "#" && currEnd < map.length) {
        if (currEl === "O") {
          count++;
        }
        currEnd++;
        if (currEnd < map.length) {
          currEl = map[currEnd][x];
        }
      }
      for (let i = currStart; i < currEnd; i++) {
        if (i - currStart < count) {
          map[i][x] = "O";
        } else {
          map[i][x] = ".";
        }
      }
      currEnd++;
      currStart = currEnd;
      count = 0;
    }
  }
}

function shiftWest(map: string[][]): void {
  for (let y = 0; y < map.length; y++) {
    let currStart = 0;
    let currEnd = 0;
    let count = 0;
    while (currStart < map[0].length) {
      let currEl = map[y][currEnd];
      while (currEl !== "#" && currEnd < map.length) {
        if (currEl === "O") {
          count++;
        }
        currEnd++;
        if (currEnd < map[0].length) {
          currEl = map[y][currEnd];
        }
      }
      for (let i = currStart; i < currEnd; i++) {
        if (i - currStart < count) {
          map[y][i] = "O";
        } else {
          map[y][i] = ".";
        }
      }
      currEnd++;
      currStart = currEnd;
      count = 0;
    }
  }
}

function shiftSouth(map: string[][]): void {
  for (let x = 0; x < map[0].length; x++) {
    let currStart = map.length - 1;
    let currEnd = currStart;
    let count = 0;
    while (currStart >= 0) {
      let currEl = map[currEnd][x];
      while (currEl !== "#" && currEnd >= 0) {
        if (currEl === "O") {
          count++;
        }
        currEnd--;
        if (currEnd >= 0) {
          currEl = map[currEnd][x];
        }
      }
      for (let i = currStart; i > currEnd; i--) {
        if (currStart - i < count) {
          map[i][x] = "O";
        } else {
          map[i][x] = ".";
        }
      }
      currEnd--;
      currStart = currEnd;
      count = 0;
    }
  }
}

function shiftEast(map: string[][]): void {
  for (let y = 0; y < map.length; y++) {
    let currStart = map[0].length - 1;
    let currEnd = currStart;
    let count = 0;
    while (currStart >= 0) {
      let currEl = map[y][currEnd];
      while (currEl !== "#" && currEnd >= 0) {
        if (currEl === "O") {
          count++;
        }
        currEnd--;
        if (currEnd >= 0) {
          currEl = map[y][currEnd];
        }
      }
      for (let i = currStart; i > currEnd; i--) {
        if (currStart - i < count) {
          map[y][i] = "O";
        } else {
          map[y][i] = ".";
        }
      }
      currEnd--;
      currStart = currEnd;
      count = 0;
    }
  }
}

function getNorthLoad(map: string[][]): number {
  let result = 0;
  for (let y = 0; y < map.length; y++) {
    result += map[y].filter((v) => v === "O").length * (map.length - y);
  }
  return result;
}

function getCycleCycleSize(cycles: number[]): number {
  const lastIndex = cycles.length - 1;
  const lastEl = cycles[lastIndex];

  for (let i = lastIndex - 1; i >= 0; i--) {
    if (cycles[i] === lastEl) {
      const size = lastIndex - i;

      const lastSlice = cycles.slice(i + 1);
      const otherSlice = cycles.slice(i - size + 1, i + 1);
      if (isEqual(lastSlice, otherSlice)) {
        return size;
      }
    }
  }

  return -1;
}

shiftNorth(part1Map);
console.log(getNorthLoad(part1Map));

const cycleResults: number[] = [];
const CYCLE_DETECTION_RANGE = 1000;

for (let i = 0; i < CYCLE_DETECTION_RANGE; i++) {
  shiftNorth(part2Map);
  shiftWest(part2Map);
  shiftSouth(part2Map);
  shiftEast(part2Map);

  cycleResults.push(getNorthLoad(part2Map));
}

const size = getCycleCycleSize(cycleResults);
const lastCycleStart = CYCLE_DETECTION_RANGE - size;
const resultOffset = (CYCLE_TOTAL - lastCycleStart) % size;

// -1 because of zero-indexing of array
console.log(cycleResults[lastCycleStart + resultOffset - 1]);
