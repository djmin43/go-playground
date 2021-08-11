import '../../styles/globals.scss'
import { Provider } from 'react-redux'
import type { AppProps } from 'next/app'
import store from '../redux/store/index'
import Layout from '../common/components/layout/Layout'

function MyApp({ Component, pageProps }: AppProps) {
  return (
    <Provider store={store}>
      <Layout>
        <Component {...pageProps} />
      </Layout>
    </Provider>
  )
}

export default MyApp
