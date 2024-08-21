import React from 'react';
import './App.css';
import GameBoard from './components/GameBoard';

const App: React.FC = () => {
  return (
    <div className='App'>
      <h1>Snake Game</h1>
      <GameBoard />
    </div>
  );
};

export default App;
