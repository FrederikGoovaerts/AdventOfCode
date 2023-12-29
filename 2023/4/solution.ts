import { sum } from "lodash";
import { asList } from "../utils/inputReader";
const input = asList("input");

const worth: number[] = [];

const cardAmount: number[] = [];

for (let i = 0; i < input.length; i++) {
  cardAmount[i] = 1;
}

input.forEach((line, cardNumber) => {
  const [winningRaw, cardRaw] = line.split(": ")[1].split(" | ");

  const winning = winningRaw.split(" ").filter((s) => s !== "");
  const card = cardRaw.split(" ").filter((s) => s !== "");
  let correct = 0;
  for (const cardNum of card) {
    if (winning.includes(cardNum)) {
      correct++;
    }
  }

  if (correct > 0) {
    worth.push(Math.pow(2, correct - 1));
  }

  for (let i = cardNumber + 1; i < cardNumber + 1 + correct; i++) {
    if (i < cardAmount.length) {
      cardAmount[i] += cardAmount[cardNumber];
    }
  }
});

console.log(sum(worth));
console.log(sum(cardAmount));
