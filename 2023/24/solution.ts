import { asList } from "../utils/inputReader";

// const [low, high, input] = [7, 27, asList("ex1")];
const [low, high, input] = [200000000000000, 400000000000000, asList("input")];

interface Stone {
  x: number;
  y: number;
  z: number;
  vx: number;
  vy: number;
  vz: number;
}

// Represents a hail stone when only looking at the 2D x-y plane, in the form of a*x + b = y
interface SimpleStone {
  a: number;
  b: number;
}

const stones: Stone[] = [];
const simpleStones: SimpleStone[] = [];

for (const line of input) {
  const matches = line.match(/(.*), (.*), (.*) @ (.*), (.*), (.*)/)!;

  const stone = {
    x: parseInt(matches[1]),
    y: parseInt(matches[2]),
    z: parseInt(matches[3]),
    vx: parseInt(matches[4]),
    vy: parseInt(matches[5]),
    vz: parseInt(matches[6]),
  };

  const yForZeroX = stone.y - (stone.x / stone.vx) * stone.vy;

  stones.push(stone);
  simpleStones.push({
    a: stone.vy / stone.vx,
    b: yForZeroX,
  });
}

let crossCount = 0;

function isFutureDirection(x: number, s: Stone): boolean {
  return x - s.x > 0 === s.vx > 0;
}

for (let s1i = 0; s1i < stones.length; s1i++) {
  for (let s2i = s1i + 1; s2i < stones.length; s2i++) {
    const s1Full = stones[s1i];
    const s1 = simpleStones[s1i];
    const s2Full = stones[s2i];
    const s2 = simpleStones[s2i];

    const crossX = (s2.b - s1.b) / (s1.a - s2.a);
    const crossY = s1.a * crossX + s1.b;

    if (
      isFutureDirection(crossX, s1Full) &&
      isFutureDirection(crossX, s2Full)
    ) {
      if (crossX >= low && crossX <= high && crossY >= low && crossY <= high) {
        crossCount++;
      }
    }
  }
}
console.log(crossCount);
