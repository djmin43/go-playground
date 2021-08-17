import React, { ReactElement, useEffect, useState } from 'react';
import Image from 'next/image';
import axios from 'axios';
import 'regenerator-runtime/runtime';

export type RandomImage = {
  id: string;
  author: string;
  download_url: string;
  url: string;
  height: number;
  width: number;
};

const TodaysDeal = (): ReactElement => {
  const [randomImages, setRandomImages] = useState<RandomImage[]>([]);

  const fetchRandomImages = async () => {
    const randomImageArray: RandomImage[] = await axios
      .get('https://picsum.photos/v2/list')
      .then((response) => response.data);
    setRandomImages(randomImageArray);
  };

  useEffect(() => {
    fetchRandomImages();
  }, []);

  return (
    <div>
      <h4>images</h4>
      {randomImages.map(({ download_url, id, author }) => (
        <div key={id}>
          <Image src={download_url} alt={author} height={300} width={300} />
        </div>
      ))}
    </div>
  );
};

export default TodaysDeal;
