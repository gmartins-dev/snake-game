import React from 'react';
import { Position } from '../types/types';

interface AppleProps {
  position: Position;
}

const Apple: React.FC<AppleProps> = ({ position }) => {
  return (
    <div
      className='apple'
      style={{ top: position.y * 20, left: position.x * 20 }}
    />
  );
};

export default Apple;
