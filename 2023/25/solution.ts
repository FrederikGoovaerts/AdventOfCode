import { asList } from "../utils/inputReader";
import { MinHeap } from "../utils/minheap";

const input = asList("input");

function combine(a: string, b: string): string {
  if (a < b) {
    return a + b;
  }
  return b + a;
}

const components = new Set<string>();
const connectsTo = new Map<string, string[]>();
const singleConnections: [string, string][] = [];

function addConnection(orig: string, dest: string): void {
  const list = connectsTo.get(orig) ?? [];
  if (!list.includes(dest)) {
    list.push(dest);
  }
  connectsTo.set(orig, list);
}

for (const line of input) {
  const [origin, rawDestinations] = line.split(": ");
  const destinations = rawDestinations.split(" ");
  components.add(origin);
  for (const dest of destinations) {
    components.add(dest);
    addConnection(origin, dest);
    addConnection(dest, origin);

    const singleConn: [string, string] = [origin, dest];
    singleConn.sort();

    singleConnections.push(singleConn);
  }
}

function shortestPath(
  source: string,
  target: string,
  ignore: Set<string>
): string[] | undefined {
  const prev: Record<string, string> = {};
  const dist: Record<string, number> = {};

  const toExplore = new MinHeap<string>((a, b) => dist[a] - dist[b]);
  const addedForExploration = new Set<string>();
  const explored = new Set<string>();

  dist[source] = 0;
  toExplore.insert(source);

  while (toExplore.size() > 0) {
    const curr = toExplore.removeMin()!;
    if (curr === target) {
      break;
    }
    explored.add(curr);

    const connections = connectsTo.get(curr)!;
    for (const next of connections) {
      if (!explored.has(next) && !ignore.has(combine(curr, next))) {
        const nextDist = dist[curr] + 1;
        if (dist[next] === undefined || dist[next] > nextDist) {
          dist[next] = nextDist;
          prev[next] = curr;
        }
        if (!addedForExploration.has(next)) {
          toExplore.insert(next);
          addedForExploration.add(next);
        }
      }
    }
  }
  if (dist[target] === undefined) {
    return undefined;
  }
  const path: string[] = [target];
  let curr = target;
  while (curr !== source) {
    path.push(prev[curr]);
    curr = prev[curr];
  }
  path.reverse();
  return path;
}

function getConnections(): [string, string][] {
  for (const base of singleConnections) {
    const firstIgnore = combine(base[0], base[1]);
    const firstPath = shortestPath(base[0], base[1], new Set([firstIgnore]))!;

    for (let i = 0; i < firstPath.length - 1; i++) {
      const secondIgnore = combine(firstPath[i], firstPath[i + 1]);
      const secondPath = shortestPath(
        base[0],
        base[1],
        new Set([firstIgnore, secondIgnore])
      )!;
      for (let j = 0; j < firstPath.length - 1; j++) {
        const thirdIgnore = combine(secondPath[j], secondPath[j + 1]);
        const thirdPath = shortestPath(
          base[0],
          base[1],
          new Set([firstIgnore, secondIgnore, thirdIgnore])
        );

        if (!thirdPath) {
          return [
            [base[0], base[1]],
            [firstPath[i], firstPath[i + 1]],
            [secondPath[j], secondPath[j + 1]],
          ];
        }
      }
    }
  }
  throw new Error("Not found!");
}

function randomElementReaches(toIgnore: [string, string][]): number {
  const startElement = input[0].split(": ")[0];
  const visited = new Set([startElement]);
  const toExplore = [startElement];
  while (toExplore.length > 0) {
    const curr = toExplore.pop()!;
    for (const next of connectsTo.get(curr)!) {
      if (
        !toIgnore.some((i) => i.includes(next) && i.includes(curr)) &&
        !visited.has(next)
      ) {
        visited.add(next);
        toExplore.push(next);
      }
    }
  }
  return visited.size;
}

const toIgnore = getConnections();
const reach = randomElementReaches(toIgnore);
console.log(reach * (components.size - reach));
