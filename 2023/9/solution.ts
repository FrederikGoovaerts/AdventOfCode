import { sum } from "lodash";
import { asList } from "../utils/inputReader";
const input = asList("input");

const nexts: number[] = [];
const prevs: number[] = [];

function getNext(list: number[]): number {
  if (list.every((v) => v === list[0])) {
    return list[0];
  }
  const subList: number[] = [];

  for (let i = 0; i < list.length - 1; i++) {
    subList.push(list[i + 1] - list[i]);
  }

  const subNext = getNext(subList);

  return list.at(-1)! + subNext;
}

function getPrev(list: number[]): number {
  if (list.every((v) => v === list[0])) {
    return list[0];
  }
  const subList: number[] = [];

  for (let i = 0; i < list.length - 1; i++) {
    subList.push(list[i + 1] - list[i]);
  }

  const subPrev = getPrev(subList);

  return list[0] - subPrev;
}

for (const line of input) {
  const parsedLine = line.split(" ").map((v) => parseInt(v));
  nexts.push(getNext(parsedLine));
  prevs.push(getPrev(parsedLine));
}

console.log(sum(nexts));
console.log(sum(prevs));
