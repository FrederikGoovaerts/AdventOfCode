import { asList } from "../utils/inputReader";
import {
  Direction,
  getDirectionNeighboursLocs,
  getNeighboursLocs,
  reverseDirection,
} from "../utils/neighbours";
const input = asList("input");

type mapSymbols = "." | "S" | "|" | "-" | "J" | "L" | "F" | "7";

function connects(char: string, dir: Direction): boolean {
  if (dir === "up") {
    return ["|", "J", "L"].includes(char);
  } else if (dir === "down") {
    return ["|", "F", "7"].includes(char);
  } else if (dir === "left") {
    return ["-", "J", "7"].includes(char);
  } else if (dir === "right") {
    return ["-", "F", "L"].includes(char);
  }
  return false;
}

const map: mapSymbols[][] = [];
const path: { x: number; y: number }[] = [];

input.forEach((line, y) => {
  const splitLine = line.split("") as mapSymbols[];
  map.push(splitLine);

  const sIndex = splitLine.indexOf("S");

  if (sIndex > -1) {
    path.push({ x: sIndex, y: y });
  }
});

// Choose random starting direction of path
for (const [dir, loc] of Object.entries(
  getDirectionNeighboursLocs(path[0].x, path[0].y, map)
)) {
  const dirPath: { x: number; y: number }[] = [path[0]];
  const char = map[loc[1]][loc[0]];
  if (!connects(char, reverseDirection(dir as Direction))) {
    continue;
  }
  let curr = { x: loc[0], y: loc[1] };
  dirPath.push(curr);

  let expanded = true;

  while (expanded) {
    expanded = false;
    for (const [dir, loc] of Object.entries(
      getDirectionNeighboursLocs(curr.x, curr.y, map)
    )) {
      const char = map[loc[1]][loc[0]];
      if (
        connects(map[curr.y][curr.x], dir as Direction) &&
        connects(char, reverseDirection(dir as Direction)) &&
        !dirPath.some((p) => p.x === loc[0] && p.y === loc[1])
      ) {
        curr = { x: loc[0], y: loc[1] };
        dirPath.push(curr);
        expanded = true;
        break;
      }
    }
  }

  if (dirPath.length > path.length) {
    path.length = 0;
    path.push(...dirPath);
  }
}

console.log(path.length / 2);

// This starts out as the first F corner we encounter
let currPathElem = { x: -1, y: -1 };
// We'll traverse the path starting from a corner and looking down
let lookDir: Direction = "down";

const cleanMap: ("." | "P" | "I" | "O")[][] = [];
for (let y = 0; y < map.length; y++) {
  const cleanMapLine: ("." | "P" | "I" | "O")[] = [];
  for (let x = 0; x < map[0].length; x++) {
    if (path.some((p) => p.x === x && p.y === y)) {
      cleanMapLine.push("P");
      if (currPathElem.x === -1) {
        currPathElem = { x, y };
      }
    } else {
      cleanMapLine.push(".");
    }
  }
  cleanMap.push(cleanMapLine);
}

const firstCornerIndex = path.findIndex(
  (p) => p.x === currPathElem.x && p.y === currPathElem.y
);

const applyPathToMap = (el: { x: number; y: number }, dir: Direction) => {
  switch (dir) {
    case "up":
      {
        if (cleanMap[el.y]?.[el.x + 1] === ".") {
          cleanMap[el.y][el.x + 1] = "O";
        }
        if (cleanMap[el.y]?.[el.x - 1] === ".") {
          cleanMap[el.y][el.x - 1] = "I";
        }
      }
      break;
    case "down":
      {
        if (cleanMap[el.y]?.[el.x + 1] === ".") {
          cleanMap[el.y][el.x + 1] = "I";
        }
        if (cleanMap[el.y]?.[el.x - 1] === ".") {
          cleanMap[el.y][el.x - 1] = "O";
        }
      }
      break;
    case "left":
      {
        if (cleanMap[el.y - 1]?.[el.x] === ".") {
          cleanMap[el.y - 1][el.x] = "O";
        }
        if (cleanMap[el.y + 1]?.[el.x] === ".") {
          cleanMap[el.y + 1][el.x] = "I";
        }
      }
      break;
    case "right":
      {
        if (cleanMap[el.y - 1]?.[el.x] === ".") {
          cleanMap[el.y - 1][el.x] = "I";
        }
        if (cleanMap[el.y + 1]?.[el.x] === ".") {
          cleanMap[el.y + 1][el.x] = "O";
        }
      }
      break;
  }
};

applyPathToMap(currPathElem, lookDir);

let last = currPathElem;
for (
  let i = (firstCornerIndex + 1) % path.length;
  i !== firstCornerIndex;
  i = (i + 1) % path.length
) {
  currPathElem = path[i];
  if (currPathElem.x > last.x) {
    lookDir = "right";
  } else if (currPathElem.x < last.x) {
    lookDir = "left";
  } else if (currPathElem.y < last.y) {
    lookDir = "up";
  } else if (currPathElem.y > last.y) {
    lookDir = "down";
  }

  applyPathToMap(currPathElem, lookDir);
  const next = path[(i + 1) % path.length];
  let dirToNext: Direction;
  if (next.x > currPathElem.x) {
    dirToNext = "right";
  } else if (next.x < currPathElem.x) {
    dirToNext = "left";
  } else if (next.y < currPathElem.y) {
    dirToNext = "up";
  } else {
    dirToNext = "down";
  }

  applyPathToMap(currPathElem, dirToNext);

  last = currPathElem;
}

// Flood fill I
for (let y = 0; y < map.length; y++) {
  for (let x = 0; x < map[0].length; x++) {
    if (cleanMap[y][x] === "I") {
      for (const [xN, yN] of getNeighboursLocs(x, y, cleanMap, false)) {
        if (cleanMap[yN][xN] === ".") {
          cleanMap[yN][xN] = "I";
        }
      }
    }
  }
}

let iCount = 0;
for (let y = 0; y < map.length; y++) {
  // This can be used to visually debug
  // console.log(cleanMap[y].join(""));
  iCount += cleanMap[y].filter((v) => v === "I").length;
}
console.log(iCount);
