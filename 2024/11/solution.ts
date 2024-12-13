import { asIs } from "../utils/inputReader";

const input = asIs("input").split(" ").map(Number);

const blinkMap: Map<string, number> = new Map();

function getBlinkCount(values: number[], count: number): number {
  if (count === 0) {
    return values.length;
  }

  let countSum = 0;

  for (const val of values) {
    const token = `${val},${count}`;
    if (blinkMap.has(token)) {
      countSum += blinkMap.get(token)!;
    } else {
      let newValues: number[];

      if (val === 0) {
        newValues = [1];
      } else if (`${val}`.length % 2 === 0) {
        const numberString = `${val}`;
        const first = Number(
          numberString.substring(0, numberString.length / 2)
        );
        const second = Number(numberString.substring(numberString.length / 2));
        newValues = [first, second];
      } else {
        newValues = [val * 2024];
      }

      const newBlinkCount = getBlinkCount(newValues, count - 1);
      countSum += newBlinkCount;
      blinkMap.set(token, newBlinkCount);
    }
  }

  return countSum;
}

console.log(getBlinkCount(input, 25));
console.log(getBlinkCount(input, 75));
