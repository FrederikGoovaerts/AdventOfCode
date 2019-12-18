import * as fs from "fs";

const input: string[][] = fs
  .readFileSync("in4", "utf8")
  .trim()
  .split("\n")
  .map(val => val.split(""));

interface Node {
  pos: string;
  length: number;
  keys: Set<string>;
}

class Queue<T> {
  private a: T[] = [];
  private b = 0;
  getLength() {
    return this.a.length - this.b;
  }
  isEmpty() {
    return 0 == this.a.length;
  }
  enqueue(b: T) {
    this.a.push(b);
  }
  dequeue() {
    if (0 != this.a.length) {
      var c = this.a[this.b];
      2 * ++this.b >= this.a.length &&
        ((this.a = this.a.slice(this.b)), (this.b = 0));
      return c;
    }
  }
  peek() {
    return 0 < this.a.length ? this.a[this.b] : void 0;
  }
}

const neighbors: Map<string, string[]> = new Map();
// Map from position to door name
const doorLoc: Map<string, string> = new Map();
// Map from position to key name
const keyLocations: Map<string, string> = new Map();
let origin: [number, number] = [-1, -1];

for (let row = 1; row < input.length - 1; row++) {
  for (let column = 1; column < input[0].length - 1; column++) {
    const symbol = input[row][column];
    if (symbol === "#") {
      continue;
    } else {
      const possibleNeighbors: [number, number][] = [
        [row + 1, column],
        [row - 1, column],
        [row, column + 1],
        [row, column - 1]
      ];
      const locNeighbors = [];
      for (const n of possibleNeighbors) {
        if (input[n[0]][n[1]] !== "#") {
          locNeighbors.push(serPos(n));
        }
      }
      neighbors.set(serPos([row, column]), locNeighbors);
    }
    if (symbol.match(/[a-z]/)) {
      keyLocations.set(serPos([row, column]), symbol);
    } else if (symbol.match(/[A-Z]/)) {
      doorLoc.set(serPos([row, column]), symbol);
    } else if (symbol === "@") {
      originPos = [row, column];
    }
  }
}

const visited: Set<string> = new Set();

const dijkstraQueue: Queue<Node> = new Queue();
dijkstraQueue.enqueue({ pos: serPos(origin), length: 0, keys: new Set() });
while (true) {
  const curr = dijkstraQueue.dequeue()!;
  const nextList = neighbors.get(curr.pos)!;
  for (const next of nextList) {
    if (doorLoc.has(next) && !curr.keys.has(doorLoc.get(next)!.toLowerCase())) {
      continue;
    }
    let nextNode: Node = curr;
    if (keyLocations.has(next)) {
      const newKeys = new Set(curr.keys);
      newKeys.add(keyLocations.get(next)!);
      if (newKeys.size === keyLocations.size) {
        throw new Error(`${curr.length + 1}`);
      }
      nextNode = { pos: next, length: curr.length + 1, keys: newKeys };
    } else {
      nextNode = {
        pos: next,
        length: curr.length + 1,
        keys: curr.keys
      };
    }
    if (!visited.has(serNode(nextNode))) {
      dijkstraQueue.enqueue(nextNode);
    }
  }
  visited.add(serNode(curr));
}

function serPos(input: [number, number]): string {
  return `${input[0]},${input[1]}`;
}
function serNode(input: Node): string {
  return `${input.pos}${[...input.keys].join()}`;
}
