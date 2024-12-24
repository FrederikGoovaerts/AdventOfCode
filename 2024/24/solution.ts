import { asIs } from "../utils/inputReader";

const [input, fixedInput, maxIn, maxZ, solveOp] = [
  asIs("input"),
  asIs("input.fixed"),
  44,
  45,
  (a: number, b: number) => a + b,
];
// const [input, fixedInput, maxIn, maxZ, solveOp] = [
//   asIs("ex3"),
//   asIs("ex3.fixed"),
//   5,
//   5,
//   (a: number, b: number) => a & b,
// ];

const [startWires, ports] = input.split("\n\n").map((s) => s.split("\n"));
const [_, fixedPorts] = fixedInput.split("\n\n").map((s) => s.split("\n"));

const p1ValueMap: Map<string, 0 | 1> = new Map();

for (const wire of startWires) {
  const [label, val] = wire.split(": ");
  p1ValueMap.set(label, val === "1" ? 1 : 0);
}

type Op = "AND" | "XOR" | "OR";

interface Port {
  in1: string;
  in2: string;
  op: Op;
  out: string;
}

const portSet: Set<Port> = new Set();

for (const portLine of ports) {
  const [input, output] = portLine.split(" -> ");
  const [in1, op, in2] = input.split(" ");
  const port: Port = { in1, in2, op: op as Op, out: output };
  portSet.add(port);
}

const fixedPortSet: Set<Port> = new Set();

for (const portLine of fixedPorts) {
  const [input, output] = portLine.split(" -> ");
  const [in1, op, in2] = input.split(" ");
  const port: Port = { in1, in2, op: op as Op, out: output };
  fixedPortSet.add(port);
}

function calculateOut(port: Port, valueMap: Map<string, 0 | 1>) {
  const in1Val = valueMap.get(port.in1)!;
  const in2Val = valueMap.get(port.in2)!;

  switch (port.op) {
    case "AND":
      valueMap.set(port.out, in1Val === 1 && in2Val === 1 ? 1 : 0);
      break;
    case "OR":
      valueMap.set(port.out, in1Val === 1 || in2Val === 1 ? 1 : 0);
      break;
    case "XOR":
      valueMap.set(port.out, in1Val !== in2Val ? 1 : 0);
      break;
  }
}

function getResult(
  initialValues: Map<string, 0 | 1>,
  portSet: Set<Port>
): number {
  const valueMap = new Map(initialValues);

  let portsToHandle = new Set(portSet);

  while (portsToHandle.size > 0) {
    const prevSet = portsToHandle;
    portsToHandle = new Set();

    for (const port of prevSet) {
      if (valueMap.has(port.in1) && valueMap.has(port.in2)) {
        calculateOut(port, valueMap);
      } else {
        portsToHandle.add(port);
      }
    }
  }

  let z = 0;

  for (let i = 0; i <= maxZ; i++) {
    const zi = valueMap.get(`z${i.toString().padStart(2, "0")}`);
    if (zi === 1) {
      z += Math.pow(2, i);
    }
  }

  return z;
}

const p1Result = getResult(p1ValueMap, portSet);
console.log(p1Result);

/**
 * Part 2 solved by checking and fixing input. See "input.fixed" for working result.
 * Steps taken:
 * - All statements with "-> Zn" were sorted
 * - All statements with "Xn XOR Yn" and "Xn AND Yn" were sorted and given a clear output-label
 * - All other statements using the new labels were sorted, and any irregular statements were gathered
 * - By swapping exactly 8 outputs, listed below, all statements now followed a visible pattern
 *
 * Swapped ports:
 * - 11XOR (qjj) <-> 11AND (cbj)
 * - z35 <-> cfk
 * - z18 <-> dmn
 * - gmt <-> z07
 *
 * These are listed in the console.log below so the "solution" actually outputs the expected "answer"
 */
console.log("cbj,cfk,dmn,gmt,qjj,z07,z18,z35");

// Extra logic verifying the fixed input on the given operation with randomly generated test cases
for (let i = 0; i < 2_000; i++) {
  const x = Math.floor(Math.random() * (Math.pow(2, maxIn + 1) - 1));
  const y = Math.floor(Math.random() * (Math.pow(2, maxIn + 1) - 1));
  const expectedResult = solveOp(x, y);

  const valueMap: Map<string, 0 | 1> = new Map();

  for (let j = 0; j <= maxIn; j++) {
    let xj = ((x >> j) & 0b1) as 1 | 0;
    let yj = ((y >> j) & 0b1) as 1 | 0;

    valueMap.set(`x${j.toString().padStart(2, "0")}`, xj);
    valueMap.set(`y${j.toString().padStart(2, "0")}`, yj);
  }

  const logicResult = getResult(valueMap, fixedPortSet);

  const correctMask = expectedResult ^ logicResult;

  if (correctMask !== 0) {
    throw new Error(
      `Op on ${x} and ${y} does not produce expected result (${expectedResult})`
    );
  }
}
