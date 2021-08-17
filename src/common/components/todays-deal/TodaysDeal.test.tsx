import React from 'react';
import { render, screen, waitFor } from '@testing-library/react';
import '@testing-library/jest-dom';
import TodaysDeal from './TodaysDeal';
import { rest } from 'msw';
import { setupServer } from 'msw/node';
import { fetchMockRandomImages } from '../../../mocks/randomImages';

const server = setupServer(
  rest.get('https://picsum.photos/v2/list', (req, res, ctx) => {
    return res(ctx.json({
      data: [
        {
          id: 1,
          author: 'tester1',
          download_url: 'tester_url1',

        },
        {
          id: 2,
          author: 'tester2',
          download_url: 'tester_url2',
        },
      ],
    }))
  })
)

// beforeAll(() => server.listen())
// afterAll(() => server.close())

test('testing a test run', () => {
  render(<TodaysDeal />)
  screen.debug(screen.getByText('images'))
});

test('handles fetching images', async () => {
  server.use(
    rest.get('https://picsum.photos/v2/list', (req, res, ctx) => {
      return res(ctx.status(200))
    })
  )
})
