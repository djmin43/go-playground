import { createSlice } from '@reduxjs/toolkit'

const initialState = [
  { id: 0, title: 'Star Wars', genre: 'SF' },
  { id: 1, title: 'Departed', genre: 'Action' }
]

const postSlice = createSlice({
  name: 'value',
  initialState,
  reducers: {
    postNewMovie(state, action) {
      state.push(action.payload)
    }
  }
})

export const { postNewMovie } = postSlice.actions

export default postSlice.reducer