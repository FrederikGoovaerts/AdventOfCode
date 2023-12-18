import { asList } from "../utils/inputReader";
import { Direction, getNeighboursLocs } from "../utils/neighbours";

const input = asList("input");

type Dir = "R" | "D" | "L" | "U";

function getDirOffset(dir: Dir): { x: number; y: number } {
  switch (dir) {
    case "R":
      return { x: 1, y: 0 };
    case "D":
      return { x: 0, y: 1 };
    case "L":
      return { x: -1, y: 0 };
    case "U":
      return { x: 0, y: -1 };
  }
}

let maxX = 0;
let maxY = 0;

function createCcwPath(): { x: number; y: number }[] {
  const result: { x: number; y: number }[] = [];
  let currX = 0;
  let currY = 0;
  let minX = 0;
  let minY = 0;

  result.push({ x: currX, y: currY });

  for (const line of input) {
    const split = line.split(" ");
    const direction = split[0] as Dir;
    const dirOffset = getDirOffset(direction);
    const dist = parseInt(split[1]);
    for (let i = 0; i < dist; i++) {
      currX += dirOffset.x;
      currY += dirOffset.y;
      if (!result.some((p) => p.x === currX && p.y === currY)) {
        result.push({ x: currX, y: currY });
      }
      minX = Math.min(minX, currX);
      minY = Math.min(minY, currY);
      maxX = Math.max(maxX, currX);
      maxY = Math.max(maxY, currY);
    }
  }

  for (const el of result) {
    el.x -= minX;
    el.y -= minY;
  }
  maxX -= minX;
  maxY -= minY;

  return result;
}

const path = createCcwPath().reverse();

const cleanMap: ("." | "P" | "I" | "O")[][] = [];

// This starts out as the first F corner we encounter
let currPathElem = {
  y: 0,
  x: Math.min(...path.filter((p) => p.y === 0).map((p) => p.x)),
};

for (let y = 0; y <= maxY; y++) {
  const line: ("." | "P" | "I" | "O")[] = [];
  for (let x = 0; x <= maxX; x++) {
    if (path.some((p) => p.x === x && p.y === y)) {
      line.push("P");
    } else {
      line.push(".");
    }
  }
  cleanMap.push(line);
}

// We'll traverse the path starting from a corner and looking down
let lookDir: Direction = "down";

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
for (let y = 0; y < cleanMap.length; y++) {
  for (let x = 0; x < cleanMap[0].length; x++) {
    if (cleanMap[y][x] === "I") {
      for (const n of getNeighboursLocs(x, y, cleanMap, false)) {
        if (cleanMap[n.y][n.x] === ".") {
          cleanMap[n.y][n.x] = "I";
        }
      }
    }
  }
}

let iCount = 0;
for (let y = 0; y < cleanMap.length; y++) {
  // This can be used to visually debug
  // console.log(cleanMap[y].join(""));
  iCount += cleanMap[y].filter((v) => v === "I").length;
}

console.log(iCount + path.length);
