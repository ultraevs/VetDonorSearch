import React from 'react'
import Start from './pages/Start/Start'
import Auth from "./pages/Auth/Auth"
import Register from './pages/Register/Register'
import Profile from "./pages/Profile/Profile"
import Main from './pages/Main/Main'

import { Route, Routes } from 'react-router-dom'

const App = () => {
  return (
    <Routes>
      <Route path='/' element={ <Start />}/>
      <Route path='/auth' element={ <Auth />}/>
      <Route path='/auth/register' element={ <Register />}/>
      <Route path='/Profile' element={ <Profile />}/>
      <Route path='/Main' element={ <Main />}/>
    </Routes>
  )
}

export default App