import { asList } from "../utils/inputReader";

const input = asList("input");

const simpleOperatorOptions = ["+", "*"] as const;
const operatorOptions = ["+", "*", "||"] as const;
type Operator = (typeof operatorOptions)[number];

function doOp(a: number, b: number, op: Operator): number {
  switch (op) {
    case "+":
      return a + b;
    case "*":
      return a * b;
    case "||":
      return Number(`${a}${b}`);
  }
}

function isCorrect(
  curr: number,
  remaining: number[],
  testNumber: number,
  operators: readonly Operator[]
): boolean {
  if (remaining.length === 0) {
    return curr === testNumber;
  } else {
    const [firstVal, ...remainingValues] = remaining;
    for (const op of operators) {
      const result = isCorrect(
        doOp(curr, firstVal, op),
        remainingValues,
        testNumber,
        operators
      );
      if (result) {
        return true;
      }
    }
    return false;
  }
}

let calibration1 = 0;
let calibration2 = 0;

for (const line of input) {
  const splitLine = line.split(": ");
  const testNumber = Number(splitLine[0]);
  const [firstVal, ...remainingValues] = splitLine[1].split(" ").map(Number);

  if (isCorrect(firstVal, remainingValues, testNumber, simpleOperatorOptions)) {
    calibration1 += testNumber;
  }
  if (isCorrect(firstVal, remainingValues, testNumber, operatorOptions)) {
    calibration2 += testNumber;
  }
}

console.log(calibration1);
console.log(calibration2);
