import { cloneDeep, sum } from "lodash";
import { asList } from "../utils/inputReader";

const input = asList("input");

interface Metal {
  x: number;
  m: number;
  a: number;
  s: number;
}

type Rule =
  | {
      check: (val: Metal) => boolean;
      property: keyof Metal;
      num: number;
      condition: "<" | ">";
      target: string;
    }
  | {
      condition: "free";
      target: string;
    };

interface Workflow {
  name: string;
  rules: Rule[];
}

const workflows: Record<string, Workflow> = {};
const metals: Metal[] = [];

let workflowsInput = true;
for (const line of input) {
  if (line === "") {
    workflowsInput = false;
  } else if (workflowsInput) {
    const [name, rules] = line.slice(0, line.length - 1).split("{");
    const splitRules = rules.split(",");
    workflows[name] = {
      name,
      rules: splitRules.map((rawRule): Rule => {
        if (rawRule.includes(":")) {
          const [rawCheck, target] = rawRule.split(":");
          const property = rawCheck[0] as keyof Metal;
          const num = parseInt(rawCheck.slice(2));
          const check =
            rawCheck[1] === ">"
              ? (m: Metal) => m[property] > num
              : (m: Metal) => m[property] < num;

          return {
            target,
            property,
            num,
            condition: rawCheck[1] as "<" | ">",
            check,
          };
        } else {
          return { condition: "free", target: rawRule };
        }
      }),
    };
  } else {
    const metal = JSON.parse(
      line
        .replace("x", '"x"')
        .replace("m", '"m"')
        .replace("a", '"a"')
        .replace("s", '"s"')
        .replaceAll("=", ":")
    );
    metals.push(metal);
  }
}

//Simplify workflows step

for (const w of Object.values(workflows)) {
  while (
    w.rules.length > 1 &&
    w.rules.at(-1)!.target === w.rules.at(-2)!.target
  ) {
    const last = w.rules.pop()!;
    w.rules[w.rules.length - 1] = last;
  }
}

const accepted: Metal[] = [];

for (const m of metals) {
  let currWorkFlow = "in";
  while (currWorkFlow !== "A" && currWorkFlow !== "R") {
    const w = workflows[currWorkFlow];
    for (const rule of w.rules) {
      if (rule.condition === "free" || rule.check(m)) {
        currWorkFlow = rule.target;
        break;
      }
    }
  }

  if (currWorkFlow === "A") {
    accepted.push(m);
  }
}

console.log(sum(accepted.map((m) => m.x + m.m + m.a + m.s)));

interface MetalInterval {
  xStart: number;
  xEnd: number;
  mStart: number;
  mEnd: number;
  aStart: number;
  aEnd: number;
  sStart: number;
  sEnd: number;
}

function recursiveGetConfigurations(
  interval: MetalInterval,
  currWorkFlow: string
): number {
  if (currWorkFlow === "R") {
    return 0;
  } else if (currWorkFlow === "A") {
    return (
      (interval.xEnd - interval.xStart + 1) *
      (interval.mEnd - interval.mStart + 1) *
      (interval.aEnd - interval.aStart + 1) *
      (interval.sEnd - interval.sStart + 1)
    );
  }

  const w = workflows[currWorkFlow];

  let result = 0;
  const currentInterval = cloneDeep(interval);

  for (const r of w.rules) {
    if (r.condition === "free") {
      result += recursiveGetConfigurations(currentInterval, r.target);
    } else if (r.condition === "<") {
      const propertyStartKey = (r.property + "Start") as keyof MetalInterval;
      if (currentInterval[propertyStartKey] < r.num) {
        const propertyEndKey = (r.property + "End") as keyof MetalInterval;
        const recInterval = { ...currentInterval };
        recInterval[propertyEndKey] = r.num - 1;
        result += recursiveGetConfigurations(recInterval, r.target);

        currentInterval[propertyStartKey] = r.num;
      }
    } else {
      const propertyEndKey = (r.property + "End") as keyof MetalInterval;
      if (currentInterval[propertyEndKey] > r.num) {
        const propertyStartKey = (r.property + "Start") as keyof MetalInterval;
        const recInterval = { ...currentInterval };
        recInterval[propertyStartKey] = r.num + 1;
        result += recursiveGetConfigurations(recInterval, r.target);

        currentInterval[propertyEndKey] = r.num;
      }
    }
  }

  return result;
}

console.log(
  recursiveGetConfigurations(
    {
      xStart: 1,
      xEnd: 4000,
      mStart: 1,
      mEnd: 4000,
      aStart: 1,
      aEnd: 4000,
      sStart: 1,
      sEnd: 4000,
    },
    "in"
  )
);
