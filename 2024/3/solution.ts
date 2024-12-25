import { asList } from "../utils/inputReader";

const linedInput = asList("input");
const input = linedInput.join("");

const regex = /mul\(\d+,\d+\)/g;
const matches = [...input.matchAll(regex)];

let total = 0;

for (const match of matches) {
  const [a, b] = match[0].split("(")[1].split(")")[0].split(",");

  total += Number(a) * Number(b);
}

console.log(total);

const regexDoDont = /do\(\)|don\'t\(\)/g;
const matchesDoDont = [...input.matchAll(regexDoDont)];

function shouldDo(index: number): boolean {
  const beforeOp = matchesDoDont.findLastIndex((m) => m.index <= index);
  if (beforeOp === -1) {
    return true;
  }
  return matchesDoDont[beforeOp][0] === "do()";
}

total = 0;

for (const match of matches) {
  if (!shouldDo(match.index)) {
    continue;
  }

  const [a, b] = match[0].split("(")[1].split(")")[0].split(",");

  total += Number(a) * Number(b);
}

console.log(total);
