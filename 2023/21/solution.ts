import { asList } from "../utils/inputReader";
import { MinHeap } from "../utils/minheap";
import { getNeighboursLocs } from "../utils/neighbours";

// const [input, maxSteps] = [asList("ex1"), 6];
const [input, maxSteps] = [asList("input"), 64];

function expandMap(original: string[], times: number): string[] {
  if (times % 2 === 0) {
    throw new Error("Only equipped to handle odds");
  }

  const expandedHorizontally: string[] = [];
  for (const line of original) {
    let repeated = line.repeat(times);
    while (
      repeated.indexOf("S") !== -1 &&
      repeated.indexOf("S") !== repeated.lastIndexOf("S")
    ) {
      repeated =
        repeated.substring(0, repeated.indexOf("S")) +
        "." +
        repeated.substring(
          repeated.indexOf("S") + 1,
          repeated.lastIndexOf("S")
        ) +
        "." +
        repeated.substring(repeated.lastIndexOf("S") + 1);
    }

    expandedHorizontally.push(repeated);
  }

  const expanded: string[] = [];

  for (let i = 0; i < times; i++) {
    if (i === Math.floor(times / 2)) {
      expanded.push(...expandedHorizontally);
    } else {
      expanded.push(...expandedHorizontally.map((l) => l.replaceAll("S", ".")));
    }
  }

  return expanded;
}

const rawMap = expandMap(input, 3);
const map = rawMap.map((line) => line.split(""));

const startNode = { x: 0, y: 0 };

for (let y = 0; y < map.length; y++) {
  for (let x = 0; x < map[0].length; x++) {
    if (map[y][x] === "S") {
      startNode.x = x;
      startNode.y = y;
    }
  }
}

interface Node {
  x: number;
  y: number;
}

function stringify(node: Node): string {
  return `${node.x}|${node.y}`;
}

const prev = new Map<string, string>();
const dist = new Map<string, number>();
const toExplore = new MinHeap<Node>(
  (a, b) => dist.get(stringify(a))! - dist.get(stringify(b))!
);
const explored = new Set<string>();

dist.set(stringify(startNode), 0);
toExplore.insert(startNode);

while (toExplore.size() > 0) {
  const curr = toExplore.removeMin()!;
  explored.add(stringify(curr));

  const neighbours = getNeighboursLocs(curr.x, curr.y, map, false);
  for (const neighbour of neighbours) {
    if (
      !explored.has(stringify(neighbour)) &&
      map[neighbour.y][neighbour.x] === "."
    ) {
      const nDist = dist.get(stringify(curr))! + 1;
      if (nDist <= maxSteps && !toExplore.has(neighbour)) {
        const nDistStored = dist.get(stringify(neighbour));
        toExplore.insert(neighbour);
        if (nDistStored === undefined || nDistStored > nDist) {
          prev.set(stringify(neighbour), stringify(curr));
          dist.set(stringify(neighbour), nDist);
        }
      }
    }
  }
}

let stepOptions = 0;

for (const steps of dist.values()) {
  if (steps <= maxSteps && steps % 2 === 1) {
    stepOptions++;
  }
}

console.log(stepOptions);

// Part 2 was mostly solved on paper by looking at the structure of the input.
// The values below were generated with the algorithm above

// 65 steps, just the center diamond (an even diamond)
const CENTER = 3742;
// 327 (65 + 2*131) steps, 25 diamonds: 13 even diamonds, 4 inner odd diamonds, 8 outer odd diamonds
const TWO_STEP_RESULT = 93148;
// 589 (65 + 4*131) steps, 81 diamonds: 41 even diamonds, 24 inner odd diamonds, 16 outer odd diamonds
const FOUR_STEP_RESULT = 301602;

// Taking the center as the known value for even diamonds, we have a system of equations with
// outer odd and inner odd diamond sizes as unknowns, and two equations. The result:
const evens = 3742;
const oddsOut = 7427 / 2;
const oddsIn = 7397 / 2;

// const STEP_MULT = 2;
// const STEP_MULT = 4;
// The requested value: 26501365 = 65 + 202300 * 131
const STEP_MULT = 202300;

const diamondSide = STEP_MULT * 2 + 1;
const totalDiamonds = diamondSide * diamondSide;
const totalEvenDiamonds = Math.ceil(totalDiamonds / 2);
const totalOddDiamonds = Math.floor(totalDiamonds / 2);
const oddOutDiamonds = Math.ceil((diamondSide - 2) / 2) * 4;
const oddInDiamonds = totalOddDiamonds - oddOutDiamonds;

console.log(
  totalEvenDiamonds * evens + oddOutDiamonds * oddsOut + oddInDiamonds * oddsIn
);
