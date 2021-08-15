import * as React from 'react';
import '@testing-library/jest-dom';
import { render } from '@testing-library/react';
import TodaysDeal from './TodaysDeal';

test('sample test', () => {
  const { container } = render(<TodaysDeal />);
  expect(container.firstChild).toBeVisible();
});
