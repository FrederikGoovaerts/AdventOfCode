import { asIs } from "../utils/inputReader";
const input = asIs("input");
const PART_2_MODE = true;
const [instructions, rawMaps] = input.split("\n\n");

interface Map {
  node: string;
  R: string;
  L: string;
}

const maps: Record<string, Map> = {};
const endingWithA: string[] = [];

for (const rawMap of rawMaps.split("\n")) {
  const matches = rawMap.match(/([\w]*) = \(([\w]*), ([\w]*)\)/)!;
  maps[matches[1]] = { node: matches[1], L: matches[2], R: matches[3] };
  if (matches[1].endsWith("A")) {
    endingWithA.push(matches[1]);
  }
}

function getCurrentInstruction(step: number): "L" | "R" {
  return instructions[step % instructions.length] as "L" | "R";
}

function allDiv(n: number, state: number[]): boolean {
  return state.every((s) => n % s === 0);
}

if (!PART_2_MODE) {
  let stepCount = 0;
  let curr = "AAA";

  while (curr !== "ZZZ") {
    const inst = getCurrentInstruction(stepCount);
    curr = maps[curr][inst];
    stepCount++;
  }

  console.log(stepCount);
} else {
  const zs: number[][] = [];
  for (const currStart of endingWithA) {
    let stepCount = 0;
    let curr = currStart;
    let zCount = 0;
    const zSteps: number[] = [];

    while (zCount < 2) {
      const inst = getCurrentInstruction(stepCount);
      curr = maps[curr][inst];
      stepCount++;
      if (curr.endsWith("Z")) {
        zCount++;
        zSteps.push(stepCount);
      }
    }
    zs.push(zSteps);
  }

  // Apparently the intervals between endings are exactly as big every time, even the first

  const z = [...zs.map((z) => z[0])];
  const largest = Math.max(...z);
  let mult = 1;

  while (!allDiv(largest * mult, z)) {
    mult++;
  }
  console.log(largest * mult);
}
