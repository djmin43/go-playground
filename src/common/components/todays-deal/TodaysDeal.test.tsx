import React from 'react';
import { render, screen } from '@testing-library/react';
import '@testing-library/jest-dom';
import TodaysDeal from './TodaysDeal';

it('testing a test run', () => {
  render(<TodaysDeal />);
  const hello = screen.getByTestId('hello');
  expect(hello).toHaveTextContent('hello');
});
