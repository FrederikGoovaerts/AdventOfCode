import { as2d } from "../utils/inputReader";
import {
  Direction,
  getDirectionNeighbourLoc,
  turnClockwise,
} from "../utils/neighbours";

const input = as2d("input");

let startX = -1;
let startY = -1;
let currDir: Direction = "up";

for (let y = 0; y < input.length; y++) {
  for (let x = 0; x < input[0].length; x++) {
    if (input[y][x] === "^") {
      startX = x;
      startY = y;
    }
  }
}

let currX = startX;
let currY = startY;

const visitedSquares: Set<string> = new Set();
visitedSquares.add(`${currX}|${currY}`);

let out = false;

while (!out) {
  const dir = getDirectionNeighbourLoc(currX, currY, input, currDir);

  if (dir) {
    const val = input[dir.y][dir.x];

    if (val === "#") {
      currDir = turnClockwise(currDir);
    } else {
      currX = dir.x;
      currY = dir.y;
      visitedSquares.add(`${currX}|${currY}`);
    }
  } else {
    out = true;
  }
}

console.log(visitedSquares.size);

let count = 0;

for (const sq of visitedSquares) {
  const [x, y] = sq.split("|").map(Number);

  currX = startX;
  currY = startY;
  currDir = "up";

  const visitedSquaresWithDir: Set<string> = new Set();
  visitedSquaresWithDir.add(`${currX}|${currY}|${currDir}`);

  let out = false;
  let stuck = false;

  while (!out && !stuck) {
    const dir = getDirectionNeighbourLoc(currX, currY, input, currDir);

    if (dir) {
      const val = input[dir.y][dir.x];

      if (val === "#" || (dir.y === y && dir.x === x)) {
        currDir = turnClockwise(currDir);
      } else {
        currX = dir.x;
        currY = dir.y;
      }

      if (visitedSquaresWithDir.has(`${currX}|${currY}|${currDir}`)) {
        stuck = true;
      } else {
        visitedSquaresWithDir.add(`${currX}|${currY}|${currDir}`);
      }
    } else {
      out = true;
    }
  }

  if (stuck) {
    count++;
  }
}

console.log(count);
