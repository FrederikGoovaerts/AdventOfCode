import { isEqual } from "lodash";
import { as2d } from "../utils/inputReader";
import { Direction, getDirectionNeighboursLocs } from "../utils/neighbours";

const input = as2d("input");

function charFor(c: Coords): "#" | "." | ">" | "v" | undefined {
  return input[c.y]?.[c.x] as "#" | "." | ">" | "v" | undefined;
}

interface Coords {
  x: number;
  y: number;
}

interface PathTo {
  dest: Coords;
  length: number;
  icyPassable: boolean;
}

function stringify(c: Coords): string {
  return `${c.x}|${c.y}`;
}
function destringify(c: string): Coords {
  const [x, y] = c.split("|").map((v) => parseInt(v));
  return { x, y };
}

const start = { x: 1, y: 0 };
const end = { x: input[0].length - 2, y: input.length - 1 };

function recursiveMakePath(
  current: Coords,
  last: Coords,
  visited: Set<string>,
  startImpassable: boolean
): PathTo {
  const neighbours = getDirectionNeighboursLocs(
    current.x,
    current.y,
    input,
    false
  );

  const pathNeighbours = Object.values(neighbours).filter(
    (c) => input[c.y][c.x] !== "#"
  );

  if (pathNeighbours.length <= 1) {
    if (isEqual(current, start) || isEqual(current, end)) {
      return {
        dest: current,
        length: visited.size,
        icyPassable: !startImpassable,
      };
    }
    throw new Error(`illegal spot!: ${stringify(current)}`);
  }

  if (pathNeighbours.length > 2) {
    const lastChar = charFor(last);
    if (
      (lastChar === ">" && isEqual(neighbours.right, current)) ||
      (lastChar === "v" && isEqual(neighbours.down, current))
    ) {
      return { dest: current, length: visited.size, icyPassable: false };
    }
    return {
      dest: current,
      length: visited.size,
      icyPassable: !startImpassable,
    };
  }

  for (const option of pathNeighbours) {
    if (!isEqual(option, last)) {
      const optionVisited = new Set(visited);
      optionVisited.add(stringify(option));
      return recursiveMakePath(option, current, optionVisited, startImpassable);
    }
  }

  throw new Error("Could not find next step in path");
}

function getPaths(
  c: Coords,
  pathNeighbours: Partial<Record<Direction, { x: number; y: number }>>
): PathTo[] {
  const result: PathTo[] = [];

  if (pathNeighbours.down) {
    result.push(
      recursiveMakePath(pathNeighbours.down, c, new Set([stringify(c)]), false)
    );
  }
  if (pathNeighbours.up) {
    result.push(
      recursiveMakePath(
        pathNeighbours.up,
        c,
        new Set([stringify(c)]),
        charFor(pathNeighbours.up) === "v"
      )
    );
  }
  if (pathNeighbours.right) {
    result.push(
      recursiveMakePath(pathNeighbours.right, c, new Set([stringify(c)]), false)
    );
  }
  if (pathNeighbours.left) {
    result.push(
      recursiveMakePath(
        pathNeighbours.left,
        c,
        new Set([stringify(c)]),
        charFor(pathNeighbours.left) === ">"
      )
    );
  }

  return result;
}

const pathMap = new Map<string, PathTo[]>();

for (let y = 0; y < input.length; y++) {
  for (let x = 0; x < input[0].length; x++) {
    if (charFor({ x, y }) === ".") {
      const neighbours = getDirectionNeighboursLocs(x, y, input, false);
      const pathNeighbours: Partial<
        Record<Direction, { x: number; y: number }>
      > = {};

      let nCount = 0;
      for (const [dir, n] of Object.entries(neighbours)) {
        if (charFor(n) !== "#") {
          pathNeighbours[dir as Direction] = n;
          nCount++;
        }
      }

      if (nCount > 2) {
        pathMap.set(stringify({ x, y }), getPaths({ x, y }, pathNeighbours));
      }
    }
  }
}

let startCrossing: Coords = start;
let startDist: number = 0;

for (const [k, v] of pathMap.entries()) {
  const startPath = v.find((p) => isEqual(p.dest, start));
  if (startPath) {
    startCrossing = destringify(k);
    startDist = startPath.length;
  }
}

function recursiveSearch(
  current: Coords,
  visitedCrossings: Set<string>,
  dist: number,
  globalBest: number,
  icy: boolean
): number {
  if (isEqual(current, end)) {
    const res = Math.max(globalBest, dist);
    return res;
  }

  const paths = pathMap.get(stringify(current))!;

  const options: PathTo[] = paths.filter(
    (p) => (!icy || p.icyPassable) && !visitedCrossings.has(stringify(p.dest))
  );

  let localBest = globalBest;

  for (const option of options) {
    const optionVisited = new Set(visitedCrossings);
    optionVisited.add(stringify(option.dest));
    localBest = Math.max(
      localBest,
      recursiveSearch(
        option.dest,
        optionVisited,
        dist + option.length,
        localBest,
        icy
      )
    );
  }

  return localBest;
}

const startingVisited = new Set<string>([
  stringify(start),
  stringify(startCrossing),
]);

const result = recursiveSearch(
  startCrossing,
  startingVisited,
  startDist,
  0,
  true
);

console.log(result);

const result2 = recursiveSearch(
  startCrossing,
  startingVisited,
  startDist,
  0,
  false
);

// TODO: figure out how to make this end quickly
console.log(result2);
