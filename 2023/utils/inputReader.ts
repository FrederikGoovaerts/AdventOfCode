import * as fs from "fs";

export const asIs = (fileName: string = "input"): string =>
  fs.readFileSync(fileName, "utf8").trim();

export const asLineList = (fileName: string = "input"): string[] =>
  asIs(fileName).split(",");

export const asNumberLineList = (fileName: string = "input"): number[] =>
  asLineList(fileName).map(Number);

export const asList = (fileName: string = "input"): string[] =>
  asIs(fileName).split("\n");

export const asNumberList = (fileName: string = "input"): number[] =>
  asList(fileName).map(Number);

export const as2d = (fileName: string = "input"): string[][] =>
  asIs(fileName)
    .split("\n")
    .map((line) => line.split(""));

export const as2dNumbers = (fileName: string = "input"): number[][] =>
  asIs(fileName)
    .split("\n")
    .map((line) => line.split("").map((v) => parseInt(v)));
