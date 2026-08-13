
import { BrowserRouter, Routes, Route} from 'react-router'
import './App.css'
import Videos from './assets/Videos'
import Video from "./assets/Video"
// import { useEffect } from 'react'

function App() {
  return (
    <BrowserRouter>
      <Routes>
        <Route path='/' element={<Videos />} />
        <Route path='/Videos/:id' element={<Video />} />
      </Routes>
    </BrowserRouter>
  )
}

export default App
