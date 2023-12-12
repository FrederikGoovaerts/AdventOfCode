import { sum } from "lodash";
import { asList } from "../utils/inputReader";
const input = asList("ex2");

const marker = "V";

function getArrangementsRec(
  vis: string,
  remainingNumbers: {
    val: number;
    position?: number;
    spaceBefore: number;
    spaceAfter: number;
  }[]
): number {
  if (remainingNumbers.length === 0) {
    if (vis.split("").includes("#")) {
      return 0;
    }
    return 1;
  }

  const nextNumber = remainingNumbers[0];
  let result = 0;

  for (
    let i = Math.max(nextNumber.spaceBefore, vis.lastIndexOf("V") + 1);
    i <= vis.length - (nextNumber.val + nextNumber.spaceAfter);
    i++
  ) {
    const place = vis.slice(i, i + nextNumber.val);
    const before = vis[i - 1];
    const after = vis[i + nextNumber.val];
    if (
      !place.includes(".") &&
      !place.includes(marker) &&
      !["#", marker].includes(before) &&
      !["#", marker].includes(after)
    ) {
      const updated =
        vis.slice(0, i) +
        "".padStart(nextNumber.val, marker) +
        vis.slice(i + nextNumber.val);

      const subRes = getArrangementsRec(updated, remainingNumbers.slice(1));

      result += subRes;
    }
  }

  return result;
}

function getArrangements(vis: string, rawNumbers: string): number {
  const rawNumbersList = rawNumbers.split(",").map((n) => parseInt(n));
  const numbers: {
    val: number;
    position?: number;
    spaceBefore: number;
    spaceAfter: number;
  }[] = [];

  for (let i = 0; i < rawNumbersList.length; i++) {
    const numbersBefore = rawNumbersList.slice(0, i);
    const numbersAfter = rawNumbersList.slice(i + 1);
    numbers.push({
      val: rawNumbersList[i],
      spaceBefore: sum(numbersBefore) + numbersBefore.length,
      spaceAfter: sum(numbersAfter) + numbersAfter.length,
    });
  }

  return getArrangementsRec(vis, numbers);
}

let part1Result = 0;

for (const line of input) {
  const [vis, rawNumbers] = line.split(" ");
  const arr = getArrangements(vis, rawNumbers);

  part1Result += arr;
}

console.log(part1Result);

let part2Result = 0;
let count = 0;

for (const line of input) {
  console.log(++count);
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
