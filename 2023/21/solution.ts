import { asList } from "../utils/inputReader";
import { MinHeap } from "../utils/minheap";
import { getNeighboursLocs } from "../utils/neighbours";

// const [input, maxSteps] = [asList("ex1"), 6];
const [input, maxSteps] = [asList("input"), 64];

const map = input.map((line) => line.split(""));

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
      if (!toExplore.has(neighbour)) {
        toExplore.insert(neighbour);
      }
      const nDist = dist.get(stringify(curr))! + 1;
      const nDistStored = dist.get(stringify(neighbour));
      if (nDistStored === undefined || nDistStored > nDist) {
        prev.set(stringify(neighbour), stringify(curr));
        dist.set(stringify(neighbour), nDist);
      }
    }
  }
}

let stepOptions = 0;

for (const steps of dist.values()) {
  if (steps <= maxSteps && steps % 2 === 0) {
    stepOptions++;
  }
}

console.log(stepOptions);
