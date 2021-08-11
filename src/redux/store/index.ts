import { configureStore } from '@reduxjs/toolkit'
import postReducer from '../reducers/postSlice'

const store = configureStore({
  reducer: {
    value: postReducer
  }
})

export type RootState = ReturnType<typeof store.getState>

export default store
