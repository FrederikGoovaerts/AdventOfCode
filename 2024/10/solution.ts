import { as2dNumbers } from "../utils/inputReader";
import { getNeighboursLocs } from "../utils/neighbours";

const input = as2dNumbers("input");

// This one checks all possible paths, might come in handy later
function getHeadScore(y: number, x: number, val: number): number {
  if (val === 9) {
    return 1;
  }

  const ns = getNeighboursLocs(x, y, input, false);

  let score = 0;

  for (const n of ns) {
    if (input[n.y][n.x] === val + 1) {
      score += getHeadScore(n.y, n.x, val + 1);
    }
  }
  return score;
}

function checkHeads(
  y: number,
  x: number,
  val: number,
  result: Set<string>
): void {
  if (val === 9) {
    result.add(`${y},${x}`);
    return;
  }

  const ns = getNeighboursLocs(x, y, input, false);

  for (const n of ns) {
    if (input[n.y][n.x] === val + 1) {
      checkHeads(n.y, n.x, val + 1, result);
    }
  }
}

let total = 0;

for (let y = 0; y < input.length; y++) {
  for (let x = 0; x < input[0].length; x++) {
    if (input[y][x] === 0) {
      const results = new Set<string>();
      checkHeads(y, x, 0, results);

      total += results.size;
    }
  }
}

console.log(total);

let totalScore = 0;

for (let y = 0; y < input.length; y++) {
  for (let x = 0; x < input[0].length; x++) {
    if (input[y][x] === 0) {
      totalScore += getHeadScore(y, x, 0);
    }
  }
}

console.log(totalScore);
