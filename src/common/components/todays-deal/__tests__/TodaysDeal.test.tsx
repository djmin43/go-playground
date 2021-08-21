import * as React from 'react'
import { render, screen } from '@testing-library/react'
import TodaysDeal from '../TodaysDeal'

it('testing a server', async () => {
  render(<TodaysDeal />)
  screen.debug()
})