import { as2d } from "../utils/inputReader";
import { MinHeap } from "../utils/minheap";
import {
  Direction,
  getDirectionStep,
  turnClockwise,
  turnCounterclockwise,
} from "../utils/neighbours";

const input = as2d("input");

interface Node {
  x: number;
  y: number;
  dir: Direction;
}

function ser(node: Node): string {
  return `${node.x}|${node.y}|${node.dir}`;
}
function deser(val: string): Node {
  const [x, y, dir] = val.split("|");
  return {
    x: Number(x),
    y: Number(y),
    dir: dir as Direction,
  };
}

const startNode: Node = { x: 1, y: input.length - 2, dir: "right" };

const endNode1: Node = { x: input[0].length - 2, y: 1, dir: "right" };
const endNode2: Node = { x: input[0].length - 2, y: 1, dir: "up" };

const prev = new Map<string, string[]>();
const dist = new Map<string, number>();
const explored = new Set<string>();
const toExplore = new MinHeap<Node>(
  (a, b) => dist.get(ser(a))! - dist.get(ser(b))!
);

dist.set(ser(startNode), 0);
toExplore.insert(startNode);

function getSteps(
  node: Node
): { x: number; y: number; dir: Direction; stepCost: number }[] {
  const result: { x: number; y: number; dir: Direction; stepCost: number }[] =
    [];

  const goDirStep = getDirectionStep(node.dir);
  const goNode: Node = {
    x: node.x + goDirStep.x,
    y: node.y + goDirStep.y,
    dir: node.dir,
  };
  const cwDirStep = getDirectionStep(turnClockwise(node.dir));
  const cwNode: Node = {
    x: node.x + cwDirStep.x,
    y: node.y + cwDirStep.y,
    dir: turnClockwise(node.dir),
  };
  const ccwDirStep = getDirectionStep(turnCounterclockwise(node.dir));
  const ccwNode: Node = {
    x: node.x + ccwDirStep.x,
    y: node.y + ccwDirStep.y,
    dir: turnCounterclockwise(node.dir),
  };

  if (input[goNode.y][goNode.x] !== "#") {
    result.push({ ...goNode, stepCost: 1 });
  }
  if (input[cwNode.y][cwNode.x] !== "#") {
    result.push({ ...node, dir: turnClockwise(node.dir), stepCost: 1000 });
  }
  if (input[ccwNode.y][ccwNode.x] !== "#") {
    result.push({
      ...node,
      dir: turnCounterclockwise(node.dir),
      stepCost: 1000,
    });
  }

  return result;
}

while (
  toExplore.size() > 0 &&
  !(explored.has(ser(endNode1)) || explored.has(ser(endNode2)))
) {
  const curr = toExplore.removeMin()!;
  explored.add(ser(curr));

  const neighbours = getSteps(curr);

  for (const neighbour of neighbours) {
    const nNode = { x: neighbour.x, y: neighbour.y, dir: neighbour.dir };

    if (!explored.has(ser(nNode))) {
      const nDist = dist.get(ser(curr))! + neighbour.stepCost;
      const nDistStored = dist.get(ser(nNode));

      if (nDistStored === undefined) {
        prev.set(ser(nNode), [ser(curr)]);
        dist.set(ser(nNode), nDist);
      } else if (nDistStored > nDist) {
        prev.set(ser(nNode), [ser(curr)]);
        dist.set(ser(nNode), nDist);
      } else if (nDistStored === nDist) {
        prev.set(ser(nNode), [...prev.get(ser(nNode))!, ser(curr)]);
        dist.set(ser(nNode), nDist);
      }

      if (!toExplore.has(nNode)) {
        toExplore.insert(nNode);
      }
    }
  }
}

console.log(dist.get(ser(endNode1)) ?? dist.get(ser(endNode2)));

const endNode = dist.has(ser(endNode1)) ? endNode1 : endNode2;

const bestExplored = new Set<string>();
const bestToEplore = new Set<string>();
bestToEplore.add(ser(endNode));

while (bestToEplore.size > 0) {
  const res = bestToEplore.values().next().value!;
  const deserRes = deser(res);
  bestToEplore.delete(res);
  bestExplored.add(`${deserRes.y}|${deserRes.x}`);

  const bestPrevs = prev.get(res);

  for (const p of bestPrevs ?? []) {
    bestToEplore.add(p);
  }
}

console.log(bestExplored.size);
