import { cloneDeep } from "lodash";
import { asList } from "../utils/inputReader";
const input = asList("input").map((line) => line.split(""));

const part1Map = cloneDeep(input);

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

function getNorthLoad(map: string[][]): number {
  let result = 0;
  for (let y = 0; y < map.length; y++) {
    result += map[y].filter((v) => v === "O").length * (map.length - y);
  }
  return result;
}

shiftNorth(part1Map);

console.log(getNorthLoad(part1Map));
