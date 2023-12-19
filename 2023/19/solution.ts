import { sum } from "lodash";
import { asList } from "../utils/inputReader";

const input = asList("input");

interface Metal {
  x: number;
  m: number;
  a: number;
  s: number;
}

interface Rule {
  check: (val: Metal) => boolean;
  target: string;
}

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

          return { target, check };
        } else {
          return { check: () => true, target: rawRule };
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

const accepted: Metal[] = [];

for (const m of metals) {
  let currWorkFlow = "in";
  while (currWorkFlow !== "A" && currWorkFlow !== "R") {
    const w = workflows[currWorkFlow];
    for (const rule of w.rules) {
      if (rule.check(m)) {
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
