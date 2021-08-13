import '../../styles/globals.scss'
import type { AppProps } from 'next/app'
import Layout from '../common/components/layout/Layout'
import React, { ReactElement } from 'react'

function MyApp({ Component, pageProps }: AppProps): ReactElement {
  return (
      <Layout>
        <Component {...pageProps} />
      </Layout>
  )
}

export default MyApp
