import readline from "readline";

const rl = readline.createInterface({
  input: process.stdin,
  output: process.stdout
});

export const ask = (q) =>
  new Promise(resolve => rl.question(q, resolve));

export const closePrompt = () => rl.close();
