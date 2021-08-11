import { configureStore } from '@reduxjs/toolkit'

import postReducer from './reducers/postSlice'

export default configureStore({
  reducer: {
    value: postReducer
  }
})