import React, { ReactElement } from 'react'
import Header from './Header'
import Footer from './Footer'
import Navbar from './Navbar'

interface ParentCompProps {
  children: ReactElement;
}

const Layout = ({children}: ParentCompProps): ReactElement => {
  return (
    <div>
      <Header />
      <Navbar />
        {children}
      <Footer />
    </div>

  )
}

export default Layout
