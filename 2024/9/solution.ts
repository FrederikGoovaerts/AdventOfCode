import { asIs } from "../utils/inputReader";

const input = asIs("input").split("").map(Number);

let disk: (number | undefined)[] = [];

function getChecksum(): number {
  let checksum = 0;

  disk.forEach((val, index) => {
    if (val !== undefined) {
      checksum += val * index;
    }
  });

  return checksum;
}

let inputPointer = 0;
let diskPointer = 0;
let fileId = 0;

while (inputPointer < input.length) {
  const inputVal = input[inputPointer];

  if (inputPointer % 2 === 0) {
    const currFileId = fileId++;

    for (let i = 0; i < inputVal; i++) {
      disk[diskPointer + i] = currFileId;
    }
  } else {
    for (let i = 0; i < inputVal; i++) {
      disk[diskPointer + i] = undefined;
    }
  }

  diskPointer += inputVal;

  inputPointer++;
}

const backup = [...disk];

let front = disk.indexOf(undefined);
let back = diskPointer - 1;

while (front < back) {
  disk[front] = disk[back];
  disk[back] = undefined;

  while (disk[front] !== undefined) {
    front++;
  }

  while (disk[back] === undefined) {
    back--;
  }
}

console.log(getChecksum());

// Part 2

disk = [...backup];
back = diskPointer - 1;

function isEmpty(start: number, size: number): boolean {
  for (let i = start; i < start + size; i++) {
    if (disk[i] !== undefined) {
      return false;
    }
  }
  return true;
}

while (back > 0) {
  if (disk[back] === undefined) {
    back--;
    continue;
  }

  let val = disk[back];

  let backBegin = back;

  while (disk[backBegin - 1] === val) {
    backBegin--;
  }

  const size = back - backBegin + 1;

  let gapPointer = disk.indexOf(undefined);

  while (
    !isEmpty(gapPointer, size) &&
    gapPointer < backBegin &&
    gapPointer !== -1
  ) {
    gapPointer = disk.indexOf(undefined, gapPointer + 1);
  }

  if (gapPointer < backBegin && gapPointer !== -1) {
    for (let i = gapPointer; i < gapPointer + size; i++) {
      disk[i] = val;
    }

    for (let i = backBegin; i <= back; i++) {
      disk[i] = undefined;
    }
  }

  back = backBegin - 1;
}

console.log(getChecksum());
