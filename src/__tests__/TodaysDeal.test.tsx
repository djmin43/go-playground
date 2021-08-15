import React from 'react';
import {
  render, fireEvent, waitFor, screen,
} from '@testing-library/react';
import '@testing-library/jest-dom';
import TodaysDeal from '../common/components/todays-deal/TodaysDeal';

test('empty test', () => {
  render(<TodaysDeal />);
  expect('deal-image').toBeVisible();
});
