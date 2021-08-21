import React, { ReactElement, useEffect, useState } from 'react';
import Image from 'next/image';
import axios from 'axios';

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
    await axios
      .get('https://picsum.photos/v2/list')
      .then((response) => setRandomImages(response.data));
  };

  useEffect(() => {
    fetchRandomImages();
  }, []);

  return (
    <div>
      <h4>images</h4>
      {randomImages.map(({ download_url, id, author }) => (
        <div key={id}>
          <img src={download_url} alt={author} />
        </div>
      ))}
    </div>
  );
};

export default TodaysDeal;
