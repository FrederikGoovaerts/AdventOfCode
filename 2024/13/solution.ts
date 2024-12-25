import { asIs } from "../utils/inputReader";

const input = asIs("input").split("\n\n").map(toMachine);

interface Machine {
  aX: number;
  aY: number;
  bX: number;
  bY: number;
  pX: number;
  pY: number;
}

function toMachine(input: string): Machine {
  const [aPart, bPart, prizePart] = input.split("\n");
  const [aX, aY] = aPart.split("X+")[1].split(", Y+").map(Number);
  const [bX, bY] = bPart.split("X+")[1].split(", Y+").map(Number);
  const [pX, pY] = prizePart.split("X=")[1].split(", Y=").map(Number);

  return {
    aX,
    aY,
    bX,
    bY,
    pX,
    pY,
  };
}

function getLowestCostFast({ aX, aY, bX, bY, pX, pY }: Machine): number {
  const aPress = (pX * bY - pY * bX) / (bY * aX - aY * bX);
  const bPress = (pX * aY - pY * aX) / (aY * bX - bY * aX);

  if (aPress % 1 === 0 && bPress % 1 === 0) {
    return aPress * 3 + bPress;
  }

  return 0;
}

let cost = 0;
let cost2 = 0;

for (const machine of input) {
  cost += getLowestCostFast(machine);
  cost2 += getLowestCostFast({
    ...machine,
    pX: machine.pX + 10000000000000,
    pY: machine.pY + 10000000000000,
  });
}

console.log(cost);
console.log(cost2);
