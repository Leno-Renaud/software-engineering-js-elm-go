export const CARD_TYPES = {
  NUMBER: "number",
  FREEZE: "freeze",
  FLIP_THREE: "flip_three",
  SECOND_CHANCE: "second_chance",
  BONUS: "bonus",
  MULTIPLIER: "multiplier"
};

export const numberCard = (value) => ({
  type: CARD_TYPES.NUMBER,
  value
});
