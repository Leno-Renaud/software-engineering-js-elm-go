export default class Player {
  constructor(name) {
    this.name = name;
    this.totalScore = 0;
    this.resetRound();
  }

  resetRound() {
    this.numbers = [];
    this.bonuses = [];
    this.multiplier = false;
    this.secondChance = false;

    this.active = true;
    this.stayed = false;
    this.frozen = false;
  }

  addNumber(value) {
    this.numbers.push(value);
  }

  hasDuplicate(value) {
    return this.numbers.includes(value);
  }

  scoreRound() {
    if (!this.stayed || this.frozen) return 0;

    let sum = this.numbers.reduce((a, b) => a + b, 0);

    if (this.multiplier) sum *= 2;

    const bonus = this.bonuses.reduce((a, b) => a + b, 0);

    if (this.numbers.length === 7) sum += 15;

    return sum + bonus;
  }
}
