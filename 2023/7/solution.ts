import { asList } from "../utils/inputReader";
const input = asList("input");
const PART_2_MODE = true;

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
  Z: 1,
};

function sortCards(cards: string): string[] {
  const sortedCards = cards.split("").sort().join("");
  if (!sortedCards.includes("Z") || sortedCards === "ZZZZZ") {
    return [sortedCards];
  } else {
    const result: string[] = [];
    const normals = sortedCards.substring(0, sortedCards.indexOf("Z"));
    const jokers = sortedCards.substring(sortedCards.indexOf("Z"));
    for (let i = 1; i < normals.length; i++) {
      result.push(normals.slice(0, i) + jokers + normals.slice(i));
    }
    result.push(sortedCards);
    return result;
  }
}

function getTypeValue(cards: string): number {
  const mapType = (c: string) => {
    if (/(\w)(?:\1|Z){4}/.test(c)) {
      return 6;
    } else if (/(\w)(?:\1|Z){3}/.test(c)) {
      return 5;
    } else if (
      /(\w)(?:\1|Z){2}.*(\w)(?:\2|Z){1}/.test(c) ||
      /(\w)(?:\1|Z){1}.*(\w)(?:\2|Z){2}/.test(c)
    ) {
      return 4;
    } else if (/(\w)(?:\1|Z){2}/.test(c)) {
      return 3;
    } else if (/(\w)(?:\1|Z){1}.*(\w)(?:\2|Z){1}/.test(c)) {
      return 2;
    } else if (/(\w)(?:\1|Z){1}/.test(c)) {
      return 1;
    }

    return 0;
  };

  const sortedCards = sortCards(cards);
  const mappedValues = sortedCards.map((c) => mapType(c));

  return Math.max(...mappedValues);
}

function getAbsoluteValue(cards: string): number {
  const usedCards = PART_2_MODE ? cards.replaceAll("J", "Z") : cards;

  let total = 0;

  total += getTypeValue(usedCards) * Math.pow(10, 10);

  for (let i = 0; i < 5; i++) {
    const cardVal = cardValueMap[usedCards[i]] ?? parseInt(usedCards[i]);

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
