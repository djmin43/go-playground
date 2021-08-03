import React, { ReactElement } from 'react'
import Link from 'next/link'

const Navbar = (): ReactElement => {
  return (
    <div className="navbar">
      <Link href="/todays-deal">
        <a>Today&apos;s deals</a>
      </Link>
      <Link href="/customer-service">
        <a>Customer Service</a>
      </Link>      
      <Link href="/about-us">
        <a>About Us</a>
      </Link>    
    </div>
  )
}

export default Navbar
