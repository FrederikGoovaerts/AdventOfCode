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

function hits(a: Stone, b: Stone): boolean {
  if (a.vx === b.vx) {
    return a.x === b.x;
  }
  const xHitTime = (b.x - a.x) / (a.vx - b.vx);
  if (xHitTime < 0) {
    return false;
  }

  return (
    a.y + xHitTime * a.vy === b.y + xHitTime * b.vy &&
    a.z + xHitTime * a.vz === b.z + xHitTime * b.vz
  );
}

const firstReference = stones[0];
const secondReference = stones.find(
  (s) => s !== firstReference && !hits(firstReference, s)
)!;

let done = false;
for (let t1 = 1; !done; t1++) {
  for (let t2 = 0; t2 < t1; t2++) {
    const candidateOne = constructStone(
      t2,
      t1,
      firstReference,
      secondReference
    );
    if (candidateOne && stones.every((s) => hits(s, candidateOne))) {
      console.log(candidateOne);
      done = true;
    } else {
      const candidateTwo = constructStone(
        t2,
        t1,
        firstReference,
        secondReference
      );
      if (candidateTwo && stones.every((s) => hits(s, candidateTwo))) {
        console.log(candidateTwo);
        done = true;
      }
    }
  }
}

function constructStone(
  tFirst: number,
  tSecond: number,
  first: Stone,
  second: Stone
): Stone | undefined {
  const firstPos = {
    x: first.x + tFirst * first.vx,
    y: first.y + tFirst * first.vy,
    z: first.z + tFirst * first.vz,
  };
  const secondPos = {
    x: second.x + tSecond * second.vx,
    y: second.y + tSecond * second.vy,
    z: second.z + tSecond * second.vz,
  };
  const vx = (secondPos.x - firstPos.x) / (tSecond - tFirst);
  const vy = (secondPos.y - firstPos.y) / (tSecond - tFirst);
  const vz = (secondPos.z - firstPos.z) / (tSecond - tFirst);
  if (Number.isInteger(vx) && Number.isInteger(vy) && Number.isInteger(vz)) {
    return {
      x: first.x - vx * tFirst,
      y: first.y - vy * tFirst,
      z: first.z - vz * tFirst,
      vx,
      vy,
      vz,
    };
  }
  return undefined;
}
