import { asList } from "../utils/inputReader";
const input = asList("input");
const space = input.map((line) => line.split(""));

const maxY = input.length - 1;
const maxX = input[0].length - 1;
const emptyRows = new Set(Array.from(Array(maxY).keys()));
const emptyColumns = new Set(Array.from(Array(maxX).keys()));

interface Galaxy {
  x: number;
  y: number;
}

const galaxies: Galaxy[] = [];

for (let y = 0; y <= maxX; y++) {
  for (let x = 0; x <= maxX; x++) {
    if (space[y][x] === "#") {
      emptyColumns.delete(x);
      emptyRows.delete(y);

      galaxies.push({ x, y });
    }
  }
}

function solve(expansion: number) {
  let dist = 0;

  const getDistanceBetween = (gA: Galaxy, gB: Galaxy): number => {
    const distX = Math.abs(gA.x - gB.x);
    const distY = Math.abs(gA.y - gB.y);

    let expansionX = 0;

    for (const col of emptyColumns) {
      if ((gA.x < col && col < gB.x) || (gB.x < col && col < gA.x)) {
        expansionX += expansion;
      }
    }

    let expansionY = 0;

    for (const row of emptyRows) {
      if ((gA.y < row && row < gB.y) || (gB.y < row && row < gA.y)) {
        expansionY += expansion;
      }
    }

    return distX + distY + expansionX + expansionY;
  };

  for (let i1 = 0; i1 < galaxies.length; i1++) {
    for (let i2 = i1 + 1; i2 < galaxies.length; i2++) {
      dist += getDistanceBetween(galaxies[i1], galaxies[i2]);
    }
  }

  console.log(dist);
}

solve(1);
solve(999999);
