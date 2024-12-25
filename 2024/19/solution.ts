import { asIs } from "../utils/inputReader";

const [inputTowels, inputPatterns] = asIs("input").split("\n\n");

const towels: Set<string> = new Set();
const truePatterns: Map<string, number> = new Map();
const falsePatterns: Set<string> = new Set();

for (const towel of inputTowels.split(", ")) {
  towels.add(towel);
}

const maxTowelLength = Math.max(...[...towels.values()].map((t) => t.length));

function checkPattern(pattern: string): number {
  if (pattern === "") {
    return 1;
  }

  if (falsePatterns.has(pattern)) {
    return 0;
  }

  let ways = 0;

  for (let i = 1; i <= maxTowelLength && i <= pattern.length; i++) {
    const part = pattern.substring(0, i);

    if (towels.has(part)) {
      if (i === pattern.length) {
        ways++;
        continue;
      }

      const remainder = pattern.substring(i);

      if (truePatterns.has(remainder)) {
        ways += truePatterns.get(remainder)!;
        continue;
      }

      if (falsePatterns.has(remainder)) {
        continue;
      }

      const remainderResult = checkPattern(remainder);

      if (remainderResult !== 0) {
        ways += remainderResult;
        truePatterns.set(remainder, remainderResult);
      } else {
        falsePatterns.add(remainder);
      }
    }
  }

  return ways;
}

const patterns = inputPatterns.split("\n");

let count = 0;
let arrangements = 0;

patterns.forEach((pattern, i) => {
  const result = checkPattern(pattern);
  if (result) {
    count++;
    arrangements += result;
  }
});

console.log(count);
console.log(arrangements);
