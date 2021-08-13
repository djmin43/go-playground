import React, { ReactElement, useEffect, useState } from 'react'
import Image from 'next/image'
import axios from 'axios'

type RandomImage = {
  id: string;
  author: string;
  download_url: string;
  url: string;
  height: number;
  width: number;
}

const TodaysDeal = (): ReactElement => {
  const [randomImages, setRandomImages] = useState<RandomImage[]>([])

  const fetchRandomImages = async () => {
    try {
      const randomImageArray: RandomImage[] = await axios.get('https://picsum.photos/v2/list').then(response => response.data)
      console.log(randomImageArray)
      setRandomImages(randomImageArray)
    } catch {
      console.log(Error)
    }
  }

  useEffect(() => {
    fetchRandomImages()
  }, [])


  return (
    <div>
      {randomImages.map(({url, download_url, id, author, height, width}) => 
      <div key={id}>
        <Image src={download_url} alt={author} height={height/10} width={width/10}/>
      </div>
      )}
    </div>
  )
}

export default TodaysDeal
