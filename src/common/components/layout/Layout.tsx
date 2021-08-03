import React from 'react'
import Header from './Header'
import Footer from './Footer'
import Navbar from './Navbar'

interface ParentCompProps {
  children: React.ReactNode;
}

const Layout = ({children}: ParentCompProps) => {
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
