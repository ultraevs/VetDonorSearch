import React from 'react'
import Auth from "./pages/Auth/Auth"
import Register from './pages/Register/Register'
import Profile from "./pages/Profile/Profile"
import Main from './pages/Main/Main'
import Info from './pages/Info/Info'
import ProfileClinic from './pages/ProfileClinic/ProfileClinic'
import Forgot from './pages/Forgot/Forgot'

import { Route, Routes } from 'react-router-dom'
import Where from './pages/Where/Where'

const App = () => {
  return (
    <Routes>
      <Route path='/' element={ <Main />}/>
      <Route path='/auth' element={ <Auth />}/>
      <Route path='/auth/register' element={ <Register />}/>
      <Route path="/Forgot" element={<Forgot/>} />
      <Route path='/Profile' element={ <Profile />}/>
      <Route path='/Where' element={ <Where />}/>
      <Route path='/Info' element={<Info/>}/>
      <Route path='/ProfileClinic' element={ <ProfileClinic />}/>
    </Routes>
  )
}

export default App