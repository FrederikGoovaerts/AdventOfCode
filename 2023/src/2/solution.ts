import { sum } from "lodash";
import { asList } from "../utils/inputReader";

const input = asList();

const part1Max: Record<string, number> = {
  blue: 14,
  red: 12,
  green: 13,
};

const part1PossibleGames: number[] = [];
const part2Powers: number[] = [];

for (const line of input) {
  const splitMain = line.split(": ");
  const gameId = parseInt(splitMain[0].split(" ")[1]);
  const grabs = splitMain[1].split("; ");

  let part1Possible = true;
  const part2Min: Record<string, number> = {
    blue: 0,
    red: 0,
    green: 0,
  };

  for (const grab of grabs) {
    const cubeList = grab.split(", ");
    for (const cube of cubeList) {
      const color = cube.split(" ")[1];
      const amount = parseInt(cube.split(" ")[0]);
      if (part1Max[color] < amount) {
        part1Possible = false;
      }
      part2Min[color] = Math.max(part2Min[color], amount);
    }
  }

  if (part1Possible) {
    part1PossibleGames.push(gameId);
  }

  part2Powers.push(part2Min.red * part2Min.blue * part2Min.green);
}

console.log(sum(part1PossibleGames));
console.log(sum(part2Powers));
