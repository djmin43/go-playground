import { rest } from 'msw'
import { setupServer } from 'msw/node'
import { handlers } from '../src/mocks/server-handler'

const server = setupServer(...handlers)
export { server, rest }