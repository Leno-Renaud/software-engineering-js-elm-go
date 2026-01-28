import { CARD_TYPES, numberCard } from "./card.js";

export default class Deck {
  constructor() {
    this.cards = [];
    this.build();
    this.shuffle();
  }

  build() {
    // nombres : N exemplaires de N
    for (let v = 0; v <= 12; v++) {
      for (let i = 0; i < v; i++) {
        this.cards.push(numberCard(v));
      }
    }

    // actions (quantités approximatives raisonnables)
    for (let i = 0; i < 3; i++) this.cards.push({ type: CARD_TYPES.FREEZE });
    for (let i = 0; i < 3; i++) this.cards.push({ type: CARD_TYPES.FLIP_THREE });
    for (let i = 0; i < 3; i++) this.cards.push({ type: CARD_TYPES.SECOND_CHANCE });

    // bonus +2 à +10
    for (let v = 2; v <= 10; v++) {
      this.cards.push({ type: CARD_TYPES.BONUS, value: v });
    }

    // x2
    this.cards.push({ type: CARD_TYPES.MULTIPLIER });
  }

  shuffle() {
    for (let i = this.cards.length - 1; i > 0; i--) {
      const j = Math.floor(Math.random() * (i + 1));
      [this.cards[i], this.cards[j]] = [this.cards[j], this.cards[i]];
    }
  }

  draw() {
    return this.cards.pop();
  }
}
