import { as2dNumbers } from "../utils/inputReader";
import { MinHeap } from "../utils/minheap";
import {
  Direction,
  getDirectionNeighboursLocs,
  reverseDirection,
} from "../utils/neighbours";

const input = as2dNumbers("input");
const goal = { x: input[0].length - 1, y: input.length - 1 };

interface Node {
  x: number;
  y: number;
  loss: number;
  ld: Direction;
  ldstep: number;
}

const start: Node = { x: 0, y: 0, loss: 0, ld: "right", ldstep: 0 };

function stringify(coords: {
  x: number;
  y: number;
  ld: Direction;
  ldstep: number;
}): string {
  return `${coords.x}|${coords.y}|${coords.ld}|${coords.ldstep}`;
}

const currentBest = new Map<string, number>();
currentBest.set(stringify(start), 0);

const nodes = new MinHeap<Node>((a, b) => a.loss - b.loss);
nodes.insert(start);

while (nodes.size() > 0) {
  const curr = nodes.removeMin()!;
  if (curr.x === goal.x && curr.y === goal.y) {
    console.log(curr);
    break;
  }
  const ns = getDirectionNeighboursLocs(curr.x, curr.y, input, false);
  for (const [direction, coords] of Object.entries(ns)) {
    if (
      (curr.ld === direction && curr.ldstep === 3) ||
      curr.ld === reverseDirection(direction as Direction)
    ) {
      continue;
    } else {
      const neighbourLoss = input[coords.y][coords.x];
      const next = {
        x: coords.x,
        y: coords.y,
        loss: curr.loss + neighbourLoss,
        ld: direction as Direction,
        ldstep: direction === curr.ld ? curr.ldstep + 1 : 1,
      };

      const bestForNeighbour = currentBest.get(stringify(next));

      if (
        bestForNeighbour === undefined ||
        curr.loss + neighbourLoss < bestForNeighbour
      ) {
        nodes.insert(next);
        if (curr.ld !== direction || curr.ldstep < 3) {
          currentBest.set(stringify(next), curr.loss + neighbourLoss);
        }
      }
    }
  }
}
