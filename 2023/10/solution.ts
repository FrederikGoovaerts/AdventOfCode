import { asList } from "../utils/inputReader";
import {
  Direction,
  getDirectionNeighboursLocs,
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
let curr: { x: number; y: number } = { x: -1, y: -1 };
// for (const [dir, loc] of Object.entries(
//   getDirectionNeighboursLocs(path[0].x, path[0].y, map)
// )) {
//   const char = map[loc[1]][loc[0]];
//   if (connects(char, reverseDirection(dir as Direction))) {
//     curr = { x: loc[0], y: loc[1] };
//     path.push(curr);
//     break;
//   }
// }
// TODO: Clean this up later to make it complete.
// FOR INPUT THIS IS CORRECT
curr = { x: path[0].x, y: path[0].y + 1 };
path.push(curr);

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
      !path.some((p) => p.x === loc[0] && p.y === loc[1])
    ) {
      curr = { x: loc[0], y: loc[1] };
      path.push(curr);
      expanded = true;
      break;
    }
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

let iCount = 0;
for (let y = 0; y < map.length; y++) {
  console.log(cleanMap[y].join(""));
  iCount += cleanMap[y].filter((v) => v === "I").length;
}
console.log(iCount, "(Currently not floodfilled)");
// 567 with some manual checks lol
