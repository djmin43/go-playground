import React, { useState } from 'react'

const Calculate = () => {

  const [currentNumber, setCurrentNumber] = useState<number>(0)

  const addNumber = () => {
    setCurrentNumber(currentNumber + 1)
  }

  return (
    <div>
      <h1>{currentNumber}</h1>
      <button onClick={() => addNumber()} >
        Add
      </button>
    </div>
  )
}

export default Calculate
