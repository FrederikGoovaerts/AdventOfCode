import { sum } from "lodash";
import { asNumberList } from "../utils/inputReader";

const input = asNumberList();

console.log(sum(input));
