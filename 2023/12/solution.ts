import { sum } from "lodash";
import { asList } from "../utils/inputReader";
const input = asList("ex1");

function replaceAt(input: string, i: number, char: string): string {
  if (i > input.length - 1) {
    return input;
  }
  return input.substring(0, i) + char + input.substring(i + 1);
}

function createResult(
  vis: string,
  numbers: string,
  choices: number[]
): string[] {
  let newVis = vis;
  for (const choice of choices) {
    newVis = replaceAt(newVis, choice, "#");
  }
  newVis = newVis.replaceAll("?", ".");

  const newVisNumberString = newVis
    .split(".")
    .filter((v) => v !== "")
    .map((v) => v.length)
    .join(",");

  if (newVisNumberString === numbers) {
    return [vis];
  } else {
    return [];
  }
}

function getArrangementsRec(
  remaining: number,
  options: number[],
  optionsIndex: number,
  choices: number[],
  vis: string,
  numbers: string
): string[] {
  if (remaining > options.length - optionsIndex) {
    return [];
  } else if (remaining === 0) {
    return createResult(vis, numbers, choices);
  }

  const results: string[] = [];
  for (let i = optionsIndex; i < options.length; i++) {
    results.push(
      ...getArrangementsRec(
        remaining - 1,
        options,
        i + 1,
        [...choices, options[i]],
        vis,
        numbers
      )
    );
  }

  return results;
}

function getArrangements(vis: string, numbers: string): string[] {
  const totalBroken = sum(numbers.split(",").map((v) => parseInt(v)));
  const totalBrokenVisible = vis.split("").filter((v) => v === "#").length;
  const options: number[] = [];
  for (let i = 0; i < vis.length; i++) {
    if (vis[i] === "?") {
      options.push(i);
    }
  }

  return getArrangementsRec(
    totalBroken - totalBrokenVisible,
    options,
    0,
    [],
    vis,
    numbers
  );
}

const arrangementsList: string[][] = [];

for (const line of input) {
  const [vis, rawNumbers] = line.split(" ");
  const arr = getArrangements(vis, rawNumbers);

  arrangementsList.push(arr);
}

console.log(sum(arrangementsList.map((a) => a.length)));

const largeArrangementsList: string[][] = [];

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

  largeArrangementsList.push(arr);
}

console.log(sum(largeArrangementsList.map((a) => a.length)));
