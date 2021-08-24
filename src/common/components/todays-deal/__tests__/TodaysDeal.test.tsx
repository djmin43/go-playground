import * as React from 'react'
import { render, screen, waitFor } from '@testing-library/react'
import '@testing-library/jest-dom'
import { setupServer } from 'msw/node'
import { handlers } from '../../../../mocks/server-handler'
import TodaysDeal from '../TodaysDeal'
import { rest } from 'msw'

const server = setupServer(...handlers)

// beforeAll(() => {
//   server.listen()
// })

// afterAll(() => {
//   server.close()
// })

it('testing a server', async () => {
  render(<TodaysDeal />)
  expect(await screen.findByText('images')).toBeInTheDocument()
  await waitFor(() => expect(screen.findByText('images')).toBeInTheDocument())
  screen.debug()
})