import React, { useState } from 'react'

const Caclulator = () => {
  
  const [number, setNumber] = useState<number>(0)

  const addNumber = () => {
    setNumber(number + 1)
  }

  const subtractNumber = () => {
    setNumber(number - 1)
  }

  return (
    <div>
      {number}
      <button onClick={ () => addNumber() }>Add</button>
      <button onClick={ () => subtractNumber() }>Subtract</button>
    </div>
  )
}

export default Caclulator
