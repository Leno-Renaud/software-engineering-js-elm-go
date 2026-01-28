import fs from "fs";

export default class Logger {
  constructor() {
    const date = new Date().toISOString().replace(/[:.]/g, "-");
    this.path = `logs/game-${date}.json`;
    this.data = { rounds: [] };
  }

  startRound(roundNumber, players) {
    this.current = {
      round: roundNumber,
      players: players.map(p => p.name),
      events: [],
      scores: {}
    };
  }

  log(event) {
    this.current.events.push(event);
  }

  endRound(players) {
    for (const p of players) {
      this.current.scores[p.name] = p.totalScore;
    }
    this.data.rounds.push(this.current);
    fs.mkdirSync("logs", { recursive: true });
    fs.writeFileSync(this.path, JSON.stringify(this.data, null, 2));
  }
}
