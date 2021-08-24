/** @jsx jsx */
import { jsx, css } from '@emotion/react'
import React, { ReactElement, useEffect, useState } from 'react';
import axios from 'axios';

const imageSize = css({
  width: '30vw',
  height: '30vw',
})

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
      {randomImages.map(({ download_url, id, author }) => (
        <div key={id}>
          <img css={imageSize} src={download_url} alt={author} />
        </div>
      ))}
    </div>
  );
};

export default TodaysDeal;
