import { as2d } from "../utils/inputReader";

const input = as2d("input");

type Location = { y: number; x: number };
const locations: Record<string, Location[]> = {};

for (let y = 0; y < input.length; y++) {
  for (let x = 0; x < input[0].length; x++) {
    const value = input[y][x];
    if (value !== ".") {
      const valueLocations = [...(locations[value] ?? []), { x, y }];
      locations[value] = valueLocations;
    }
  }
}

function isWithinBounds(x: number, y: number): boolean {
  return x >= 0 && y >= 0 && x < input[0].length && y < input.length;
}

const antinodes1 = new Set<string>();
const antinodes2 = new Set<string>();

for (const valLocations of Object.values(locations)) {
  for (let i = 0; i < valLocations.length - 1; i++) {
    const loc1 = valLocations[i];
    for (let j = i + 1; j < valLocations.length; j++) {
      const loc2 = valLocations[j];

      const diffX = loc1.x - loc2.x;
      const diffY = loc1.y - loc2.y;

      // Part 1
      const n1: Location = { x: loc1.x + diffX, y: loc1.y + diffY };
      const n2: Location = { x: loc2.x - diffX, y: loc2.y - diffY };

      if (isWithinBounds(n1.x, n1.y)) {
        antinodes1.add(`${n1.x}|${n1.y}`);
      }

      if (isWithinBounds(n2.x, n2.y)) {
        antinodes1.add(`${n2.x}|${n2.y}`);
      }

      // Part 2
      const n3: Location = { x: loc1.x, y: loc1.y };
      while (isWithinBounds(n3.x, n3.y)) {
        antinodes2.add(`${n3.x}|${n3.y}`);

        n3.x += diffX;
        n3.y += diffY;
      }

      const n4: Location = { x: loc2.x, y: loc2.y };
      while (isWithinBounds(n4.x, n4.y)) {
        antinodes2.add(`${n4.x}|${n4.y}`);
        n4.x -= diffX;
        n4.y -= diffY;
      }
    }
  }
}

console.log(antinodes1.size);
console.log(antinodes2.size);
