import { asIs } from "../utils/inputReader";

const input = asIs("input").split("\n\n");

const rulesAsWritten = input[0].split("\n");
const updates = input[1].split("\n");

// For each page, contains an array of pages the should PRECEDE it if present
const ruleMap: Record<string, string[]> = {};

for (const rule of rulesAsWritten) {
  const [a, b] = rule.split("|");

  ruleMap[b] = [...(ruleMap[b] ?? []), a];
}

const correctUpdates: string[] = [];
const fixedIncorrectUpdates: string[] = [];

for (const update of updates) {
  let correct = true;
  const pages = update.split(",");

  const forbiddenPages: Set<string> = new Set();

  for (let i = 0; i < pages.length; ) {
    const page = pages[i];
    if (forbiddenPages.has(page)) {
      correct = false;
      // hard restart of this one after some fiddling
      pages[i] = pages[i - 1];
      pages[i - 1] = page;
      i = 0;
      forbiddenPages.clear();
    } else {
      for (const forbiddenPage of ruleMap[page] ?? []) {
        forbiddenPages.add(forbiddenPage);
      }
      i++;
    }
  }

  if (correct) {
    correctUpdates.push(update);
  } else {
    fixedIncorrectUpdates.push(pages.join(","));
  }
}

let middleSum = 0;

for (const update of correctUpdates) {
  const parts = update.split(",");
  const middlePart = Number(parts[Math.floor(parts.length / 2)]);
  middleSum += middlePart;
}

console.log(middleSum);

let fixedMiddleSum = 0;

for (const update of fixedIncorrectUpdates) {
  const parts = update.split(",");
  const middlePart = Number(parts[Math.floor(parts.length / 2)]);
  fixedMiddleSum += middlePart;
}

console.log(fixedMiddleSum);
