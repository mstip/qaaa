import { RunEl } from "./elements.js";

function run(event) {
  event.preventDefault();
  // TODO: add
  console.log("run");
}

export default function Run() {
  RunEl.onclick = run;
}
