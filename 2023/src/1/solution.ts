import { sum } from "lodash";
import { asList } from "../utils/inputReader";

const input = asList();

const values: number[] = [];

for (const line of input) {
  const digitOnlyLine = line.replaceAll(/[a-z]/g, "");
  const digits = digitOnlyLine[0] + digitOnlyLine[digitOnlyLine.length - 1];
  values.push(parseInt(digits));
}
console.log(sum(values));

const numberMap: Record<string, string> = {
  one: "1",
  two: "2",
  three: "3",
  four: "4",
  five: "5",
  six: "6",
  seven: "7",
  eight: "8",
  nine: "9",
};

const newValues: number[] = [];
const regex = /(?=([1-9]|one|two|three|four|five|six|seven|eight|nine))/g;

for (const line of input) {
  const matches = [...line.matchAll(regex)]

    const first = matches.at(0)![1];
    const last = matches.at(-1)![1];
    const newValue = parseInt(
      (numberMap[first] ?? first) + (numberMap[last] ?? last)
    );
    newValues.push(newValue);
}
console.log(sum(newValues));
