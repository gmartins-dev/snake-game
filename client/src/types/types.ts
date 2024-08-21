export interface Position {
  x: number;
  y: number;
}

export interface GameState {
  snake: Position;
  apple: Position;
  score: number;
}
