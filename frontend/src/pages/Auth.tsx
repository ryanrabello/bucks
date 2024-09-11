import { useState } from 'react'

export default function Auth() {
  const [isLogin, setIsLogin] = useState(true)

  const toggleAuthMode = () => {
    setIsLogin(!isLogin)
  }

  return (
    <div className="auth-container">
      <h2>{isLogin ? 'Login' : 'Register'}</h2>
      <form>
        <input type="email" placeholder="Email" required />
        <input type="password" placeholder="Password" required />
        {!isLogin && <input type="password" placeholder="Confirm Password" required />}
        <button type="submit">{isLogin ? 'Login' : 'Register'}</button>
      </form>
      <p>
        {isLogin ? "Don't have an account? " : "Already have an account? "}
        <button onClick={toggleAuthMode}>
          {isLogin ? 'Register' : 'Login'}
        </button>
      </p>
    </div>
  )
}