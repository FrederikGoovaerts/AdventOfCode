import { asNumberList } from "../utils/inputReader";

const input = asNumberList("input");

function getNext(sec: number): number {
  let curr = BigInt(sec);

  const mult1 = curr * 64n;
  curr = (curr ^ mult1) % 16777216n;

  const div = curr / 32n;
  curr = (curr ^ div) % 16777216n;

  const mult2 = curr * 2048n;
  curr = (curr ^ mult2) % 16777216n;

  return Number(curr);
}

let sum = 0;

for (const init of input) {
  let sec = init;
  for (let i = 0; i < 2000; i++) {
    sec = getNext(sec);
  }
  sum += sec;
}
console.log(sum);

const bananaMap: Map<string, number[]> = new Map();

input.forEach((init, inputIndex) => {
  let sec = init;
  let changes: number[] = [];

  for (let i = 0; i < 2000; i++) {
    const prevVal = sec;

    sec = getNext(sec);

    if (changes.length === 4) {
      changes = [
        changes[1],
        changes[2],
        changes[3],
        (sec % 10) - (prevVal % 10),
      ];

      if (changes[3] >= 0) {
        const changeKey = changes.join();

        const matchingArray = bananaMap.get(changeKey);
        if (matchingArray) {
          if (matchingArray[inputIndex] === undefined) {
            matchingArray[inputIndex] = sec % 10;
          }
        } else {
          const arr: number[] = [];
          arr[inputIndex] = sec % 10;
          bananaMap.set(changeKey, arr);
        }
      }
    } else {
      changes.push((sec % 10) - (prevVal % 10));
    }
  }
});

let most = 0;

for (const [_, bananas] of bananaMap) {
  let sum = 0;
  for (const b of bananas) {
    if (b !== undefined) {
      sum += b;
    }
  }
  if (sum > most) {
    most = sum;
  }
}
console.log(most);
