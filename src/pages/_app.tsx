import '../../styles/globals.scss'
import Layout from '../common/components/layout/Layout'
import type { AppProps } from 'next/app'

function MyApp({ Component, pageProps }: AppProps) {
  return (
  <Layout>
    <Component {...pageProps} />
  </Layout>
  )
}
export default MyApp
