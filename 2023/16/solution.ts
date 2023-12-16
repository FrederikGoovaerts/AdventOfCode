import { as2d } from "../utils/inputReader";

const input = as2d("input");

type Direction = "u" | "d" | "l" | "r";

interface BeamEl {
  x: number;
  y: number;
  dir: Direction;
}

function stringify(el: BeamEl): string {
  return `${el.x}|${el.y}|${el.dir}`;
}
function stringifyNoDir(el: BeamEl): string {
  return `${el.x}|${el.y}`;
}

function checkEnergy(start: BeamEl) {
  const visited = new Set<string>();
  const energized = new Set<string>();
  const toExpand: BeamEl[] = [];

  toExpand.push(start);

  while (toExpand.length > 0) {
    const curr = toExpand.pop()!;
    if (visited.has(stringify(curr))) {
      continue;
    }
    if (
      curr.x < 0 ||
      curr.x >= input[0].length ||
      curr.y < 0 ||
      curr.y >= input.length
    ) {
      continue;
    }
    energized.add(stringifyNoDir(curr));
    visited.add(stringify(curr));
    const char = input[curr.y][curr.x];

    if (char === ".") {
      switch (curr.dir) {
        case "u":
          toExpand.push({ ...curr, y: curr.y - 1 });
          break;
        case "d":
          toExpand.push({ ...curr, y: curr.y + 1 });
          break;
        case "l":
          toExpand.push({ ...curr, x: curr.x - 1 });
          break;
        case "r":
          toExpand.push({ ...curr, x: curr.x + 1 });
          break;
      }
    } else if (char === "/") {
      switch (curr.dir) {
        case "u":
          toExpand.push({ ...curr, x: curr.x + 1, dir: "r" });
          break;
        case "d":
          toExpand.push({ ...curr, x: curr.x - 1, dir: "l" });
          break;
        case "l":
          toExpand.push({ ...curr, y: curr.y + 1, dir: "d" });
          break;
        case "r":
          toExpand.push({ ...curr, y: curr.y - 1, dir: "u" });
          break;
      }
    } else if (char === "\\") {
      switch (curr.dir) {
        case "u":
          toExpand.push({ ...curr, x: curr.x - 1, dir: "l" });
          break;
        case "d":
          toExpand.push({ ...curr, x: curr.x + 1, dir: "r" });
          break;
        case "l":
          toExpand.push({ ...curr, y: curr.y - 1, dir: "u" });
          break;
        case "r":
          toExpand.push({ ...curr, y: curr.y + 1, dir: "d" });
          break;
      }
    } else if (char === "-") {
      switch (curr.dir) {
        case "u":
        case "d":
          toExpand.push({ ...curr, x: curr.x - 1, dir: "l" });
          toExpand.push({ ...curr, x: curr.x + 1, dir: "r" });
          break;
        case "l":
          toExpand.push({ ...curr, x: curr.x - 1 });
          break;
        case "r":
          toExpand.push({ ...curr, x: curr.x + 1 });
          break;
      }
    } else if (char === "|") {
      switch (curr.dir) {
        case "u":
          toExpand.push({ ...curr, y: curr.y - 1 });
          break;
        case "d":
          toExpand.push({ ...curr, y: curr.y + 1 });
          break;
        case "l":
        case "r":
          toExpand.push({ ...curr, y: curr.y - 1, dir: "u" });
          toExpand.push({ ...curr, y: curr.y + 1, dir: "d" });
          break;
      }
    }
  }
  return energized.size;
}

console.log(checkEnergy({ x: 0, y: 0, dir: "r" }));

let max = 0;

for (let y = 0; y < input.length; y++) {
  max = Math.max(max, checkEnergy({ x: 0, y, dir: "r" }));
  max = Math.max(max, checkEnergy({ x: input[0].length - 1, y, dir: "l" }));
}
for (let x = 0; x < input.length; x++) {
  max = Math.max(max, checkEnergy({ x, y: 0, dir: "d" }));
  max = Math.max(max, checkEnergy({ x, y: input.length - 1, dir: "u" }));
}

console.log(max);
