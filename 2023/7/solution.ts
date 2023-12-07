import { asList } from "../utils/inputReader";
const input = asList("input");

interface HandValue {
  originalHand: string;
  absoluteValue: number;
  bet: number;
}

const handValues: HandValue[] = [];

const cardValueMap: Record<string, number> = {
  A: 14,
  K: 13,
  Q: 12,
  J: 11,
  T: 10,
};

function getTypeValue(cards: string): number {
  const sortedCards = cards.split("").sort().join("");
  if (/(\w)\1{4}/.test(sortedCards)) {
    return 6;
  } else if (/(\w)\1{3}/.test(sortedCards)) {
    return 5;
  } else if (
    /(\w)\1{2}.*(\w)\2{1}/.test(sortedCards) ||
    /(\w)\1{1}.*(\w)\2{2}/.test(sortedCards)
  ) {
    return 4;
  } else if (/(\w)\1{2}/.test(sortedCards)) {
    return 3;
  } else if (/(\w)\1{1}.*(\w)\2{1}/.test(sortedCards)) {
    return 2;
  } else if (/(\w)\1{1}/.test(sortedCards)) {
    return 1;
  }

  return 0;
}

function getAbsoluteValue(cards: string): number {
  let total = 0;

  total += getTypeValue(cards) * Math.pow(10, 10);

  for (let i = 0; i < 5; i++) {
    const cardVal = cardValueMap[cards[i]] ?? parseInt(cards[i]);

    total += cardVal * Math.pow(10, 8 - i * 2);
  }
  return total;
}

for (const line of input) {
  const [handPart, betPart] = line.split(" ");
  handValues.push({
    originalHand: handPart,
    absoluteValue: getAbsoluteValue(handPart),
    bet: parseInt(betPart),
  });
}

// Sorted from lowest rank (1) to highest (N)
const sortedValues = handValues.sort(
  (a, b) => a.absoluteValue - b.absoluteValue
);

let winnings = 0;

sortedValues.forEach((v, i) => {
  winnings += v.bet * (i + 1);
});

console.log(winnings);
