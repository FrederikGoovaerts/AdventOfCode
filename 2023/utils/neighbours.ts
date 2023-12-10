export type Direction = "up" | "down" | "left" | "right";

export function reverseDirection(dir: Direction): Direction {
  switch (dir) {
    case "up":
      return "down";
    case "down":
      return "up";
    case "left":
      return "right";
    case "right":
      return "left";
  }
}

export function getNeighbours<T = number>(
  x: number,
  y: number,
  field: T[][],
  diagonally = true
): T[] {
  return getNeighboursLocs(x, y, field, diagonally).map(
    (val) => field[val[0]][val[1]]
  );
}

export function getDirectionNeighboursLocs(
  x: number,
  y: number,
  field: unknown[][],
  diagonally = false
): Partial<Record<Direction, [number, number]>> {
  if (diagonally) {
    throw new Error("You haven't implemented this yet");
  }

  const all: Record<Direction, [number, number]> = {
    down: [x, y + 1],
    right: [x + 1, y],
    left: [x - 1, y],
    up: [x, y - 1],
  };

  const result: Partial<Record<Direction, [number, number]>> = {};
  for (const entry of Object.entries(all)) {
    if (field[entry[1][1]]?.[entry[1][0]] !== undefined) {
      result[entry[0] as Direction] = entry[1];
    }
  }

  return result;
}

export function getNeighboursLocs(
  x: number,
  y: number,
  field: unknown[][],
  diagonally = true
): [number, number][] {
  const result: [number, number][] = diagonally
    ? [
        [x - 1, y - 1],
        [x - 1, y],
        [x - 1, y + 1],
        [x, y - 1],
        [x, y + 1],
        [x + 1, y - 1],
        [x + 1, y],
        [x + 1, y + 1],
      ]
    : [
        [x, y + 1],
        [x + 1, y],
        [x - 1, y],
        [x, y - 1],
      ];

  return result.filter((val) => field[val[0]]?.[val[1]] !== undefined);
}
