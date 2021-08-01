import Calculator from '../components/Caclulator'

import { render } from '@testing-library/react';

test('fake test', () => {
  expect(true).toBeTruthy()
})

test('renders calculator', () => {
  render(<Calculator />);
});

export {}