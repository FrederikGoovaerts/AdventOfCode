import { as2d } from "../utils/inputReader";
import {
  Direction,
  getNeighboursLocs,
  turnClockwise,
  turnCounterclockwise,
} from "../utils/neighbours";

const input = as2d("input");

const handledSquares = new Set<string>();

type Location = { x: number; y: number };

function getPlotFor(startY: number, startX: number): Location[] {
  const val = input[startY][startX];
  const result = [{ x: startX, y: startY }];
  const toHandle = [{ x: startX, y: startY }];
  const seen = new Set<string>();
  seen.add(`${startY}|${startX}`);

  while (toHandle.length > 0) {
    const curr = toHandle.pop()!;
    const neighbours = getNeighboursLocs(curr.x, curr.y, input, false).filter(
      (loc) => !seen.has(`${loc.y}|${loc.x}`)
    );

    for (const n of neighbours) {
      if (input[n.y][n.x] === val) {
        result.push({ x: n.x, y: n.y });
        toHandle.push({ x: n.x, y: n.y });
        seen.add(`${n.y}|${n.x}`);
      }
    }
  }

  return result;
}

const dirMap: Record<Direction, { baseCheck: Location; ccwCheck: Location }> = {
  right: {
    baseCheck: { x: 1, y: 0 },
    ccwCheck: { x: 1, y: -1 },
  },
  down: {
    baseCheck: { x: 0, y: 1 },
    ccwCheck: { x: 1, y: 1 },
  },
  left: {
    baseCheck: { x: -1, y: 0 },
    ccwCheck: { x: -1, y: 1 },
  },
  up: {
    baseCheck: { x: 0, y: -1 },
    ccwCheck: { x: -1, y: -1 },
  },
};

// Unused, does not handle insides
function getNbSides(startY: number, startX: number): number {
  const val = input[startY][startX];

  let nbOfSides = 1;
  let x = startX;
  let y = startY;
  let dir: Direction = "right";

  while (nbOfSides < 4 || !(x === startX && y === startY && dir === "up")) {
    const currChecks = dirMap[dir];
    const baseOk =
      input[y + currChecks.baseCheck.y]?.[x + currChecks.baseCheck.x] === val;

    if (!baseOk) {
      nbOfSides++;
      dir = turnClockwise(dir);
    } else {
      const ccwTurn =
        input[y + currChecks.ccwCheck.y]?.[x + currChecks.ccwCheck.x] === val;

      if (!ccwTurn) {
        x = x + currChecks.baseCheck.x;
        y = y + currChecks.baseCheck.y;
      } else {
        nbOfSides++;
        dir = turnCounterclockwise(dir);
        x = x + currChecks.ccwCheck.x;
        y = y + currChecks.ccwCheck.y;
      }
    }
  }

  return nbOfSides;
}

let price = 0;
let modernPrice = 0;

for (let y = 0; y < input.length; y++) {
  for (let x = 0; x < input[0].length; x++) {
    if (!handledSquares.has(`${y}|${x}`)) {
      const val = input[y][x];
      const plot = getPlotFor(y, x);

      let fences = 0;

      for (const l of plot) {
        handledSquares.add(`${l.y}|${l.x}`);
        const neighbours = getNeighboursLocs(l.x, l.y, input, false);
        fences += 4 - neighbours.length;

        for (const n of neighbours) {
          if (input[n.y][n.x] !== val) {
            fences++;
          }
        }
      }

      let modernFences = fences;

      for (const l of plot) {
        const upVal = input[l.y - 1]?.[l.x];
        const rightVal = input[l.y]?.[l.x + 1];
        const downVal = input[l.y + 1]?.[l.x];
        const leftVal = input[l.y]?.[l.x - 1];

        // Check right
        if (rightVal === val) {
          const rightUpVal = input[l.y - 1]?.[l.x + 1];
          const rightDownVal = input[l.y + 1]?.[l.x + 1];
          if (rightUpVal !== val && upVal !== val) {
            modernFences--;
          }
          if (rightDownVal !== val && downVal !== val) {
            modernFences--;
          }
        }

        // Check down
        if (downVal === val) {
          const downRightVal = input[l.y + 1]?.[l.x + 1];
          const downLeftVal = input[l.y + 1]?.[l.x - 1];
          if (downRightVal !== val && rightVal !== val) {
            modernFences--;
          }
          if (downLeftVal !== val && leftVal !== val) {
            modernFences--;
          }
        }
      }

      price += plot.length * fences;
      modernPrice += plot.length * modernFences;
    }
  }
}

console.log(price);
console.log(modernPrice);
