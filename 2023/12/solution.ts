import { sum } from "lodash";
import { asList } from "../utils/inputReader";
const input = asList("input");

const globalMap = new Map<string, number>();

function getBeforeAndAfter(
  numbers: number[],
  i: number
): { before: number; after: number } {
  const numbersBefore = numbers.slice(0, i);
  const numbersAfter = numbers.slice(i + 1);
  return {
    before: sum(numbersBefore) + numbersBefore.length,
    after: sum(numbersAfter) + numbersAfter.length,
  };
}

function stringify(vis: string, numbers: number[]): string {
  return vis + numbers.join(",");
}

function getArrangementsRec(vis: string, numbers: number[]): number {
  if (numbers.length === 0) {
    if (vis.includes("#")) {
      return 0;
    } else {
      return 1;
    }
  }

  const stringified = stringify(vis, numbers);
  if (globalMap.has(stringified)) {
    return globalMap.get(stringified)!;
  }

  if (numbers.length === 1) {
    const n = numbers[0];
    let result = 0;

    for (let i = 0; i <= vis.length - n; i++) {
      const place = vis.slice(i, i + n);

      const a = Math.max(i, 0);
      const beforeVisSlice = vis.slice(0, a);

      const b = Math.min(i + n, vis.length);
      const afterVisSlice = vis.slice(b);

      if (
        !place.includes(".") &&
        !beforeVisSlice.includes("#") &&
        !afterVisSlice.includes("#")
      ) {
        result += 1;
      }
    }

    globalMap.set(stringified, result);
    return result;
  }

  const nextNumber = Math.max(...numbers);
  const nextNumberIndex = numbers.indexOf(nextNumber);
  const bAndA = getBeforeAndAfter(numbers, nextNumberIndex);

  let result = 0;

  for (
    let i = bAndA.before;
    i <= vis.length - (nextNumber + bAndA.after);
    i++
  ) {
    const place = vis.slice(i, i + nextNumber);
    const before = vis[i - 1];
    const after = vis[i + nextNumber];
    if (!place.includes(".") && before !== "#" && after !== "#") {
      let localResult = 1;

      const a = Math.max(i - 1, 0);
      const beforeVisSlice = vis.slice(0, a);
      const beforeNumSlice = numbers.slice(0, nextNumberIndex);

      localResult *= getArrangementsRec(beforeVisSlice, beforeNumSlice);

      const b = Math.min(i + nextNumber + 1, vis.length);
      const afterVisSlice = vis.slice(b);
      const afterNumSlice = numbers.slice(nextNumberIndex + 1);

      localResult *= getArrangementsRec(afterVisSlice, afterNumSlice);

      result += localResult;
    }
  }

  globalMap.set(stringified, result);
  return result;
}

function getArrangements(vis: string, rawNumbers: string): number {
  const rawNumbersList = rawNumbers.split(",").map((n) => parseInt(n));

  return getArrangementsRec(vis, rawNumbersList);
}

let part1Result = 0;

for (const line of input) {
  const [vis, rawNumbers] = line.split(" ");
  const arr = getArrangements(vis, rawNumbers);

  part1Result += arr;
}

console.log(part1Result);

let part2Result = 0;

for (const line of input) {
  const [vis, rawNumbers] = line.split(" ");

  let fullVis = "";
  let fullRawNumbers = "";

  for (let i = 0; i < 5; i++) {
    if (i !== 0) {
      fullVis += "?";
      fullRawNumbers += ",";
    }
    fullVis += vis;
    fullRawNumbers += rawNumbers;
  }

  const arr = getArrangements(fullVis, fullRawNumbers);

  part2Result += arr;
}

console.log(part2Result);
