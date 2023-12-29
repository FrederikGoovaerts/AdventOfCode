import { asIs } from "../utils/inputReader";
const input = asIs("input").split("\n\n");

function solve(desiredDifferences: number): void {
  function checkMatch(lines: string[], i: number): boolean {
    let differences = 0;

    for (let offSet = 0; ; offSet++) {
      const upper = lines[i - offSet];
      const lower = lines[i + 1 + offSet];
      if (upper === undefined || lower === undefined) {
        break;
      } else if (upper !== lower) {
        for (let i = 0; i < upper.length; i++) {
          if (upper[i] !== lower[i]) {
            differences++;
            if (differences > desiredDifferences) {
              return false;
            }
          }
        }
      }
    }

    return differences === desiredDifferences;
  }

  function checkHorizontal(lines: string[]): number {
    for (let i = 0; i < lines.length - 1; i++) {
      if (checkMatch(lines, i)) {
        return i + 1;
      }
    }
    return -1;
  }

  function getFlipped(lines: string[]): string[] {
    const result: string[] = [];
    for (let i = 0; i < lines[0].length; i++) {
      const lineChars: string[] = [];
      for (const line of lines) {
        lineChars.push(line[i]);
      }
      result.push(lineChars.join(""));
    }

    return result;
  }

  function checkVertical(lines: string[]): number {
    return checkHorizontal(getFlipped(lines));
  }

  let result = 0;

  for (const field of input) {
    const lines = field.split("\n");

    const h = checkHorizontal(lines);
    if (h !== -1) {
      result += 100 * h;
    }
    const v = checkVertical(lines);
    if (v !== -1) {
      result += v;
    }
  }
  console.log(result);
}

solve(0);
solve(1);
