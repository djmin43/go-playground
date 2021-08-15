import React, { ReactElement } from 'react';
import Head from 'next/head';
import LandingPage from '../common/components/landing/LandingPage';

export default function Home(): ReactElement {
  return (
    <div className="App">
      <Head>
        <title>Makgoli</title>
        <link rel="icon" href="/favicon.ico" />
      </Head>
      <LandingPage />
    </div>
  );
}
