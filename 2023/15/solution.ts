import { sum } from "lodash";
import { asLineList } from "../utils/inputReader";

const input = asLineList("input");

function hash(text: string): number {
  let curr = 0;
  for (const char of text.split("")) {
    const code = char.charCodeAt(0);
    curr += code;
    curr = (curr * 17) % 256;
  }
  return curr;
}

const results: number[] = [];

for (const line of input) {
  results.push(hash(line));
}

console.log(sum(results));

interface Lens {
  label: string;
  str: number;
}

const boxes: Record<number, Lens[]> = {};

for (let i = 0; i < 256; i++) {
  boxes[i] = [];
}

for (const line of input) {
  if (line.endsWith("-")) {
    const label = line.slice(0, line.length - 1);
    const h = hash(label);
    const box = boxes[h];
    const updatedBox = box.filter((v) => v.label !== label);
    boxes[h] = updatedBox;
  } else {
    const label = line.slice(0, line.length - 2);
    const str = parseInt(line[line.length - 1]);
    const h = hash(label);
    const matchingLens = boxes[h].find((v) => v.label === label);
    if (matchingLens) {
      matchingLens.str = str;
    } else {
      boxes[h].push({ label, str });
    }
  }
}

const powers: number[] = [];

for (let i = 0; i < 256; i++) {
  const box = boxes[i];
  box.forEach((l, li) => {
    powers.push((i + 1) * (li + 1) * l.str);
  });
}
console.log(sum(powers));
