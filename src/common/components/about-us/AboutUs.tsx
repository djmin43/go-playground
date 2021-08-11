import React, { ReactElement } from 'react'
import { useSelector } from 'react-redux'
import { RootState } from '../../../redux/store'


const AboutUs = (): ReactElement => {
  const reduxValue = useSelector((state: RootState) => state.value)

  return (
    <div>
      asdf
      {reduxValue.map(({id, title, genre}) => 
        <div key={id}>
          <p>{title}</p>
          <p>{genre}</p>
        </div>
      )}
    </div>
  )
}

export default AboutUs
