import React, { useState } from 'react'
import axios from 'axios'
import { useNavigate } from 'react-router-dom'

const Login = () => {
  const [email, setEmail] = useState('')
  const [password, setPassword] = useState('')
  const [popup, setPopup] = useState({ show: false, success: false, message: '' })
  const navigate = useNavigate()

  const handleSubmit = async (e) => {
    e.preventDefault()
    try {
      const response = await axios.post('http://localhost:3000/api/user/login', { email, password })
      if (response.data.statusCode === 200) {
        localStorage.setItem('accessToken', response.data.accessToken)
        setPopup({ show: true, success: true, message: response.data.message })
        setTimeout(() => {
          setPopup({ show: false, success: false, message: '' })
          navigate('/dashboard')
        }, 1200)
      } else {
        setPopup({ show: true, success: false, message: response.data.message || 'Login failed' })
        setTimeout(() => setPopup({ show: false, success: false, message: '' }), 3000)
      }
    } catch (err) {
      setPopup({ show: true, success: false, message: err.response?.data?.message || 'Login failed' })
      setTimeout(() => setPopup({ show: false, success: false, message: '' }), 3000)
    }
  }

  return (
    <div className="flex items-center justify-center min-h-[60vh]">
      <div className="w-full max-w-md bg-white rounded-lg shadow-md p-8">
        <h2 className="text-2xl font-bold mb-6 text-center text-gray-800">Login to your account</h2>
        <form className="space-y-5" onSubmit={handleSubmit}>
          <div>
            <label htmlFor="email" className="block text-sm font-medium text-gray-700 mb-1">Email</label>
            <input
              type="email"
              id="email"
              className="w-full px-4 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
              placeholder="Enter your email"
              required
              value={email}
              onChange={e => setEmail(e.target.value)}
            />
          </div>
          <div>
            <label htmlFor="password" className="block text-sm font-medium text-gray-700 mb-1">Password</label>
            <input
              type="password"
              id="password"
              className="w-full px-4 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
              placeholder="Enter your password"
              required
              value={password}
              onChange={e => setPassword(e.target.value)}
            />
          </div>
          <button
            type="submit"
            className="w-full py-2 px-4 bg-blue-600 text-white font-semibold rounded-md hover:bg-blue-700 transition-colors"
          >
            Login
          </button>
        </form>
        {popup.show && (
          <div className={`fixed top-6 left-1/2 transform -translate-x-1/2 z-50 flex items-center gap-2 px-6 py-3 rounded shadow-lg text-white ${popup.success ? 'bg-green-600' : 'bg-red-600'}`}>
            {popup.success ? (
              <svg className="w-6 h-6 text-white" fill="none" stroke="currentColor" strokeWidth="2" viewBox="0 0 24 24"><path strokeLinecap="round" strokeLinejoin="round" d="M5 13l4 4L19 7" /></svg>
            ) : (
              <svg className="w-6 h-6 text-white" fill="none" stroke="currentColor" strokeWidth="2" viewBox="0 0 24 24"><path strokeLinecap="round" strokeLinejoin="round" d="M6 18L18 6M6 6l12 12" /></svg>
            )}
            <span>{popup.message}</span>
          </div>
        )}
      </div>
    </div>
  )
}

export default Login 