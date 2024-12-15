import { asIs } from "../utils/inputReader";

const [inputMap, inputPath] = asIs("input")
  .split("\n\n")
  .map((s) => s.split("\n"));

const path = inputPath.join("").split("");
const rawMap = inputMap.map((s) => s.split(""));

let robotX = -1;
let robotY = -1;

const wallSet = new Set<string>();
const boxSet = new Set<string>();

function ser(y: number, x: number): string {
  return `${y}|${x}`;
}
function deser(val: string): { x: number; y: number } {
  const [y, x] = val.split("|").map(Number);
  return { y, x };
}

for (let y = 0; y < rawMap.length; y++) {
  for (let x = 0; x < rawMap[0].length; x++) {
    const val = rawMap[y][x];
    if (val === "@") {
      robotX = x;
      robotY = y;
    } else if (val === "O") {
      boxSet.add(ser(y, x));
    } else if (val === "#") {
      wallSet.add(ser(y, x));
    }
  }
}

function getDirection(input: string): { x: number; y: number } {
  switch (input) {
    case "<":
      return { x: -1, y: 0 };
    case ">":
      return { x: 1, y: 0 };
    case "^":
      return { x: 0, y: -1 };
    case "v":
      return { x: 0, y: 1 };
  }
  throw Error("Wrong input");
}

function getCollision(
  dirX: number,
  dirY: number
): { x: number; y: number; type: "wall" | "open" } {
  let currX = robotX;
  let currY = robotY;
  let currVal = "@";
  while (currVal === "@" || currVal === "O") {
    currX += dirX;
    currY += dirY;
    if (boxSet.has(ser(currY, currX))) {
      currVal = "O";
    } else if (wallSet.has(ser(currY, currX))) {
      currVal = "#";
    } else {
      currVal = ".";
    }
  }

  return { x: currX, y: currY, type: currVal === "#" ? "wall" : "open" };
}

for (const step of path) {
  const dir = getDirection(step);
  const coll = getCollision(dir.x, dir.y);

  if (coll.type === "open") {
    let currX = coll.x;
    let currY = coll.y;

    while (currX !== robotX + dir.x || currY !== robotY + dir.y) {
      boxSet.add(ser(currY, currX));
      currX -= dir.x;
      currY -= dir.y;
    }

    boxSet.delete(ser(currY, currX));
    robotX = currX;
    robotY = currY;
  }
}

let coordSum = 0;

for (const box of boxSet) {
  const loc = deser(box);
  coordSum += 100 * loc.y + loc.x;
}

console.log(coordSum);

// Part 2

robotX = -1;
robotY = -1;

wallSet.clear();
boxSet.clear();

for (let y = 0; y < rawMap.length; y++) {
  for (let x = 0; x < rawMap[0].length; x++) {
    const val = rawMap[y][x];
    if (val === "@") {
      robotX = x * 2;
      robotY = y;
    } else if (val === "O") {
      boxSet.add(ser(y, x * 2));
    } else if (val === "#") {
      wallSet.add(ser(y, x * 2));
      wallSet.add(ser(y, x * 2 + 1));
    }
  }
}

function getBroadCollision(
  dirX: number,
  dirY: number
): { boxes: Set<string>; type: "wall" | "open" } {
  const pushingSquares: { x: number; y: number }[] = [{ x: robotX, y: robotY }];
  const pushedBoxes: Set<string> = new Set();

  let wallHit = false;

  while (pushingSquares.length > 0 && !wallHit) {
    const curr = pushingSquares.pop()!;
    const dirTarget = ser(curr.y + dirY, curr.x + dirX);
    const shiftedTarget = ser(curr.y + dirY, curr.x + dirX - 1);

    if (wallSet.has(dirTarget)) {
      wallHit = true;
    } else {
      if (dirY !== 0) {
        if (boxSet.has(dirTarget)) {
          pushedBoxes.add(dirTarget);
          pushingSquares.push({ x: curr.x + dirX, y: curr.y + dirY });
          pushingSquares.push({ x: curr.x + dirX + 1, y: curr.y + dirY });
        }

        if (boxSet.has(shiftedTarget)) {
          pushedBoxes.add(shiftedTarget);
          pushingSquares.push({ x: curr.x + dirX - 1, y: curr.y + dirY });
          pushingSquares.push({ x: curr.x + dirX, y: curr.y + dirY });
        }
      } else if (dirX === -1) {
        if (boxSet.has(shiftedTarget)) {
          pushedBoxes.add(shiftedTarget);
          pushingSquares.push({ x: curr.x + dirX - 1, y: curr.y + dirY });
          pushingSquares.push({ x: curr.x + dirX, y: curr.y + dirY });
        }
      } else {
        if (boxSet.has(dirTarget)) {
          pushedBoxes.add(dirTarget);
          pushingSquares.push({ x: curr.x + dirX + 1, y: curr.y + dirY });
          pushingSquares.push({ x: curr.x + dirX, y: curr.y + dirY });
        }
      }
    }
  }

  if (wallHit) {
    return { boxes: new Set(), type: "wall" };
  }

  return { boxes: pushedBoxes, type: "open" };
}

function vis(): void {
  for (let y = 0; y < rawMap.length; y++) {
    let line = "";
    for (let x = 0; x < rawMap[0].length * 2; x++) {
      if (y === robotY && x === robotX) {
        line += "@";
      } else if (wallSet.has(ser(y, x))) {
        line += "#";
      } else if (boxSet.has(ser(y, x))) {
        line += "[";
      } else if (boxSet.has(ser(y, x - 1))) {
        line += "]";
      } else {
        line += ".";
      }
    }
    console.log(line);
  }
  console.log();
}

for (const step of path) {
  const dir = getDirection(step);
  const coll = getBroadCollision(dir.x, dir.y);

  if (coll.type === "open") {
    for (const box of coll.boxes) {
      boxSet.delete(box);
    }
    for (const box of coll.boxes) {
      const deserBox = deser(box);
      const pushedBox = ser(deserBox.y + dir.y, deserBox.x + dir.x);
      boxSet.add(pushedBox);
    }

    robotX += dir.x;
    robotY += dir.y;
  }

  // goshdarn I needed this
  // console.log(step, coll.type);
  // vis();
}

coordSum = 0;

for (const box of boxSet) {
  const loc = deser(box);
  coordSum += 100 * loc.y + loc.x;
}

console.log(coordSum);
