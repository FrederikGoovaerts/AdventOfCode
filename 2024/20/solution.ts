import { as2d } from "../utils/inputReader";
import { MinHeap } from "../utils/minheap";
import { getNeighboursLocs } from "../utils/neighbours";

const input = as2d("input");

let sX = 0;
let sY = 0;
let eX = 0;
let eY = 0;

for (let y = 0; y < input.length; y++) {
  for (let x = 0; x < input[0].length; x++) {
    if (input[y][x] === "S") {
      sX = x;
      sY = y;
    }
    if (input[y][x] === "E") {
      eX = x;
      eY = y;
    }
  }
}

interface Node {
  x: number;
  y: number;
}

function ser(node: Node): string {
  return `${node.x}|${node.y}`;
}

function deser(val: string): Node {
  const [x, y] = val.split("|");
  return {
    x: Number(x),
    y: Number(y),
  };
}

const prev = new Map<string, string>();
const dist = new Map<string, number>();
const explored = new Set<string>();
const toExplore = new MinHeap<Node>(
  (a, b) => dist.get(ser(a))! - dist.get(ser(b))!
);

dist.set(ser({ x: sX, y: sY }), 0);
toExplore.insert({ x: sX, y: sY });

const end = ser({ x: eX, y: eY });

while (toExplore.size() > 0 && !explored.has(end)) {
  const curr = toExplore.removeMin()!;
  explored.add(ser(curr));

  const neighbours = getNeighboursLocs(
    curr.x,
    curr.y,
    input[0].length - 1,
    input.length - 1
  ).filter((n) => input[n.y][n.x] !== "#");

  for (const neighbour of neighbours) {
    if (!explored.has(ser(neighbour))) {
      const nDist = dist.get(ser(curr))! + 1;
      const nDistStored = dist.get(ser(neighbour));

      if (nDistStored === undefined || nDistStored > nDist) {
        prev.set(ser(neighbour), ser(curr));
        dist.set(ser(neighbour), nDist);
      }

      if (!toExplore.has(neighbour)) {
        toExplore.insert(neighbour);
      }
    }
  }
}

const path: Node[] = [];
let curr: string | undefined = end;
path.push(deser(end));

while (curr !== undefined) {
  curr = prev.get(curr);
  if (curr) {
    path.unshift(deser(curr));
  }
}

function getJumps(x: number, y: number): { x: number; y: number }[] {
  const result: { x: number; y: number }[] = [
    { x: x, y: y + 2 },
    { x: x + 2, y: y },
    { x: x - 2, y: y },
    { x: x, y: y - 2 },
  ];

  return result.filter(
    (val) =>
      val.x >= 0 &&
      val.x < input[0].length &&
      val.y >= 0 &&
      val.y < input.length
  );
}

let count = 0;

for (const p of path) {
  const pDist = dist.get(ser(p))!;
  const jumps = getJumps(p.x, p.y);

  for (const j of jumps) {
    const jDist = dist.get(ser(j));
    if (jDist) {
      if (jDist - pDist - 2 >= 100) {
        count++;
      }
    }
  }
}

console.log(count);

function getLongJumps(
  x: number,
  y: number
): { x: number; y: number; jumpSize: number }[] {
  const result: { x: number; y: number; jumpSize: number }[] = [];
  for (let offY = -20; offY <= 20; offY++) {
    for (let offX = -20 + Math.abs(offY); offX <= 20 - Math.abs(offY); offX++) {
      const newX = x + offX;
      const newY = y + offY;
      if (
        newX >= 0 &&
        newX < input[0].length &&
        newY >= 0 &&
        newY < input.length
      ) {
        result.push({
          x: newX,
          y: newY,
          jumpSize: Math.abs(offX) + Math.abs(offY),
        });
      }
    }
  }

  return result;
}

count = 0;

const distMap = new Map<number, number>();

for (const p of path) {
  const pDist = dist.get(ser(p))!;
  const jumps = getLongJumps(p.x, p.y);

  for (const j of jumps) {
    const jDist = dist.get(ser({ ...j }));
    if (jDist) {
      const saved = jDist - pDist - j.jumpSize;
      if (saved >= 100) {
        count++;
        distMap.set(saved, (distMap.get(saved) ?? 0) + 1);
      }
    }
  }
}

console.log(count);
