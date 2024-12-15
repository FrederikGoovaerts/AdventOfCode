import { asList } from "../utils/inputReader";

// const yMax = 6;
// const xMax = 10;
// const input = asList("ex1").map(toRobot);
// const input = asList("ex2").map(toRobot);

const yMax = 102;
const xMax = 100;
const input = asList("input").map(toRobot);

interface Robot {
  px: number;
  py: number;
  vx: number;
  vy: number;
}

function toRobot(input: string): Robot {
  const [pPart, vPart] = input.substring(2).split(" v=");
  const [px, py] = pPart.split(",").map(Number);
  const [vx, vy] = vPart.split(",").map(Number);

  return {
    px,
    py,
    vx: vx < 0 ? vx + xMax + 1 : vx,
    vy: vy < 0 ? vy + yMax + 1 : vy,
  };
}

function moveRobot(r: Robot, seconds: number): void {
  let newX = (r.px + r.vx * seconds) % (xMax + 1);
  if (newX < 0) {
    newX += xMax;
  }

  let newY = (r.py + r.vy * seconds) % (yMax + 1);
  if (newY < 0) {
    newY += yMax;
  }

  r.px = newX;
  r.py = newY;
}

const robots = structuredClone(input);

for (const robot of robots) {
  moveRobot(robot, 100);
}

const q1 = robots.filter((r) => r.px < xMax / 2 && r.py < yMax / 2);
const q2 = robots.filter((r) => r.px > xMax / 2 && r.py < yMax / 2);
const q3 = robots.filter((r) => r.px < xMax / 2 && r.py > yMax / 2);
const q4 = robots.filter((r) => r.px > xMax / 2 && r.py > yMax / 2);

console.log(q1.length * q2.length * q3.length * q4.length);

// Part 2: Solved semi-manually while visualising and checking for recurring patterns (every 101 cycles with an initial offset of 18)

function visualize(robots: Robot[]): void {
  for (let y = 0; y <= yMax; y++) {
    let line = "";
    for (let x = 0; x <= xMax; x++) {
      if (robots.some((r) => r.px === x && r.py === y)) {
        line += "X";
      } else {
        line += ".";
      }
    }
    console.log(line);
  }
  console.log();
}

const treeBots = structuredClone(input);

const treeCount = 7492;
for (const robot of treeBots) {
  moveRobot(robot, treeCount);
}
console.log(treeCount);
visualize(treeBots);
