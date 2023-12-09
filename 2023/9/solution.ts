import { sum } from "lodash";
import { asList } from "../utils/inputReader";
const input = asList("input");

const nexts: number[] = [];
const prevs: number[] = [];

function extrapolate(list: number[], mode: "NEXT" | "PREV"): number {
  if (list.every((v) => v === list[0])) {
    return list[0];
  }
  const subList: number[] = [];

  for (let i = 0; i < list.length - 1; i++) {
    subList.push(list[i + 1] - list[i]);
  }

  const subEx = extrapolate(subList, mode);

  return mode === "NEXT" ? list.at(-1)! + subEx : list[0] - subEx;
}

for (const line of input) {
  const parsedLine = line.split(" ").map((v) => parseInt(v));
  nexts.push(extrapolate(parsedLine, "NEXT"));
  prevs.push(extrapolate(parsedLine, "PREV"));
}

console.log(sum(nexts));
console.log(sum(prevs));
