import { asList } from "../utils/inputReader";

const input = asList("input");

let regA = BigInt(input[0].substring(12));
let regB = 0n;
let regC = 0n;

const program = input[4].substring(9).split(",").map(BigInt);
const out: bigint[] = [];

let pointer = 0;

function combo(val: bigint): bigint {
  switch (val) {
    case 0n:
    case 1n:
    case 2n:
    case 3n:
      return val;
    case 4n:
      return regA;
    case 5n:
      return regB;
    case 6n:
      return regC;
  }
  throw Error("illegal combo");
}

function doOp(code: bigint, operand: bigint): bigint | undefined {
  if (code === 0n) {
    // adv
    const numerator = regA;
    const denominator = 2n ** combo(operand);
    const result = numerator / denominator;
    regA = result;

    return undefined;
  } else if (code === 1n) {
    // bxl
    const result = regB ^ operand;
    regB = result;

    return undefined;
  } else if (code === 2n) {
    // bst
    const result = combo(operand) % 8n;
    regB = result;

    return undefined;
  } else if (code === 3n) {
    // jnz
    if (regA !== 0n) {
      return operand;
    }

    return undefined;
  } else if (code === 4n) {
    // bxc
    const result = regB ^ regC;
    regB = result;

    return undefined;
  } else if (code === 5n) {
    // out
    const result = combo(operand) % 8n;
    out.push(result);

    return undefined;
  } else if (code === 6n) {
    // bdv
    const numerator = regA;
    const denominator = 2n ** combo(operand);
    const result = numerator / denominator;
    regB = result;

    return undefined;
  } else if (code === 7n) {
    // cdv
    const numerator = regA;
    const denominator = 2n ** combo(operand);
    const result = numerator / denominator;
    regC = result;

    return undefined;
  }

  throw new Error("illegal op");
}

function run() {
  while (pointer < program.length) {
    const oc = program[pointer];
    const op = program[pointer + 1];

    const newPoint = doOp(oc, op);

    if (newPoint !== undefined) {
      pointer = Number(String(newPoint));
    } else {
      pointer += 2;
    }
  }
}

run();

console.log(out.join(","));

// Part 2

// The program step by step:
// 2,4 - B = A mod 8
// 1,1 - B = B XOR 1
// 7,5 - C = A / 2^B
// 4,4 - B = B XOR C
// 1,4 - B = B XOR 4
// 0,3 - A = A / 8
// 5,5 - Output B % 8
// 3,0 - Start over or end if A is 0

// This means that B, before output, is equal to:
// (((A % 8) XOR 1) XOR (A / 2^((A % 8) XOR 1))) XOR 4

const reversedProgram = program.toReversed();

function recGet(step: number, lastA: bigint): bigint | undefined {
  if (step === reversedProgram.length) {
    return lastA;
  }

  const b = reversedProgram[step];

  for (let a = lastA * 8n; a < (lastA + 1n) * 8n; a++) {
    const run = a % 8n ^ 1n ^ (a / 2n ** (a % 8n ^ 1n)) ^ 4n;

    if (run % 8n === b) {
      const result = recGet(step + 1, a);

      if (result !== undefined) {
        return result;
      }
    }
  }

  return undefined;
}

console.log(recGet(0, 0n)!.toString());
