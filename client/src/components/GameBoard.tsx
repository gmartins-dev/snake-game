import React, { useEffect, useState } from 'react';
import { GameState } from '../types/types';
import Apple from './Apple';
import Snake from './Snake';

const GameBoard: React.FC = () => {
  const [gameState, setGameState] = useState<GameState | null>(null);

  useEffect(() => {
    const startGame = async () => {
      try {
        const response = await fetch('/api/start', {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json',
          },
        });
        const data: GameState = await response.json();
        setGameState(data);
      } catch (error) {
        console.error('Error starting game:', error);
      }
    };

    startGame();
  }, []);

  if (!gameState) {
    return <div>Loading...</div>;
  }

  return (
    <div className='game-board'>
      <Snake position={gameState.snake} />
      <Apple position={gameState.apple} />
      <div className='score'>Score: {gameState.score}</div>
    </div>
  );
};

export default GameBoard;
