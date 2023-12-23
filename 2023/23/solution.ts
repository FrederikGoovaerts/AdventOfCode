import { isEqual } from "lodash";
import { as2d } from "../utils/inputReader";
import { getDirectionNeighboursLocs } from "../utils/neighbours";

const input = as2d("input");

interface Coords {
  x: number;
  y: number;
}

function stringify(c: Coords): string {
  return `${c.x}|${c.y}`;
}

const start = { x: 1, y: 0 };
const end = { x: input[0].length - 2, y: input.length - 1 };

function recursiveSearch(
  current: Coords,
  visited: Set<string>,
  globalBest: number,
  icy: boolean
): number {
  if (isEqual(current, end)) {
    // for (let y = 0; y < input.length; y++) {
    //   let line = "";
    //   for (let x = 0; x < input[0].length; x++) {
    //     if (visited.has(stringify({ x, y }))) {
    //       line += "O";
    //     } else {
    //       line += input[y][x];
    //     }
    //   }
    //   console.log(line);
    // }
    const res = Math.max(globalBest, visited.size - 1);
    // console.log(visited.size - 1);
    // console.log();
    return res;
  }

  const currentChar = input[current.y][current.x];

  const neighbours = getDirectionNeighboursLocs(
    current.x,
    current.y,
    input,
    false
  );

  const options: Coords[] = [];

  if (icy && currentChar === "v" && neighbours.down) {
    options.push(neighbours.down);
  } else if (icy && currentChar === "^" && neighbours.up) {
    options.push(neighbours.up);
  } else if (icy && currentChar === ">" && neighbours.right) {
    options.push(neighbours.right);
  } else if (icy && currentChar === "<" && neighbours.left) {
    options.push(neighbours.left);
  } else {
    options.push(
      ...Object.values(neighbours).filter((c) => input[c.y][c.x] !== "#")
    );
  }

  let localBest = globalBest;

  for (const option of options) {
    if (!visited.has(stringify(option))) {
      const optionVisited = new Set(visited);
      optionVisited.add(stringify(option));
      localBest = Math.max(
        localBest,
        recursiveSearch(option, optionVisited, localBest, icy)
      );
    }
  }

  return localBest;
}

const startingVisited = new Set<string>([stringify(start)]);
const result = recursiveSearch(start, startingVisited, 0, true);

console.table(result);

const result2 = recursiveSearch(start, startingVisited, 0, false);

console.table(result2);
