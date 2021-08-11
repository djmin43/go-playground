import { createSlice } from '@reduxjs/toolkit'

const initialState = [
  { id: 0, title: 'Star Wars', genre: 'SF' },
  { id: 1, title: 'Departed', genre: 'Action' }
]

const postSlice = createSlice({
  name: 'value',
  initialState,
  reducers: {}
})

export default postSlice.reducer