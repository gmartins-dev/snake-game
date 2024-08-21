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

  return (
    <>
      <h1 className='title'>Snake Game 🐍</h1>
      <div className='score'>
        Score: {gameState ? gameState.score : 'Loading...'}
      </div>
      <div className='game-board'>
        {gameState ? (
          <>
            <Snake position={gameState.snake} />
            <Apple position={gameState.apple} />
          </>
        ) : (
          <div className='loading-placeholder'>Loading...</div>
        )}
      </div>
    </>
  );
};

export default GameBoard;
