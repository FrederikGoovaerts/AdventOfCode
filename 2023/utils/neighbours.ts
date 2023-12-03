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
