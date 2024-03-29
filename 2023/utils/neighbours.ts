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
    (val) => field[val.y][val.x]
  );
}

export function getDirectionNeighboursLocs(
  x: number,
  y: number,
  field: unknown[][],
  diagonally = false
): Partial<Record<Direction, { x: number; y: number }>> {
  if (diagonally) {
    throw new Error("You haven't implemented this yet");
  }

  const all: Record<Direction, { x: number; y: number }> = {
    down: { x: x, y: y + 1 },
    right: { x: x + 1, y: y },
    left: { x: x - 1, y: y },
    up: { x: x, y: y - 1 },
  };

  const result: Partial<Record<Direction, { x: number; y: number }>> = {};
  for (const [direction, neighbour] of Object.entries(all)) {
    if (field[neighbour.y]?.[neighbour.x] !== undefined) {
      result[direction as Direction] = neighbour;
    }
  }

  return result;
}

export function getNeighboursLocs(
  x: number,
  y: number,
  field: unknown[][],
  diagonally = true
): { x: number; y: number }[] {
  const result: { x: number; y: number }[] = diagonally
    ? [
        { x: x - 1, y: y - 1 },
        { x: x - 1, y: y },
        { x: x - 1, y: y + 1 },
        { x: x, y: y - 1 },
        { x: x, y: y + 1 },
        { x: x + 1, y: y - 1 },
        { x: x + 1, y: y },
        { x: x + 1, y: y + 1 },
      ]
    : [
        { x: x, y: y + 1 },
        { x: x + 1, y: y },
        { x: x - 1, y: y },
        { x: x, y: y - 1 },
      ];

  return result.filter((val) => field[val.y]?.[val.x] !== undefined);
}
