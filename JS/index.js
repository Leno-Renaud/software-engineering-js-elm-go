import Player from "./player.js";
import Game from "./game.js";
import { ask, closePrompt } from "./prompt.js";

async function main() {
  const n = parseInt(await ask("Nombre de joueurs : "));

  const players = [];

  for (let i = 0; i < n; i++) {
    const name = await ask(`Nom joueur ${i + 1}: `);
    players.push(new Player(name));
  }

  const game = new Game(players);

  await game.start();

  closePrompt();
}

main();
