import './App.css'
import { BrowserRouter as Router, Route, Routes, Link } from 'react-router-dom'
import Home from './pages/Home'
import Auth from './pages/Auth'

function App() {
  return (
    <Router>
      <div className="container">
        <header>
          <h1>Ryan Bucks</h1>
          <nav>
            <Link to="/">Home</Link>
            <Link to="/send">Send</Link>
            <Link to="/receive">Receive</Link>
            <Link to="/claim">Claim Codes</Link>
            <Link to="/auth">Login/Register</Link>
          </nav>
        </header>
        <main>
          <Routes>
            <Route path="/" element={<Home />} />
            <Route path="/auth" element={<Auth />} />
            {/* Add other routes here */}
          </Routes>
        </main>
      </div>
    </Router>
  )
}

export default App
