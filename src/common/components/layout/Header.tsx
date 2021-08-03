import React, { ReactElement } from 'react'
import Image from 'next/image'

const Header = (): ReactElement => {

  const mainLogo = "/images/main-cat-logo.png"

  return (
    <div className="header">
      <Image src={mainLogo} alt="logo" width={30} height={30} layout="fixed" />
      <Image src={mainLogo} alt="logo" width={30} height={30} layout="intrinsic" />
      <p>
        &quot;meow&quot;
      </p>
    </div>
  )
}

export default Header
