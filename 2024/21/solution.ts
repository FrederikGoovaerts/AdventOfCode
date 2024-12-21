import { asList } from "../utils/inputReader";

const input = asList("input");

type NumButton =
  | "0"
  | "1"
  | "2"
  | "3"
  | "4"
  | "5"
  | "6"
  | "7"
  | "8"
  | "9"
  | "A";

const numPaths: Record<NumButton, Record<NumButton, string[]>> = {
  "0": {
    "0": [""],
    "1": ["^<"],
    "2": ["^"],
    "3": ["^>"],
    "4": ["^^<", "^<^"],
    "5": ["^^"],
    "6": ["^^>", "^>^", ">^^"],
    "7": ["^<^^", "^^<^", "^^^<"],
    "8": ["^^^"],
    "9": ["^>^^", "^^>^", "^^^>"],
    A: [">"],
  },
  "1": {
    "0": [">v"],
    "1": [""],
    "2": [">"],
    "3": [">>"],
    "4": ["^"],
    "5": [">^", "^>"],
    "6": ["^>>", ">^>", ">>^"],
    "7": ["^^"],
    "8": ["^^>", "^>^", ">^^"],
    "9": [">>^^", ">^>^", ">^^>", "^>>^", "^>^>", "^^>>"],
    A: [">>v", ">v>"],
  },
  "2": {
    "0": ["v"],
    "1": ["<"],
    "2": [""],
    "3": [">"],
    "4": ["^<", "<^"],
    "5": ["^"],
    "6": ["^>", ">^"],
    "7": ["^^<", "^<^", "<^^"],
    "8": ["^^"],
    "9": ["^^>", "^>^", ">^^"],
    A: [">v", "v>"],
  },
  "3": {
    "0": ["<v", "v<"],
    "1": ["<<"],
    "2": ["<"],
    "3": [""],
    "4": ["^<<", "<^<", "<<^"],
    "5": ["<^", "^<"],
    "6": ["^"],
    "7": ["<<^^", "<^<^", "<^^<", "^<<^", "^<^<", "^^<<"],
    "8": ["^^<", "^<^", "<^^"],
    "9": ["^^"],
    A: ["v"],
  },
  "4": {
    "0": [">vv", "v>v"],
    "1": ["v"],
    "2": [">v", "v>"],
    "3": [">>v", ">v>", "v>>"],
    "4": [""],
    "5": [">"],
    "6": [">>"],
    "7": ["^"],
    "8": [">^", "^>"],
    "9": ["^>>", ">^>", ">>^"],
    A: [">>vv", ">v>v", ">vv>", "v>>v", "v>v>"],
  },
  "5": {
    "0": ["vv"],
    "1": ["<v", "v<"],
    "2": ["v"],
    "3": [">v", "v>"],
    "4": ["<"],
    "5": [""],
    "6": [">"],
    "7": ["^<", "<^"],
    "8": ["^"],
    "9": ["^>", ">^"],
    A: [">vv", "v>v", "vv>"],
  },
  "6": {
    "0": ["<vv", "v<v", "vv<"],
    "1": ["v<<", "<v<", "<<v"],
    "2": ["<v", "v<"],
    "3": ["v"],
    "4": ["<<"],
    "5": ["<"],
    "6": [""],
    "7": ["^<<", "<^<", "<<^"],
    "8": ["<^", "^<"],
    "9": ["^"],
    A: ["vv"],
  },
  "7": {
    "0": [">vvv", "v>vv", "vv>v"],
    "1": ["vv"],
    "2": [">vv", "v>v", "vv>"],
    "3": [">>vv", ">v>v", ">vv>", "v>>v", "v>v>", "vv>>"],
    "4": ["v"],
    "5": [">v", "v>"],
    "6": [">>v", ">v>", "v>>"],
    "7": [""],
    "8": [">"],
    "9": [">>"],
    A: [
      ">>vvv",
      ">v>vv",
      "v>>vv",
      ">vv>v",
      "v>v>v",
      "vv>>v",
      ">vvv>",
      "v>vv>",
      "vv>v>",
    ],
  },
  "8": {
    "0": ["vvv"],
    "1": ["<vv", "v<v", "vv<"],
    "2": ["vv"],
    "3": [">vv", "v>v", "vv>"],
    "4": ["<v", "v<"],
    "5": ["v"],
    "6": [">v", "v>"],
    "7": ["<"],
    "8": [""],
    "9": [">"],
    A: [">vvv", "v>vv", "vv>v", "vvv>"],
  },
  "9": {
    "0": ["vvv<", "vv<v", "v<vv", "<vvv"],
    "1": ["<<vv", "<v<v", "<vv<", "v<<v", "v<v<", "vv<<"],
    "2": ["<vv", "v<v", "vv<"],
    "3": ["vv"],
    "4": ["v<<", "<v<", "<<v"],
    "5": ["<v", "v<"],
    "6": ["v"],
    "7": ["<<"],
    "8": ["<"],
    "9": [""],
    A: ["vvv"],
  },
  A: {
    "0": ["<"],
    "1": ["^<<", "<^<"],
    "2": ["<^", "^<"],
    "3": ["^"],
    "4": ["<^<^", "<^^<", "^<<^", "^<^<", "^^<<"],
    "5": ["^^<", "^<^", "<^^"],
    "6": ["^^"],
    "7": [
      "<^<^^",
      "^<<^^",
      "<^^<^",
      "^<^<^",
      "^^<<^",
      "<^^^<",
      "^<^^<",
      "^^<^<",
      "^^^<<",
    ],
    "8": ["^^^<", "^^<^", "^<^^", "<^^^"],
    "9": ["^^^"],
    A: [""],
  },
};

