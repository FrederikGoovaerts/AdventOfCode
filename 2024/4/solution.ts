import { as2d } from "../utils/inputReader";

const input = as2d("input");

function checkHorizontal(y: number, x: number): boolean {
  return (
    (input[y][x] === "X" &&
      input[y][x + 1] === "M" &&
      input[y][x + 2] === "A" &&
      input[y][x + 3] === "S") ||
    (input[y][x] === "S" &&
      input[y][x + 1] === "A" &&
      input[y][x + 2] === "M" &&
      input[y][x + 3] === "X")
  );
}

function checkVertical(y: number, x: number): boolean {
  return (
    (input[y][x] === "X" &&
      input[y + 1][x] === "M" &&
      input[y + 2][x] === "A" &&
      input[y + 3][x] === "S") ||
    (input[y][x] === "S" &&
      input[y + 1][x] === "A" &&
      input[y + 2][x] === "M" &&
      input[y + 3][x] === "X")
  );
}

function checkDiagonal(y: number, x: number): boolean {
  return (
    (input[y][x] === "X" &&
      input[y + 1][x + 1] === "M" &&
      input[y + 2][x + 2] === "A" &&
      input[y + 3][x + 3] === "S") ||
    (input[y][x] === "S" &&
      input[y + 1][x + 1] === "A" &&
      input[y + 2][x + 2] === "M" &&
      input[y + 3][x + 3] === "X")
  );
}

function checkCounterDiagonal(y: number, x: number): boolean {
  return (
    (input[y + 3][x] === "X" &&
      input[y + 2][x + 1] === "M" &&
      input[y + 1][x + 2] === "A" &&
      input[y][x + 3] === "S") ||
    (input[y + 3][x] === "S" &&
      input[y + 2][x + 1] === "A" &&
      input[y + 1][x + 2] === "M" &&
      input[y][x + 3] === "X")
  );
}

function checkX(y: number, x: number): boolean {
  if (input[y + 1][x + 1] !== "A") {
    return false;
  }

  const diagSam =
    (input[y][x] === "S" && input[y + 2][x + 2] === "M") ||
    (input[y][x] === "M" && input[y + 2][x + 2] === "S");
  const counterDiagSam =
    (input[y][x + 2] === "S" && input[y + 2][x] === "M") ||
    (input[y][x + 2] === "M" && input[y + 2][x] === "S");

  return diagSam && counterDiagSam;
}

let total = 0;
let xTotal = 0;

for (let y = 0; y < input.length; y++) {
  for (let x = 0; x < input[0].length; x++) {
    if (x < input[0].length - 3 && checkHorizontal(y, x)) {
      total++;
    }

    if (y < input.length - 3 && checkVertical(y, x)) {
      total++;
    }

    if (
      y < input.length - 3 &&
      x < input[0].length - 3 &&
      checkDiagonal(y, x)
    ) {
      total++;
    }

    if (
      y < input.length - 3 &&
      x < input[0].length - 3 &&
      checkCounterDiagonal(y, x)
    ) {
      total++;
    }

    if (y < input.length - 2 && x < input[0].length - 2 && checkX(y, x)) {
      xTotal++;
    }
  }
}

console.log(total);
console.log(xTotal);
