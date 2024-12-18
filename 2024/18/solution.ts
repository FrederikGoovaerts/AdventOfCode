import { asList } from "../utils/inputReader";
import { MinHeap } from "../utils/minheap";
import { getNeighboursLocs } from "../utils/neighbours";

// const [inputFile, max, fallen] = ["ex1", 6, 12];
const [inputFile, max, fallen] = ["input", 70, 1024];

const input = asList(inputFile);

const fallMap: Map<string, number> = new Map();

input.forEach((val, i) => {
  fallMap.set(val, i + 1);
});

interface Node {
  x: number;
  y: number;
}

function ser(n: Node): string {
  return `${n.x},${n.y}`;
}

function getLength(fallen: number): number | undefined {
  const startNode: Node = { x: 0, y: 0 };
  const endNode: Node = { x: max, y: max };

  const prev = new Map<string, string>();
  const dist = new Map<string, number>();
  const explored = new Set<string>();
  const toExplore = new MinHeap<Node>(
    (a, b) => dist.get(ser(a))! - dist.get(ser(b))!
  );

  dist.set(ser(startNode), 0);
  toExplore.insert(startNode);

  while (toExplore.size() > 0 && !explored.has(ser(endNode))) {
    const curr = toExplore.removeMin()!;
    explored.add(ser(curr));

    const neighbours = getNeighboursLocs(curr.x, curr.y, max, max);

    for (const neighbour of neighbours) {
      if (
        !explored.has(ser(neighbour)) &&
        (fallMap.get(ser(neighbour)) ?? fallen + 1) > fallen
      ) {
        const nDist = dist.get(ser(curr))! + 1;
        if (!toExplore.has(neighbour)) {
          const nDistStored = dist.get(ser(neighbour));
          toExplore.insert(neighbour);
          if (nDistStored === undefined || nDistStored > nDist) {
            prev.set(ser(neighbour), ser(curr));
            dist.set(ser(neighbour), nDist);
          }
        }
      }
    }
  }

  return dist.get(ser(endNode));
}

console.log(getLength(fallen));

let minFall = fallen + 1;
let maxFall = input.length - 1;
let curr = Math.ceil((minFall + maxFall) / 2);

while (minFall !== maxFall) {
  const res = getLength(curr);

  if (res === undefined) {
    maxFall = curr - 1;
  } else {
    minFall = curr;
  }
  curr = Math.ceil((minFall + maxFall) / 2);
}

console.log(input[curr]);
