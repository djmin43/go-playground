import React, { ReactElement } from 'react'
import { StyleSheet, css } from 'aphrodite'

const AboutUs = (): ReactElement => {

  return (
    <div className={css(styles.aboutUs)}>
      about us page
    </div>
  )
}

const styles = StyleSheet.create({
  aboutUs: {
    height: '80vh'
  }
})  

export default AboutUs
