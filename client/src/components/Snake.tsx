import React from 'react';
import { Position } from '../types/types';

interface SnakeProps {
  position: Position;
}

const Snake: React.FC<SnakeProps> = ({ position }) => {
  return (
    <div
      className='snake'
      style={{ top: position.y * 20, left: position.x * 20 }}
    />
  );
};

export default Snake;
