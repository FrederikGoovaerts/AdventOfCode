import { asList } from "../utils/inputReader";

const input = asList("input");

function isSafeInc(report: string): boolean {
  const splitReport = report.split(" ").map(Number);

  for (let i = 0; i < splitReport.length - 1; i++) {
    const step = splitReport[i + 1] - splitReport[i];
    if (step < 1 || step > 3) {
      return false;
    }
  }
  return true;
}

function isSafeDec(report: string): boolean {
  const splitReport = report.split(" ").map(Number);

  for (let i = 0; i < splitReport.length - 1; i++) {
    const step = splitReport[i] - splitReport[i + 1];
    if (step < 1 || step > 3) {
      return false;
    }
  }
  return true;
}

function isSafe(report: string, dampen: boolean): boolean {
  const isBaseSafe = isSafeInc(report) || isSafeDec(report);

  if (!isBaseSafe && dampen) {
    const splitReport = report.split(" ").map(Number);
    for (let i = 0; i < splitReport.length; i++) {
      const dampenedReport = [...splitReport];
      dampenedReport.splice(i, 1);
      const dampenedReportString = dampenedReport.join(" ");

      if (isSafeInc(dampenedReportString) || isSafeDec(dampenedReportString)) {
        return true;
      }
    }
  }

  return isBaseSafe;
}

let safeCount = 0;
let dampenedSafeCount = 0;

for (const line of input) {
  if (isSafe(line, false)) {
    safeCount++;
  }
  if (isSafe(line, true)) {
    dampenedSafeCount++;
  }
}

console.log(safeCount);
console.log(dampenedSafeCount);
