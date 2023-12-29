import { asList } from "../utils/inputReader";
import { Direction } from "../utils/neighbours";

const input = asList("input");

interface Instruction {
  dist: number;
  dir: Direction;
}

interface DistAndCorners {
  prevTurn: Turn;
  dist: number;
  nextTurn: Turn;
  name: string;
}

type Turn = "CW" | "CCW";

function dirsToTurn(from: Direction, to: Direction): Turn {
  if (
    (from === "up" && to === "right") ||
    (from === "right" && to === "down") ||
    (from === "down" && to === "left") ||
    (from === "left" && to === "up")
  ) {
    return "CW";
  } else {
    return "CCW";
  }
}

const dirMap: Record<string, Direction> = {
  D: "down",
  U: "up",
  L: "left",
  R: "right",
};

const dirMapNumbers: Record<string, Direction> = {
  "1": "down",
  "3": "up",
  "2": "left",
  "0": "right",
};

function solve(instructions: Instruction[]) {

  const turns: DistAndCorners[] = [];

  for (let i = 0; i < instructions.length; i++) {
    const prev = instructions.at(i - 1)!;
    const curr = instructions.at(i)!;
    const next = instructions.at((i + 1) % instructions.length)!;

    turns.push({
      prevTurn: dirsToTurn(prev.dir, curr.dir),
      dist: curr.dist,
      nextTurn: dirsToTurn(curr.dir, next.dir),
      name: `${i}`,
    });
  }

  const cwOverCount =
    turns.filter((t) => t.nextTurn === "CW").length -
    turns.filter((t) => t.nextTurn === "CCW").length;

  const clockWisePath = (
    cwOverCount === 4
      ? turns
      : turns.reverse().map(
          (v, i): DistAndCorners => ({
            dist: v.dist,
            prevTurn: v.prevTurn === "CW" ? "CCW" : "CW",
            nextTurn: v.nextTurn === "CW" ? "CCW" : "CW",
            name: `${i}`,
          })
        )
  ).map((v) => {
    let newDist = v.dist;
    if (v.prevTurn === "CW" && v.nextTurn === "CW") {
      newDist++;
    }
    if (v.prevTurn === "CCW" && v.nextTurn === "CCW") {
      newDist--;
    }
    return {
      ...v,
      dist: newDist,
    };
  });


  let addedSpace = 0;
  let removedSpace = 0;

  while (clockWisePath.length > 4) {
    for (let i = 0; i < clockWisePath.length; i++) {
      const prev = clockWisePath.at(i - 1)!;
      const curr = clockWisePath.at(i)!;
      const next = clockWisePath.at((i + 1) % clockWisePath.length)!;

      if (curr.nextTurn === curr.prevTurn) {
        const shift = Math.min(prev.dist, next.dist);
        if (curr.nextTurn === "CW") {
          removedSpace += shift * curr.dist;
        } else {
          addedSpace += shift * curr.dist;
        }
        prev.dist -= shift;
        next.dist -= shift;
      }
    }

    if (clockWisePath[0].dist === 0) {
      const last = clockWisePath.pop()!;
      clockWisePath.unshift(last);
    }

    let i = 1;

    while (i < clockWisePath.length - 1) {
      const curr = clockWisePath[i];
      if (curr.dist === 0 && curr.nextTurn !== curr.prevTurn) {
        const prev = clockWisePath.at(i - 1)!;
        const next = clockWisePath.at((i + 1) % clockWisePath.length)!;
        clockWisePath.splice(i - 1, 3, {
          nextTurn: next.nextTurn,
          dist: prev.dist + next.dist,
          prevTurn: prev.prevTurn,
          name: `${prev.name}|${curr.name}|${next.name}`,
        });
      } else {
        i++;
      }
    }
  }

  console.log(
    removedSpace - addedSpace + clockWisePath[0].dist * clockWisePath[1].dist
  );
}

const basicInstructions = input.map((line): Instruction => {
  const split = line.split(" ");
  return {
    dist: parseInt(split[1]),
    dir: dirMap[split[0]],
  };
});

solve(basicInstructions);

const largeInstructions = input.map((line): Instruction => {
  const hex = line.split(" ")[2].slice(2);
  return {
    dist: Number(`0x${hex.slice(0, 5)}`),
    dir: dirMapNumbers[hex[5]],
  };
});

solve(largeInstructions);
