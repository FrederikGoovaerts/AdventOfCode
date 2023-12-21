import { cloneDeep, sum } from "lodash";
import { asList } from "../utils/inputReader";
import { Queue } from "../utils/queue";

const input = asList("input");

type module =
  | { name: string; type: "end" }
  | { name: string; type: "bc"; targets: string[] }
  | { name: string; type: "ff"; state: boolean; targets: string[] }
  | {
      name: string;
      type: "con";
      state: Record<string, boolean>;
      targets: string[];
    };

const moduleMap: Record<string, module> = {};

for (const line of input) {
  const [id, rawTargets] = line.split(" -> ");
  const targets = rawTargets.split(", ");
  if (id === "broadcaster") {
    moduleMap[id] = { name: id, type: "bc", targets };
  } else if (id.startsWith("%")) {
    moduleMap[id.slice(1)] = {
      name: id.slice(1),
      type: "ff",
      state: false,
      targets,
    };
  } else if (id.startsWith("&")) {
    moduleMap[id.slice(1)] = {
      name: id.slice(1),
      type: "con",
      state: {},
      targets,
    };
  }
}

// initialize state of Conjunctions and find unlisted end connections
for (const v of Object.values(moduleMap)) {
  if (v.type === "con") {
    for (const v2 of Object.values(moduleMap)) {
      if (v2.type !== "end" && v2.targets.includes(v.name)) {
        v.state[v2.name] = false;
      }
    }
  }
  if (v.type !== "end") {
    for (const target of v.targets) {
      if (moduleMap[target] === undefined) {
        moduleMap[target] = { type: "end", name: "target" };
      }
    }
  }
}

interface Pulse {
  origin: string;
  target: string;
  state: boolean;
}

const part1State = cloneDeep(moduleMap);

function pushButton(state: Record<string, module>): {
  highs: number;
  lows: number;
  rxPulses: boolean[];
} {
  let highs = 0;
  let lows = 0;
  const rxPulses: boolean[] = [];

  const pulses: Queue<Pulse> = new Queue();
  pulses.enqueue({ target: "broadcaster", state: false, origin: "button" });

  while (!pulses.isEmpty()) {
    const nextPulse = pulses.dequeue()!;

    if (nextPulse.state) {
      highs++;
    } else {
      lows++;
    }

    if (nextPulse.target === "rx") {
      rxPulses.push(nextPulse.state);
    }

    const pulseTarget = state[nextPulse.target];
    if (pulseTarget.type === "bc") {
      for (const t of pulseTarget.targets) {
        pulses.enqueue({
          origin: nextPulse.target,
          target: t,
          state: nextPulse.state,
        });
      }
    } else if (pulseTarget.type === "ff") {
      if (!nextPulse.state) {
        pulseTarget.state = !pulseTarget.state;
        for (const t of pulseTarget.targets) {
          pulses.enqueue({
            origin: nextPulse.target,
            target: t,
            state: pulseTarget.state,
          });
        }
      }
    } else if (pulseTarget.type === "con") {
      pulseTarget.state[nextPulse.origin] = nextPulse.state;
      const outPulseState = Object.values(pulseTarget.state).some(
        (v) => v === false
      )
        ? true
        : false;
      for (const t of pulseTarget.targets) {
        pulses.enqueue({
          origin: nextPulse.target,
          target: t,
          state: outPulseState,
        });
      }
    }
  }

  return { highs, lows, rxPulses };
}
const highs: number[] = [];
const lows: number[] = [];

for (let i = 0; i < 1000; i++) {
  const result = pushButton(part1State);
  highs.push(result.highs);
  lows.push(result.lows);
}

console.log(sum(highs) * sum(lows));

const part2State = cloneDeep(moduleMap);

for (let i = 1; ; i++) {
  const result = pushButton(part2State);
  if (result.rxPulses.includes(false)) {
    console.log(i, result.rxPulses);
    break;
  }
}
