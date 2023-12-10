import { sum } from "lodash";
import { asList } from "../utils/inputReader";
import { getNeighbours, getNeighboursLocs } from "../utils/neighbours";

const input = asList("input");

interface Num {
  value: number;
  y: number;
  xStart: number;
  xEnd: number;
}

const board: string[][] = [];
const numbers: Num[] = [];
const gears: { y: number; x: number }[] = [];

input.forEach((line, i) => {
  board.push(line.split(""));
  const matches = line.matchAll(/\d+/g);

  for (const match of matches) {
    numbers.push({
      value: parseInt(match[0]),
      y: i,
      xStart: match.index!,
      xEnd: match.index! + match[0].length - 1,
    });
  }

  const gearMatches = line.matchAll(/\*/g);

  for (const gear of gearMatches) {
    gears.push({ y: i, x: gear.index! });
  }
});

const partNumbers: number[] = [];

for (const num of numbers) {
  let isPart = false;
  for (let i = num.xStart; i <= num.xEnd; i++) {
    const neigbours = getNeighbours(i, num.y, board, true);
    if (neigbours.some((ne) => !/\d|\./.test(ne))) {
      isPart = true;
      break;
    }
  }
  if (isPart) {
    partNumbers.push(num.value);
  }
}

console.log(sum(partNumbers));

const ratios = [];

for (const gear of gears) {
  const neighbourLocs = getNeighboursLocs(gear.x, gear.y, board, true);
  const touchingNumbers = numbers.filter((nu) =>
    neighbourLocs.some(
      (ne) => ne.y === nu.y && ne.x >= nu.xStart && ne.x <= nu.xEnd
    )
  );
  if (touchingNumbers.length === 2) {
    ratios.push(touchingNumbers[0].value * touchingNumbers[1].value);
  }
}
console.log(sum(ratios));