type DirButton = "^" | "A" | "<" | "v" | ">";

const dirPaths: Record<DirButton, Record<DirButton, string[]>> = {
  "^": {
    "^": [""],
    A: [">"],
    "<": ["v<"],
    v: ["v"],
    ">": ["v>", ">v"],
  },
  A: {
    "^": ["<"],
    A: [""],
    "<": ["v<<", "<v<"],
    v: ["<v", "v<"],
    ">": ["v"],
  },
  "<": {
    "^": [">^"],
    A: [">>^", ">^>"],
    "<": [""],
    v: [">"],
    ">": [">>"],
  },
  v: {
    "^": ["^"],
    A: [">^", "^>"],
    "<": ["<"],
    v: [""],
    ">": [">"],
  },
  ">": {
    "^": ["<^", "^<"],
    A: ["^"],
    "<": ["<<"],
    v: ["<"],
    ">": [""],
  },
};

const bestDirPath: Map<string, number> = new Map();

function getPressesDir(
  dir1: DirButton,
  dir2: DirButton,
  depth: number
): number {
  if (depth === 0) {
    return dirPaths[dir1][dir2][0].length + 1;
  }

  const mapKey = `${depth}|${dir1}|${dir2}`;
  if (bestDirPath.has(mapKey)) {
    return bestDirPath.get(mapKey)!;
  }

  const paths = dirPaths[dir1][dir2];

  let shortest: number | undefined = undefined;

  for (const path of paths) {
    let curr: DirButton = "A";

    let result = 0;

    for (const next of `${path}A`.split("")) {
      const presses = getPressesDir(curr, next as DirButton, depth - 1);
      result += presses;

      curr = next as DirButton;
    }

    if (shortest === undefined || result < shortest) {
      shortest = result;
    }
  }

  bestDirPath.set(mapKey, shortest!);

  return shortest!;
}

function getPressesNum(
  num1: NumButton,
  num2: NumButton,
  dirDepth: number
): number {
  const paths = numPaths[num1][num2];

  let shortest: number | undefined = undefined;

  for (const path of paths) {
    let curr: DirButton = "A";

    let result = 0;

    for (const next of `${path}A`.split("")) {
      const presses = getPressesDir(curr, next as DirButton, dirDepth);
      result += presses;

      curr = next as DirButton;
    }

    if (shortest === undefined || result < shortest) {
      shortest = result;
    }
  }

  return shortest!;
}

function getPressesFor(code: string, dirDepth: number): number {
  let curr: NumButton = "A";

  let result = 0;

  for (const next of code.split("")) {
    const presses = getPressesNum(curr, next as NumButton, dirDepth);
    result += presses;

    curr = next as NumButton;
  }

  return result;
}

let complexitySum = 0;

for (const code of input) {
  const result = getPressesFor(code, 1);

  const complexity = result * Number(code.substring(0, code.length - 1));
  complexitySum += complexity;
}

console.log(complexitySum);

complexitySum = 0;

for (const code of input) {
  const result = getPressesFor(code, 24);

  const complexity = result * Number(code.substring(0, code.length - 1));
  complexitySum += complexity;
}

console.log(complexitySum);
