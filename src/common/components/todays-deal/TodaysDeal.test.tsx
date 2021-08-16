import React from 'react';
import { render, screen } from '@testing-library/react';
import '@testing-library/jest-dom';
import TodaysDeal from './TodaysDeal';
import { rest } from 'msw';
import { setupServer } from 'msw/node';
import { fetchMockRandomImages } from '../../../mocks/randomImages';

const server = setupServer(...fetchMockRandomImages)

beforeAll(() => server.listen())

describe('testing a test run', () => {
  it('testing a mock server', async () => {
    server.use(
      rest.get('https://picsum.photos/v2/list', (req, res, ctx) => {
        return res(ctx.status(200))
      })
    )
  })
  render(<TodaysDeal />);
});
