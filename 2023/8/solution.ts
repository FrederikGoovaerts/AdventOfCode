import { asIs } from "../utils/inputReader";
import { multipleLcm } from "../utils/math";
const input = asIs("input");
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

let stepCount = 0;
let curr = "AAA";

while (curr !== "ZZZ") {
  const inst = getCurrentInstruction(stepCount);
  curr = maps[curr][inst];
  stepCount++;
}

console.log(stepCount);

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
const zDistances = [...zs.map((z) => z[0])];

console.log(multipleLcm(zDistances));
