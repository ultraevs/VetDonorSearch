import React from 'react'
import MainLayout from '../../components/MainLayout/MainLayout'
import Chat from '../../components/Chat/Chat'
import NeedToClinic from '../../components/NeedToClinic/NeedToClinic'
import Maps from '../../components/Maps/Maps'

const Where = () => {
  return (
    <MainLayout>
      <NeedToClinic />
      <Maps />
    </MainLayout>
  )
}

export default Where