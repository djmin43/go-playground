import React, { ReactElement } from 'react'
import { useSelector } from 'react-redux'
import { RootState } from '../../../redux/store'
import { useDispatch } from 'react-redux'
import { nanoid } from '@reduxjs/toolkit'
import { postNewMovie } from '../../../redux/reducers/postSlice'

const AboutUs = (): ReactElement => {
  const reduxValue = useSelector((state: RootState) => state.value)

  const dispatch = useDispatch()

  const handlePostNewMovie = () => {
    dispatch(
      postNewMovie({
        id: nanoid(),
        title: 'New Movie',
        genre: 'new genre',
      })
    )
  }

  return (
    <div>
      asdf
      {reduxValue.map(({id, title, genre}) => 
        <div key={id}>
          <p>{title}</p>
          <p>{genre}</p>
        </div>
      )}
      <button onClick={handlePostNewMovie}>add a new movie</button>
    </div>
  )
}

export default AboutUs
