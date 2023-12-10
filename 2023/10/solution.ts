import { asList } from "../utils/inputReader";
import {
  Direction,
  getDirectionNeighboursLocs,
  reverseDirection,
} from "../utils/neighbours";
const input = asList("input");

type mapSymbols = "." | "S" | "|" | "-" | "J" | "L" | "F" | "7";

function connects(char: string, dir: Direction): boolean {
  if (dir === "up") {
    return ["|", "J", "L"].includes(char);
  } else if (dir === "down") {
    return ["|", "F", "7"].includes(char);
  } else if (dir === "left") {
    return ["-", "J", "7"].includes(char);
  } else if (dir === "right") {
    return ["-", "F", "L"].includes(char);
  }
  return false;
}

const map: mapSymbols[][] = [];
const path: { x: number; y: number }[] = [];

input.forEach((line, y) => {
  const splitLine = line.split("") as mapSymbols[];
  map.push(splitLine);

  const sIndex = splitLine.indexOf("S");

  if (sIndex > -1) {
    path.push({ x: sIndex, y: y });
  }
});

// Choose random starting direction of path
let curr: { x: number; y: number } = { x: -1, y: -1 };
// for (const [dir, loc] of Object.entries(
//   getDirectionNeighboursLocs(path[0].x, path[0].y, map)
// )) {
//   const char = map[loc[1]][loc[0]];
//   if (connects(char, reverseDirection(dir as Direction))) {
//     curr = { x: loc[0], y: loc[1] };
//     path.push(curr);
//     break;
//   }
// }
curr = { x: path[0].x, y: path[0].y + 1 };
path.push(curr);

let expanded = true;
while (expanded) {
  expanded = false;
  for (const [dir, loc] of Object.entries(
    getDirectionNeighboursLocs(curr.x, curr.y, map)
  )) {
    const char = map[loc[1]][loc[0]];
    if (
      connects(map[curr.y][curr.x], dir as Direction) &&
      connects(char, reverseDirection(dir as Direction)) &&
      !path.some((p) => p.x === loc[0] && p.y === loc[1])
    ) {
      curr = { x: loc[0], y: loc[1] };
      path.push(curr);
      expanded = true;
      break;
    } else if (
      connects(map[curr.y][curr.x], dir as Direction) &&
      path[0].x === loc[0] &&
      path[0].y === loc[1]
    ) {
      console.log("done");
    }
  }
}

// for (let y = 0; y < map.length; y++) {
//   console.log(
//     map[y].join(""),
//     map[y]
//       .map((v, i) => (path.some((p) => p.x === i && p.y === y) ? "X" : v))
//       .join("")
//   );
// }

console.log(path.length / 2);
