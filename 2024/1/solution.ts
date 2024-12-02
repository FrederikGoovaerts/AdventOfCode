import { asList } from "../utils/inputReader";

const input = asList("input");

const listOne: number[] = [];
const listTwo: number[] = [];

for (const line of input) {
  const [one, two] = line.split("   ");
  listOne.push(Number(one));
  listTwo.push(Number(two));
}

listOne.sort();
listTwo.sort();

let diff = 0;

listOne.forEach((val, i) => {
  diff += Math.abs(val - listTwo[i]);
});

console.log(diff);

const amount: Record<number, number> = {};

for (const el of listTwo) {
  const curr = amount[el];
  amount[el] = (curr ?? 0) + 1;
}

let sim = 0;

listOne.forEach((val) => {
  sim += val * (amount[val] ?? 0);
});

console.log(sim);
