import { asList } from "../utils/inputReader";

const input = asList("input");

const connectionMap: Map<string, Set<string>> = new Map();

function addConnection(c1: string, c2: string) {
  const set = connectionMap.get(c1) ?? new Set();

  set.add(c2);

  connectionMap.set(c1, set);
}

for (const line of input) {
  const [c1, c2] = line.split("-");
  addConnection(c1, c2);
  addConnection(c2, c1);
}

const tConns = new Set<string>();

for (const comp of connectionMap.keys()) {
  if (!comp.startsWith("t")) {
    continue;
  }

  const connections = [...connectionMap.get(comp)!];

  for (let a = 0; a < connections.length; a++) {
    for (let b = a + 1; b < connections.length; b++) {
      const c2 = connections[a];
      const c3 = connections[b];

      if (connectionMap.get(c2)?.has(c3)) {
        tConns.add([comp, c2, c3].toSorted().join());
      }
    }
  }
}

console.log(tConns.size);

const allComps = [...connectionMap.keys()].toSorted();

function pickComputers(chosen: string[], remaining: string[]): string {
  if (remaining.length === 0) {
    return chosen.join();
  }

  const [curr, ...newRemaining] = remaining;

  const currConn = connectionMap.get(curr)!;
  const connectedRemaining = newRemaining.filter((r) => currConn.has(r));

  if (connectedRemaining.length === newRemaining.length) {
    return pickComputers([...chosen, curr], newRemaining);
  } else {
    const resultWith = pickComputers([...chosen, curr], connectedRemaining);
    const resultWithout = pickComputers(chosen, newRemaining);

    if (resultWith.length > resultWithout.length) {
      return resultWith;
    } else {
      return resultWithout;
    }
  }
}

console.log(pickComputers([], allComps));
