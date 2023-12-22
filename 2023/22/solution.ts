import { asList } from "../utils/inputReader";

const input = asList("input");

interface Brick {
  id: number;
  sx: number;
  sy: number;
  sz: number;
  ex: number;
  ey: number;
  ez: number;
}

interface Coords {
  x: number;
  y: number;
  z: number;
}

const bricks = new Map<number, Brick>();
const spaceUsedByBrick = new Map<string, number>();

const settledBricks = new Set<number>();
const restsOn = new Map<number, number[]>();

let maxKnownZ = 0;

function coordsForBrick(id: number): Coords[] {
  const result: Coords[] = [];
  const brick = bricks.get(id)!;

  for (let x = brick.sx; x <= brick.ex; x++) {
    for (let y = brick.sy; y <= brick.ey; y++) {
      for (let z = brick.sz; z <= brick.ez; z++) {
        result.push({ x, y, z });
      }
    }
  }
  return result;
}

function stringify(c: Coords): string {
  return `${c.x}|${c.y}|${c.z}`;
}

input.forEach((line, index) => {
  const [startCoord, endCoord] = line.split("~");
  const [sx, sy, sz] = startCoord.split(",").map((v) => parseInt(v));
  const [ex, ey, ez] = endCoord.split(",").map((v) => parseInt(v));
  bricks.set(index, { id: index, sx, sy, sz, ex, ey, ez });

  if (sz === 1 || ez === 1) {
    settledBricks.add(index);
    restsOn.set(index, []);
  }

  maxKnownZ = Math.max(maxKnownZ, sz, ez);

  for (const c of coordsForBrick(index)) {
    spaceUsedByBrick.set(stringify(c), index);
  }
});

for (let z = 2; z <= maxKnownZ; z++) {
  for (const [id, brick] of bricks.entries()) {
    if (!settledBricks.has(id) && (brick.sz === z || brick.ez === z)) {
      // Remove spaces from map
      const spaces = coordsForBrick(id);
      spaces.forEach((s) => {
        spaceUsedByBrick.delete(stringify(s));
      });

      // Check how far it drops
      let drop = 0;
      const getRestsOn = (): number[] => {
        const result: number[] = [];
        for (const space of spaces) {
          const spaceBelow = spaceUsedByBrick.get(
            stringify({ ...space, z: space.z - (drop + 1) })
          );
          if (spaceBelow !== undefined && !result.includes(spaceBelow)) {
            result.push(spaceBelow);
          }
        }
        return result;
      };
      let ro = getRestsOn();
      while (ro.length === 0 && brick.sz - drop > 1 && brick.ez - drop > 1) {
        drop++;
        ro = getRestsOn();
      }

      // Drop and add used spaces again
      restsOn.set(id, ro);
      brick.sz -= drop;
      brick.ez -= drop;
      coordsForBrick(id).forEach((s) => {
        spaceUsedByBrick.set(stringify(s), id);
      });

      // Mark as settled
      settledBricks.add(id);
    }
  }
}

let dCount = 0;

for (const b of bricks.keys()) {
  if (![...restsOn.values()].some((r) => r.length === 1 && r[0] === b)) {
    dCount++;
  }
}

console.log(dCount);
