import { asIs } from "../utils/inputReader";

const input = asIs("input")
  .split("\n\n")
  .map((s) => s.split("\n").map((l) => l.split("")));

type Combination = [number, number, number, number, number];

const locks: Combination[] = [];
const keys: Combination[] = [];

function getCombination(square: string[][]): Combination {
  const combination: Combination = [0, 0, 0, 0, 0];

  for (let y = 0; y < 7; y++) {
    for (let x = 0; x < 5; x++) {
      if (square[y][x] === "#") {
        combination[x] += 1;
      }
    }
  }

  return combination;
}

for (const square of input) {
  if (square[0][0] === "#") {
    locks.push(getCombination(square));
  } else {
    keys.push(getCombination(square));
  }
}

let count = 0;

for (const lock of locks) {
  for (const key of keys) {
    let fits = true;
    for (let digit = 0; digit < 5; digit++) {
      if (lock[digit] + key[digit] > 7) {
        fits = false;
      }
    }
    if (fits) {
      count++;
    }
  }
}
console.log(count);
