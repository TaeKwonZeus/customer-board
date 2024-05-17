import './App.css'
import { Outlet } from 'react-router-dom'
import { Nav } from '@/components/nav.tsx'
import { Annoyed } from 'lucide-react'

function App() {
  return (
    <>
      <Nav isCollapsed={false} links={[
        {
          title: 'JA PIERDOLE',
          href: 'https://google.com/',
          icon: Annoyed,
          variant: 'default',
        }]}/>
      <div id="detail">
        <Outlet/>
      </div>
    </>
  )
}

export default App
