import React from 'react'
import Header from '../Header/Header'
import Footer from '../Footer/Footer'
import Chat from '../Chat/Chat'

const MainLayout = ({children}) => {
  return (
    <div>
        <Header />
        {children}
        <Chat />
        <Footer />
    </div>
  )
}

export default MainLayout